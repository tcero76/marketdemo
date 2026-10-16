import type { ChatTurn, FormState } from '@/lib/form-schema';

export type MessagesProps = {
    messages:Message[]
}

export type Message = {
    userId:string
    message:string
    extName:string
  }

export type SimpleChatProps = {
  roomName?: string;
}

export type ChatMessage = {
  id: string;
  role: "user" | "assistant";
  content: string;
};

export type ChatRequest = {
  message: string;
  conversationId:string
};
 
export type ChatSuccess = {
  fields: FormState;
  reply: string;
  done: boolean;
};