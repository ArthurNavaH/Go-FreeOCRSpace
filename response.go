package gofreeocr

//Response Respuesta de la API
type Response struct {
	ErrorDetails          string
	ErrorMessage          string
	IsErroredOnProcessing bool
	OCRExitCode           string
	ParsedResults         []ParsedResult

	ProcessingTimeInMilliseconds string
	SearchablePDFURL             string
}

type ParsedResult struct {
	ErrorDetails      string
	ErrorMessage      string
	FileParseExitCode string
	ParsedText        string
	TextOverlay       struct {
		HasOverlay bool
		Lines      []Line
	}
	Message string
}

type Line struct {
	MaxHeight int
	MinTop    int
	Words     []Word
}

type Word struct {
	Height   int
	Left     int
	Top      int
	Width    int
	WordText string
}
