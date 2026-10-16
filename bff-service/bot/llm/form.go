package llm

// Answers guarda los valores ya capturados, indexados por nombre de campo.
type Answers map[string]string

// Field describe un paso del formulario. Es solo DATOS + estrategias:
// no contiene lógica de flujo, así que agregar un campo no toca el motor.
type Field struct {
	Name   string
	Label  string // nombre legible para resúmenes y correcciones (opcional)
	Prompt string // lo que se le pregunta al usuario
	Hint   string // descripción para el LLM (p. ej. "fecha en formato YYYY-MM-DD"); si está vacío se usa Prompt
	Type   string   // tipo JSON Schema para tool calling: "string" (por defecto), "integer", "number", "boolean"
	Enum   []string // valores permitidos: van en el schema Y se validan automáticamente

	// Extractor solo lo usa PerFieldExtractor (modo reglas). En modo LLM se ignora,
	// por eso las reglas que DEBEN cumplirse van siempre en Validators.
	Extractor  Extractor
	Validators []Validator
	Skip       func(Answers) bool // omitir el campo según respuestas previas
}

// Form es una definición declarativa y ordenada de campos.
type Form struct {
	ID     string
	Fields []Field
}

// Pending devuelve, en orden, los campos sin responder y que no se omiten.
func (f Form) Pending(a Answers) []Field {
	var out []Field
	for _, fld := range f.Fields {
		if _, filled := a[fld.Name]; filled {
			continue
		}
		if fld.Skip != nil && fld.Skip(a) {
			continue
		}
		out = append(out, fld)
	}
	return out
}

// Next devuelve el siguiente campo pendiente. El "estado" de la máquina
// se deriva de las respuestas: no hay tabla de transiciones que mantener.
func (f Form) Next(a Answers) (Field, bool) {
	if p := f.Pending(a); len(p) > 0 {
		return p[0], true
	}
	return Field{}, false
}

// DisplayName devuelve el Label si existe, o el Name.
func (f Field) DisplayName() string {
	if f.Label != "" {
		return f.Label
	}
	return f.Name
}

// Filled devuelve, en orden, los campos que ya tienen respuesta (candidatos a revisión).
func (f Form) Filled(a Answers) []Field {
	var out []Field
	for _, fld := range f.Fields {
		if _, ok := a[fld.Name]; ok {
			out = append(out, fld)
		}
	}
	return out
}

// AllValidators agrega la validación automática de Enum a los Validators del campo.
func (f Field) AllValidators() []Validator {
	if len(f.Enum) == 0 {
		return f.Validators
	}
	return append([]Validator{OneOf(f.Enum...)}, f.Validators...)
}
