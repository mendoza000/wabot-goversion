package utils

import (
	"fmt"
	"go.mau.fi/whatsmeow"
)

func GetGroup(client *whatsmeow.Client) {
	groups, err := client.GetJoinedGroups()

	if err != nil {
		panic(err)
	}
	for _, g := range groups {
		fmt.Println("----")
		fmt.Println(g.Name)
		fmt.Println(g.JID)

	}
}
