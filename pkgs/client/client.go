package client

import (
	"context"
	"fmt"
	"github.com/mymmsc/gox/qrterminal"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"wabot/utils"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func ConnectClient(opt int) {
	dbLog := waLog.Stdout("Database", "INFO", true)

	wabotFolderPath := utils.FolderPath()
	if wabotFolderPath == "" {
		panic("Error al obtener la ruta de la carpeta wabot")
	}

	utils.VerifyImage()

	dbPath := wabotFolderPath + "/data.db"
	dbPath = "file:" + filepath.ToSlash(dbPath) + "?_foreign_keys=on"

	container, err := sqlstore.New("sqlite3", dbPath, dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Noop
	client := whatsmeow.NewClient(deviceStore, clientLog)

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("Scan this QR code with your phone:")
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err = client.Connect()
		if err != nil {
			panic(err)
		}

		if opt == 1 {
			utils.GetGroup(client)
		}

		if opt == 2 {
			SendClient15(client, false)
		}

		if opt == 3 {
			SendClient15(client, true)
		}

		if opt == 4 {
			SendClient25(client, false)
		}

		if opt == 5 {
			SendClient25(client, true)
		}
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
