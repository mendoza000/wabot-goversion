package db

import (
	"encoding/json"
	"fmt"
	"github.com/supabase-community/supabase-go"
	"wabot/models"
)

func GetClients(day string) []models.Client {
	client, err := supabase.NewClient("https://labouqvrnsksnthjoaig.supabase.co", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImxhYm91cXZybnNrc250aGpvYWlnIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTY5Mzg1MDQxMywiZXhwIjoyMDA5NDI2NDEzfQ.62I_t3c2U7qjt-Wo5dma9UD-FsZJKBWZFT-ilSZbOKE", nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}

	var cList []models.Client

	if day == "25" {
		data, _, err := client.From("clients").Select("*", "exact", false).Gte("day", "16").Eq("paid", "false").Execute()

		if err != nil {
			fmt.Println("error", err)
			panic(err)
		}

		if err := json.Unmarshal(data, &cList); err != nil {
			fmt.Println("error", err)
			panic(err)
		}

		return cList
	}

	data, _, err := client.From("clients").Select("*", "exact", false).Lte("day", day).Eq("paid", "false").Execute()

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}

	if err := json.Unmarshal(data, &cList); err != nil {
		fmt.Println("error", err)
		panic(err)
	}

	return cList
}
