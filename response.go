package gofreeocr

//Response Respuesta de la API
type Response struct {
	ErrorDetails          string
	ErrorMessage          string
	IsErroredOnProcessing bool
	OCRExitCode           int
	ParsedResults         []ParsedResult

	ProcessingTimeInMilliseconds string
	SearchablePDFURL             string
}

type ParsedResult struct {
	ErrorDetails      string
	ErrorMessage      string
	FileParseExitCode int
	ParsedText        string
	TextOverlay       struct {
		HasOverlay bool
		Lines      []Line
	}
	Message string
}

type Line struct {
	MaxHeight float32
	MinTop    float32
	Words     []Word
}

type Word struct {
	Height   float32
	Left     float32
	Top      float32
	Width    float32
	WordText string
}

//TextPretty Retorna el texto original formateado como lo devolvio el API
func (response Response) TextPretty() string {
	return response.ParsedResults[0].ParsedText
}

//Text Retorna el texto sin saltos de lineas y retornos de carro
func (response Response) Text() string {
	text := []byte(response.ParsedResults[0].ParsedText)

	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '\n', '\r':
			// text[i] = ' '
			text = append(text[:i], text[i+2:]...)
		}
	}

	return string(text)
}
