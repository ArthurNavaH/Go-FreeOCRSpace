# Go-FreeOCRSpace
Libreria de Golang para usar la OCR API de OCR.Space como servicio.

# Ejemplo de uso
`
client := gofreeocr.NewClient("00afe515fd88957", "true")

img := gofreeocr.NewImage(
    "https://host.com/tu/imagen",   // URL de la Imagen
    "spa",  // Lenguaje de la imagen
    gofreeocr.UNDEFINED // Tipo de archivo de la imagen
)

response, err := client.SendImage(img, gofreeocr.POST)  // Se envia la imagen, Declarando el metodo a utilizar
if err != nil {
    log.Println(err)
}

fmt.Println(response.Text())
`
