package utils

import "wabot/models"

func FilterAccountByID(id int, accounts []models.Account) models.Account {
	var account models.Account

	for _, acc := range accounts {
		if acc.ID == id {
			account = acc
		}
	}

	return account
}
