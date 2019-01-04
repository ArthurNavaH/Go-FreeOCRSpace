package gofreeocr

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
	isOverlayRequired string
}

//NewClient Constructor Client
func NewClient(ApiKey string, isOverlay string) *Client {
	return &Client{
		apikey:            ApiKey,
		isOverlayRequired: isOverlay,
	}
}

//SendImage Envia la imagen a la API
func (client *Client) SendImage(img *Image) (response *Response, err error) {
	responseApi, err := http.PostForm(APIURL, url.Values{
		"apikey":            {client.apikey},
		"url":               {img.url},
		"isoverlayrequired": {client.isOverlayRequired},
		"language":          {img.language},
	})
	defer responseApi.Body.Close()

	body, err := ioutil.ReadAll(responseApi.Body)

	err = json.Unmarshal(body, &response)

	return
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
