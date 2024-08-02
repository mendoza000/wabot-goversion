package client

import (
	"fmt"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"golang.org/x/net/context"
	"strings"
	"wabot/pkgs/db"
	"wabot/pkgs/utils"
)

func SendClient25(client *whatsmeow.Client, reminder bool) {
	clients := db.GetClients("25")
	accounts := db.GetAccounts()

	for _, c := range clients {
		if c.Phone == "" {
			continue
		}
		service := utils.GetService(c.AccountService)
		account := utils.FilterAccountByID(c.AccountID, accounts)

		msg := utils.GetMessage(reminder, c, service, account)

		phoneNumber := strings.ReplaceAll(c.Phone, "+", "")

		_, err := client.SendMessage(context.Background(), types.JID{
			User:   phoneNumber,
			Server: types.DefaultUserServer,
		}, &waProto.Message{
			Conversation: msg,
		})

		if err != nil {
			panic(err)
		}

		fmt.Println("Message sent to: " + c.Name)
	}

	fmt.Println("Messages sent to all clients! 🚀🚀🚀")
}
