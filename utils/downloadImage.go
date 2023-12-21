package utils

import (
	"io"
	"net/http"
	"os"
)

func DownloadImage(url string, destination string) error {
	// Realizar solicitud HTTP para obtener el archivo
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	// Crear el archivo en la ubicación de destino
	file, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Copiar el contenido del cuerpo de la respuesta al archivo
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
