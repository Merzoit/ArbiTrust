package handlers

import (
	"arbbot/api"
	"arbbot/menu"
	"arbbot/structures"
	"log"

	tb "github.com/tucnak/telebot"
)

func StartHandler(bot *tb.Bot, m *tb.Message) {
	user := structures.User{
		TID:  int64(m.Sender.ID),
		Name: m.Sender.Username,
	}

	if err := api.AddUserAPI(user); err != nil {
		log.Println("Error adding user:", err)
	}

	menu.ShowMainMenu(bot, m)
}
