package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOpenAICompatClient(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/engines/v1/chat/completions" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"choices":[{"message":{"content":"{\"nombre\":\"Ana\"}"}}]}`)
	}))
	defer srv.Close()

	c := OpenAICompatClient{BaseURL: srv.URL + "/engines/v1", Model: "qwen2.5"}
	out, err := c.Complete(context.Background(), "hola")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "Ana") {
		t.Fatalf("respuesta inesperada: %q", out)
	}
}

func TestOpenAICompatClientHTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "model not found", http.StatusNotFound)
	}))
	defer srv.Close()

	c := OpenAICompatClient{BaseURL: srv.URL, Model: "x"}
	if _, err := c.Complete(context.Background(), "hola"); err == nil || !strings.Contains(err.Error(), "404") {
		t.Fatalf("debía devolver error con el código 404, got: %v", err)
	}
}

func TestOpenAICompatClientCallTool(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tools, _ := body["tools"].([]any)
		if len(tools) != 1 || body["tool_choice"] != "required" {
			http.Error(w, "falta tools o tool_choice", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"choices":[{"message":{"content":null,"tool_calls":[{"type":"function","function":{"name":"registrar_datos","arguments":"{\"nombre\":\"Ana\"}"}}]}}]}`)
	}))
	defer srv.Close()

	c := OpenAICompatClient{BaseURL: srv.URL, Model: "qwen2.5"}
	args, err := c.CallTool(context.Background(), "sys", "user", ToolSpec{
		Name: "registrar_datos", Description: "d", Parameters: map[string]any{"type": "object"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(args, "Ana") {
		t.Fatalf("argumentos inesperados: %q", args)
	}
}
