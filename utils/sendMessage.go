package utils

import (
	"context"
	_ "embed"
	"github.com/golang/protobuf/proto"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"os"
)

func SendMessage(client *whatsmeow.Client, groupJID string) {
	wabotFolderPath := FolderPath()
	if wabotFolderPath == "" {
		panic("Error al obtener la ruta de la carpeta wabot")
	}

	// Leer la imagen del archivo
	imagePath := wabotFolderPath + "/cuentas-hs.png"
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		panic(err)
	}

	// Subir la imagen a WhatsApp
	resp, err := client.Upload(context.Background(), imageData, whatsmeow.MediaImage)
	if err != nil {
		panic(err)
	}

	imageMsg := &waProto.ImageMessage{
		Caption: proto.String("*Disponibles cuentas premium sin caidas 🔥* \n\n - *Netflix" +
			" original:* 3$ \n - *HBO Max:* 2$ \n - *Disney Plus:* 1." +
			"5$ \n - *Star Plus:* 1$ \n - *Spotify:* desde 2$ \n - *Prime video:* 1.5" +
			"5$ \n - *Crunchyroll:* 2$ \n\n😉 Aceptamos *Bolivares, Binance, Paypal, Zinli, " +
			"Bancolombia y efectivo* \n\n🦇 *Siguenos en instagram para no perderte de nada!* https" +
			"://instagram.com/house.streamingxx"),
		Mimetype: proto.String("image/png"), // replace this with the actual mime type
		// you can also optionally add other fields like ContextInfo and JpegThumbnail here

		Url:           &resp.URL,
		DirectPath:    &resp.DirectPath,
		MediaKey:      resp.MediaKey,
		FileEncSha256: resp.FileEncSHA256,
		FileSha256:    resp.FileSHA256,
		FileLength:    &resp.FileLength,
	}

	_, err = client.SendMessage(context.Background(), types.JID{
		User:   groupJID,
		Server: types.GroupServer,
	}, &waProto.Message{
		ImageMessage: imageMsg,
	})
	if err != nil {
		panic(err)
	}
}
