// lib/ai.ts
import type { ChatTransport, UIMessage, UIMessageChunk } from "ai"

type Step = { role: "user" | "assistant"; text: string; sleepMs: number }

const wait = (ms: number) => new Promise((r) => setTimeout(r, ms))

export function getMessageText(message: UIMessage) {
  return message.parts.map((p) => (p.type === "text" ? p.text : "")).join("")
}

class ChatScript {
  private steps: Step[] = []
  private pendingSleep = 0

  user(text: string) {
    this.steps.push({ role: "user", text, sleepMs: 0 })
    return this
  }
  sleep(ms: number) {
    this.pendingSleep = ms
    return this
  }
  assistant(text: string) {
    this.steps.push({ role: "assistant", text, sleepMs: this.pendingSleep })
    this.pendingSleep = 0
    return this
  }

  private toMessage(step: Step, i: number): UIMessage {
    return { id: `msg-${i}`, role: step.role, parts: [{ type: "text", text: step.text }] }
  }

  get(count: number): UIMessage[] {
    return this.steps.slice(0, count).map((s, i) => this.toMessage(s, i))
  }

  next(messages: UIMessage[]): UIMessage | undefined {
    const step = this.steps[messages.length]
    return step?.role === "user" ? this.toMessage(step, messages.length) : undefined
  }

  transport({ delayMs = 20 } = {}): ChatTransport<UIMessage> {
    return {
      sendMessages: async ({ messages, abortSignal }) => {
        const step = this.steps[messages.length]
        return new ReadableStream<UIMessageChunk>({
          start: async (controller) => {
            try {
              if (!step || step.role !== "assistant") return controller.close()
              await wait(step.sleepMs)
              controller.enqueue({ type: "start" })
              controller.enqueue({ type: "text-start", id: "t" })
              for (const chunk of step.text.match(/\S+\s*/g) ?? []) {
                if (abortSignal?.aborted) break
                controller.enqueue({ type: "text-delta", id: "t", delta: chunk })
                await wait(delayMs)
              }
              controller.enqueue({ type: "text-end", id: "t" })
              controller.enqueue({ type: "finish" })
              controller.close()
            } catch (e) {
              controller.error(e)
            }
          },
        })
      },
      reconnectToStream: async () => null,
    }
  }
}

export const createChat = () => new ChatScript()