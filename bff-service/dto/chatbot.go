package dto

type ChatRequest struct {
	ConversationID string `json:"conversationId"`
	Message        string `json:"message"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}
