package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mymmsc/gox/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"wabot/utils"
)

func connectClient(opt int) {
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
	clientLog := waLog.Stdout("Client", "INFO", true)
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
		c := color.New(color.FgHiBlack, color.Bold).Add(color.BgWhite).SprintFunc()
		green := color.New(color.FgHiBlack, color.Bold, color.BgGreen).SprintFunc()
		yellow := color.New(color.FgYellow, color.Bold).SprintFunc()

		if opt == 1 {
			groups := utils.GetGroupsToSend()
			for _, group := range groups {
				utils.SendMessage(client, group.JID)
				fmt.Printf(" %s "+group.Name+"\n", c(" Send message to: "))
			}
			fmt.Printf("\n\n%s press %s to exit.", green(" Messages sent successfully! "),
				yellow("Ctrl+C"))
		}
		if opt == 2 {
			utils.GetGroup(client)
		}
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}

func main() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		return
	}

	c := color.New(color.FgHiBlack, color.Bold).Add(color.BgWhite)
	cm := color.New(color.FgWhite, color.Bold).Add(color.BgWhite)
	yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
	underline := color.New(color.Underline, color.Bold).SprintFunc()

	cm.Println("\n    Welcome to Wabot - Go version    ")
	c.Println("    Welcome to Wabot - Go version    ")
	cm.Println("    Welcome to Wabot - Go version    ")

	fmt.Printf("  %s mendoza000\n\n", yellow("Developed by:"))

	fmt.Println("  Select your option:")

	fmt.Printf("  %s Send messages\n", yellow("1."))
	fmt.Printf("  %s Get groups\n\n", yellow("2."))

	var opcion int
	fmt.Printf("  %s ", underline("Type the number of your option:"))

	_, err = fmt.Scanln(&opcion)
	if err != nil {
		fmt.Println("Error al leer la opción")
	}

	switch opcion {
	case 1:
		connectClient(1)
		break
	case 2:
		connectClient(2)
		break
	default:
		fmt.Println("Opción inválida")
	}
}
