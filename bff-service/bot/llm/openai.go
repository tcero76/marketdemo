package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// OpenAICompatClient habla con cualquier servidor de API compatible con OpenAI
// (Docker Model Runner, Ollama /v1, llama.cpp, vLLM...). Implementa:
//   - Completer  (prompt → texto)
//   - ToolCaller (function calling → argumentos JSON)
type OpenAICompatClient struct {
	BaseURL string       // p. ej. http://localhost:12434/engines/v1
	Model   string       // p. ej. qwen2.5
	APIKey  string       // opcional (los servidores locales suelen no pedirla)
	HTTP    *http.Client // opcional; por defecto timeout de 60 s
}

type chatResponse struct {
	Choices []struct {
		Message struct {
			Content   string `json:"content"`
			ToolCalls []struct {
				Function struct {
					Name      string          `json:"name"`
					Arguments json.RawMessage `json:"arguments"`
				} `json:"function"`
			} `json:"tool_calls"`
		} `json:"message"`
	} `json:"choices"`
}

func (c OpenAICompatClient) chat(ctx context.Context, payload map[string]any) (*chatResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := strings.TrimRight(c.BaseURL, "/") + "/chat/completions"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	client := c.HTTP
	if client == nil {
		client = &http.Client{Timeout: 60 * time.Second}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("llamando al LLM (%s): %w", url, err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("leyendo respuesta del LLM: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("el LLM respondió %d: %s", resp.StatusCode, data[:min(len(data), 200)])
	}

	var out chatResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("respuesta del LLM no es JSON válido: %w", err)
	}
	if len(out.Choices) == 0 {
		return nil, fmt.Errorf("el LLM no devolvió choices")
	}
	return &out, nil
}

// Complete envía un prompt y devuelve el texto de la respuesta.
func (c OpenAICompatClient) Complete(ctx context.Context, prompt string) (string, error) {
	out, err := c.chat(ctx, map[string]any{
		"model":       c.Model,
		"temperature": 0, // extracción: queremos respuestas deterministas
		"messages": []map[string]string{
			{"role": "system", "content": "Eres un extractor de datos. Respondes únicamente con JSON válido."},
			{"role": "user", "content": prompt},
		},
	})
	if err != nil {
		return "", err
	}
	return out.Choices[0].Message.Content, nil
}

// CallTool expone UNA función al modelo y exige que la invoque (tool_choice "required",
// la forma más portable: algunos servidores no aceptan el objeto que fuerza una función
// concreta). Devuelve los argumentos como texto JSON.
func (c OpenAICompatClient) CallTool(ctx context.Context, system, user string, tool ToolSpec) (string, error) {
	out, err := c.chat(ctx, map[string]any{
		"model":       c.Model,
		"temperature": 0,
		"messages": []map[string]string{
			{"role": "system", "content": system},
			{"role": "user", "content": user},
		},
		"tools": []map[string]any{{
			"type": "function",
			"function": map[string]any{
				"name":        tool.Name,
				"description": tool.Description,
				"parameters":  tool.Parameters,
			},
		}},
		"tool_choice": "required",
	})
	if err != nil {
		return "", err
	}

	msg := out.Choices[0].Message
	if len(msg.ToolCalls) > 0 {
		raw := msg.ToolCalls[0].Function.Arguments
		if len(raw) > 0 && raw[0] == '"' { // el estándar OpenAI envía los argumentos como string
			var s string
			if err := json.Unmarshal(raw, &s); err != nil {
				return "", fmt.Errorf("argumentos de la herramienta ilegibles: %w", err)
			}
			return s, nil
		}
		return string(raw), nil // algunos servidores los envían como objeto
	}
	if strings.TrimSpace(msg.Content) != "" {
		return msg.Content, nil // modelos pequeños a veces emiten <tool_call>… como texto
	}
	return "", ErrNotUnderstood
}
