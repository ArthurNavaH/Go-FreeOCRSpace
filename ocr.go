package gofreeocr

const (
	// URL de la API POST
	APIURL = "https://api.ocr.space/parse/image"

	// Lista de Filetypes soportados
	PDF = "PDF"
	GIF = "GIF"
	PNG = "PNG"
	JPG = "JPG"
	TIF = "TIF"
	BMP = "BMP"
)

//Client Cliente de la API
type Client struct {
	apikey            string
	isOverlayRequired bool
}

//NewClient Constructor Client
func NewClient(ApiKey string, isOverlay bool) *Client {
	return &Client{
		apikey:            ApiKey,
		isOverlayRequired: isOverlay,
	}
}

//Image Imagen a enviar
type Image struct {
	url      string
	language string
	filetype string
}

//NewImage Constructor Image
func NewImage(url string, language string, filetype string) *Image {
	return &Image{
		url:      url,
		language: language,
		filetype: filetype,
	}
}
