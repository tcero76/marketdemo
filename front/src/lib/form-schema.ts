// Definición única del formulario. La usan tanto el route handler (servidor)
// como el componente de chat (cliente). Edita solo este archivo para agregar
// o quitar campos.
 
export const FORM_FIELDS = {
  nombre: { label: 'Nombre completo', description: 'Nombre completo de la persona' },
  email: { label: 'Correo electrónico', description: 'Correo electrónico válido' },
  telefono: { label: 'Teléfono', description: 'Número de teléfono de contacto' },
  mensaje: { label: 'Motivo de contacto', description: 'Motivo de contacto o mensaje' },
} as const;
 
export type FieldKey = keyof typeof FORM_FIELDS;
export const REQUIRED_FIELDS = Object.keys(FORM_FIELDS) as FieldKey[];
 
export type FormState = Partial<Record<FieldKey, string>>;
export type ChatTurn = { role: 'user' | 'assistant'; content: string };
 
export type ChatResponse =
  | { fields: FormState; reply: string; done: boolean }
  | { error: string };
 