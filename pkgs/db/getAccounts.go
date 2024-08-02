package db

import (
	"encoding/json"
	"fmt"
	"github.com/supabase-community/supabase-go"
	"wabot/models"
)

func GetAccounts() []models.Account {
	client, err := supabase.NewClient("https://labouqvrnsksnthjoaig.supabase.co", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImxhYm91cXZybnNrc250aGpvYWlnIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTY5Mzg1MDQxMywiZXhwIjoyMDA5NDI2NDEzfQ.62I_t3c2U7qjt-Wo5dma9UD-FsZJKBWZFT-ilSZbOKE", nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}

	var acList []models.Account

	data, _, err := client.From("accounts").Select("*", "exact", false).Execute()

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}

	if err := json.Unmarshal(data, &acList); err != nil {
		fmt.Println("error", err)
		panic(err)
	}

	return acList
}
