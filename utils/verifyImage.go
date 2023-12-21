package utils

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func VerifyImage() {
	// Obtener información del usuario actual
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error al obtener información del usuario:", err)
	}

	// Construir la ruta completa hacia la carpeta "wabot" en la carpeta de documentos del usuario
	wabotFolderPath := filepath.Join(currentUser.HomeDir, "Documents", "wabot")

	// Verificar si la carpeta "wabot" existe
	if _, err := os.Stat(wabotFolderPath); os.IsNotExist(err) {
		// La carpeta no existe, crearla
		err := os.MkdirAll(wabotFolderPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error al crear la carpeta wabot:", err)
		}

	}

	// Nombre del archivo a verificar y descargar
	fileName := "cuentas-hs.png"
	filePath := filepath.Join(wabotFolderPath, fileName)

	// Verificar si el archivo existe
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// El archivo no existe, descargarlo desde el enlace de Google Drive
		err := DownloadImage("https://res.cloudinary.com/mendoza000/image/upload/v1703182497/ujrepbomnxrhb9xusznj.png", filePath)
		if err != nil {
			fmt.Println("Error al descargar la imagen:", err)
		}

	}
}
