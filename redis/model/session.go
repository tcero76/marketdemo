package model

type SessionData struct {
	AccessToken  string `redis:"access_token"`
	RefreshToken string `redis:"refresh_token"`
	UserID       string `redis:"user_id"`
	// UserAgent    string `redis:"user_agent"`
	// IP           string `redis:"ip"`
	Idp             string       `redis:"idp"`
	SessionID       string       `redis:"session_id"`
	IsAuthenticated bool         `redis:"is_authenticated"`
	LoginChallenge  string       `redis:"login_challenge"`
	BotSession      *FormSession `redis:"bot_session"`
}

const (
	PhaseCollecting Phase = ""
	PhaseConfirming Phase = "confirming" // mostró el resumen y espera sí/no
	PhaseCorrecting Phase = "correcting" // el usuario dijo "no" y debe elegir qué dato corregir
)

// Phase (State) indica en qué fase está el diálogo. El valor cero es "recolectando".
type Phase string

// Answers guarda los valores ya capturados, indexados por nombre de campo.
type Answers map[string]string

type FormSession struct {
	ID      string   `redis:"id"`
	FormID  string   `redis:"form_id"`
	Answers Answers  `redis:"answers"`
	History []string `redis:"history"`
	Done    bool     `redis:"done"`
	Phase   Phase    `redis:"phase"`
}

func (s *FormSession) Clone() *FormSession {
	c := *s
	c.Answers = make(Answers, len(s.Answers))
	for k, v := range s.Answers {
		c.Answers[k] = v
	}
	c.History = append([]string(nil), s.History...)
	return &c
}
