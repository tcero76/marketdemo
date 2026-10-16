"use client";

import { useState } from "react";
import { MessageCircle, Minus, Send } from "lucide-react";

import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Message,
  MessageContent,
} from "@/components/ui/message";
import { Bubble, BubbleContent } from "@/components/ui/bubble";
import {
  MessageScroller,
  MessageScrollerButton,
  MessageScrollerContent,
  MessageScrollerItem,
  MessageScrollerProvider,
  MessageScrollerViewport,
} from "@/components/ui/message-scroller";
import { Input } from "@/components/ui/input";
import { useSendChatMessageMutation } from "@/http/api";
import { ChatMessage, ChatSuccess } from "@/types";

const initialMessages: ChatMessage[] = [
  {
    id: "1",
    role: "assistant",
    content:
      "Hola 👋 Soy el asistente de documentación. ¿Qué quieres saber?",
  },
  {
    id: "2",
    role: "user",
    content: "¿Qué es PostgreSQL?",
  },
  {
    id: "3",
    role: "assistant",
    content:
      "PostgreSQL es un sistema de gestión de bases de datos relacional y objeto-relacional de código abierto.",
  },
];

export function ChatBox() {
  const [open, setOpen] = useState(true);
  const [messages, setMessages] = useState(initialMessages);
  const [input, setInput] = useState("");
  const [sendChatMessage, { isLoading: loading }] = useSendChatMessageMutation();

  function sendMessage() {
    const content = input.trim();

    if (!content) return;

    setMessages((current) => [
      ...current,
      {
        id: crypto.randomUUID(),
        role: "user",
        content,
      },
    ]);

    sendChatMessage({message:content, conversationId: "abc"}).unwrap()
    .then((cs:ChatSuccess) => {
      setMessages((current) => [
        ...current,
        {
          id:crypto.randomUUID(),
          role: "assistant",
          content: cs.reply
        }
      ])
    })
    setInput("");
  }

  return (
    <div className="fixed bottom-4 right-4 z-50">
      {!open ? (
        <Button
          size="icon-lg"
          className="rounded-full shadow-lg"
          onClick={() => setOpen(true)}
          aria-label="Abrir chat"
        >
          <MessageCircle />
        </Button>
      ) : (
        <Card className="flex h-[600px] w-[380px] flex-col gap-0 overflow-hidden shadow-xl">
          <CardHeader className="flex shrink-0 flex-row items-center justify-between border-b">
            <CardTitle className="text-base">
              PostgreSQL Assistant
            </CardTitle>

            <Button
              variant="ghost"
              size="icon"
              onClick={() => setOpen(false)}
              aria-label="Minimizar chat"
            >
              <Minus />
            </Button>
          </CardHeader>

          <CardContent className="min-h-0 flex-1 p-0">
            <MessageScrollerProvider>
              <MessageScroller>
                <MessageScrollerViewport>
                  <MessageScrollerContent className="gap-4 p-4">
                    {messages.map((message) => (
                      <MessageScrollerItem
                        key={message.id}
                        messageId={message.id}
                        scrollAnchor={message.role === "user"}
                      >
                        <Message align={message.role === "user" ? "end" : "start"}>
                          <MessageContent>
                            <Bubble
                              variant={
                                message.role === "user"
                                  ? "default"
                                  : "ghost"
                              }
                            >
                              <BubbleContent>
                                {message.content}
                              </BubbleContent>
                            </Bubble>
                          </MessageContent>
                        </Message>
                      </MessageScrollerItem>
                    ))}
                  </MessageScrollerContent>
                </MessageScrollerViewport>

                <MessageScrollerButton />
              </MessageScroller>
            </MessageScrollerProvider>
          </CardContent>

          <CardFooter className="shrink-0 border-t p-3">
            <form
              className="flex w-full gap-2"
              onSubmit={(event) => {
                event.preventDefault();
                sendMessage();
              }}
            >
              <Input
                value={input}
                onChange={(event) => setInput(event.target.value)}
                placeholder="Escribe tu pregunta..."
              />

              <Button type="submit" size="icon" aria-label="Enviar">
                <Send />
              </Button>
            </form>
          </CardFooter>
        </Card>
      )}
    </div>
  );
}