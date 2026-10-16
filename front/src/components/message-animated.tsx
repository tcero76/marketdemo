// src/components/message-animated.tsx
"use client"

import type { UIMessage } from "ai"

import { getMessageText } from "@/lib/ai"
import { cn } from "@/lib/utils"
import { MessageScrollerItem } from "@/components/ui/message-scroller"

export function MessageAnimated({
  message,
  scrollAnchor,
}: {
  message: UIMessage
  scrollAnchor?: boolean
}) {
  const isUser = message.role === "user"

  return (
    <MessageScrollerItem
      scrollAnchor={scrollAnchor}
      className={cn(
        "flex animate-in fade-in-0 slide-in-from-bottom-2 duration-300",
        isUser ? "justify-end" : "justify-start"
      )}
    >
      <div
        className={cn(
          "max-w-[85%] rounded-2xl px-3 py-2 text-sm whitespace-pre-wrap",
          isUser
            ? "bg-primary text-primary-foreground"
            : "bg-muted text-foreground"
        )}
      >
        {getMessageText(message)}
      </div>
    </MessageScrollerItem>
  )
}