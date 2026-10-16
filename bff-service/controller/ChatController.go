package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	echo "github.com/labstack/echo/v4"
	basicBot "github.com/tcero76/marketplace/bff-service/bot/basic"
	llmBot "github.com/tcero76/marketplace/bff-service/bot/llm"
)

type chatRequest struct {
	ConversationID string `json:"conversationId"`
	Message        string `json:"message"`
}

type chatResponse struct {
	Reply string `json:"reply"`
}

func ChatHandler(b *basicBot.Bot) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("Entró a ChatHandler")
		c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, 16<<10)
		var req chatRequest
		if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil ||
			req.ConversationID == "" {
			return c.JSON(http.StatusBadRequest,
				map[string]string{"error": "conversationId y message son obligatorios"})

		}
		ctx, cancel := context.WithTimeout(c.Request().Context(), 30*time.Second)
		defer cancel()
		reply, err := b.Process(ctx, req.ConversationID, req.Message)
		if err != nil {
			log.Printf("process error: %v", err)
			return c.JSON(http.StatusInternalServerError,
				map[string]string{"error": "Error procesando el mensaje"})

		}
		return c.JSON(http.StatusOK, chatResponse{Reply: reply.Text})
	}
}

func ChatLLMHandler(b *llmBot.Processor) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Entró a ChatLLMHandler")
		auth := c.Request().Header.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			log.Println("Missing or invalid Authorization header")
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid Authorization header")
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, 16<<10)
		var req chatRequest
		if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil ||
			req.ConversationID == "" {
			return c.JSON(http.StatusBadRequest,
				map[string]string{"error": "conversationId y message son obligatorios"})
		}
		ctx, cancel := context.WithTimeout(c.Request().Context(), 30*time.Second)
		defer cancel()
		reply, err := b.Handle(ctx, token, req.Message)
		if err != nil {
			log.Println("error:", err)
			return c.JSON(http.StatusBadRequest, chatResponse{Reply: err.Error()})
		}
		if !reply.Done {
			return c.JSON(http.StatusOK, chatResponse{Reply: reply.Text})
		}
		return c.JSON(http.StatusOK, chatResponse{Reply: "Conversación terminada"})
	}
}

func verifyWebhook(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("hub.mode")
	log.Println("mode: ", mode)
	token := r.URL.Query().Get("hub.verify_token")
	log.Println("token: ", token)
	challenge := r.URL.Query().Get("hub.challenge")
	log.Println("challenge: ", challenge)

	verifyToken := os.Getenv("WHATSAPP_VERIFY_TOKEN")

	if mode != "subscribe" || token != verifyToken {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(challenge))
}

func receiveWebhook(w http.ResponseWriter, r *http.Request) {
	log.Println("Webhook recibido")
	w.WriteHeader(http.StatusOK)
}

func WebhookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Ingresó WebhookHandler")
		switch c.Request().Method {
		case http.MethodGet:
			verifyWebhook(c.Response(), c.Request())
		case http.MethodPost:
			receiveWebhook(c.Response(), c.Request())
		default:
			http.Error(c.Response(), "method not allowed", http.StatusMethodNotAllowed)
		}
		return nil
	}
}
