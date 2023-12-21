package utils

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func FolderPath() string {
	// Obtener información del usuario actual
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error al obtener información del usuario:", err)
		return ""
	}
	// Construir la ruta completa hacia la carpeta "wabot" en la carpeta de documentos del usuario
	wabotFolderPath := filepath.Join(currentUser.HomeDir, "Documents", "wabot")

	// Verificar si la carpeta "wabot" existe
	if _, err := os.Stat(wabotFolderPath); os.IsNotExist(err) {
		// La carpeta no existe, crearla
		err := os.MkdirAll(wabotFolderPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			return ""
		}

	}

	return wabotFolderPath
}
