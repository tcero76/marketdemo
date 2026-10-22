package llm

import (
	"context"
	"log"
	"time"
)

// ---- Decorator: agrega logs a cualquier Completer sin tocar el original ----

type loggingCompleter struct{ next Completer }

func (l loggingCompleter) Complete(ctx context.Context, prompt string) (string, error) {
	start := time.Now()
	out, err := l.next.Complete(ctx, prompt)
	log.Printf("[llm %v] respuesta=%q err=%v", time.Since(start).Round(time.Millisecond), out, err)
	return out, err
}

type loggingToolCaller struct{ next ToolCaller }

func (l loggingToolCaller) CallTool(ctx context.Context, system, user string, tool ToolSpec) (string, error) {
	start := time.Now()
	out, err := l.next.CallTool(ctx, system, user, tool)
	log.Printf("[llm tool %v] args=%q err=%v", time.Since(start).Round(time.Millisecond), out, err)
	return out, err
}
