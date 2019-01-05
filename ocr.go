package gofreeocr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// URL de la API POST
	APIURL_POST = "https://api.ocr.space/parse/image"
	// URL de la API GET
	APIURL_GET = "https://api.ocr.space/parse/imageurl"

	GET  = "GET"
	POST = "POST"

	// Lista de Filetypes soportados
	PDF = "PDF"
	GIF = "GIF"
	PNG = "PNG"
	JPG = "JPG"
	TIF = "TIF"
	BMP = "BMP"

	UNDEFINED = "UNDEFINED"
)

//Client Cliente de la API
type Client struct {
	apiurl            string
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
func (client *Client) SendImage(img *Image, method string) (response *Response, err error) {
	var responseApi *http.Response

	if method == POST {

		responseApi, err = http.PostForm(APIURL_POST, url.Values{
			"apikey":            {client.apikey},
			"url":               {img.url},
			"isoverlayrequired": {client.isOverlayRequired},
			"language":          {img.language},
		})
		defer responseApi.Body.Close()

		body, err := ioutil.ReadAll(responseApi.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &response)

	} else if method == GET {
		responseApi, err = http.Get(fmt.Sprintf("%s?apikey=%s&url=%s&language=%s&isOverlayRequired=%s&filetype=%s",
			APIURL_GET, client.apikey, img.url, img.language, client.isOverlayRequired, img.filetype,
		))

		defer responseApi.Body.Close()

		body, err := ioutil.ReadAll(responseApi.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &response)
	} else {
		err = errors.New(fmt.Sprintf("Metodo de la peticion no soportado : %s", method))
	}

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
