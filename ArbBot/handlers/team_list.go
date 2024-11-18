package handlers

import (
	"arbbot/api"

	tb "github.com/tucnak/telebot"
)

var currentIndex = make(map[int]int)

func TeamListHandler(bot *tb.Bot, m *tb.Message) {
	
	err := api.FetchTeamsAPI()
	if err != nil {
		bot.Send(m.Sender, "Не удалось получить команд")
	}

	currentIndex[m.Sender.ID] = 0
	api.SendTeam(bot, m.Sender, currentIndex[m.Sender.ID])
}

func HandleNavigation(bot *tb.Bot, c *tb.Callback) {
	userID := c.Sender.ID
	switch c.Data {
	case "next":
		if currentIndex[userID] < len(api.Teams)-1 {
			currentIndex[userID]++
			api.SendTeam(bot, c.Sender, currentIndex[userID])
		}
	case "prev":
		if currentIndex[userID] > 0 {
			currentIndex[userID]--
			api.SendTeam(bot, c.Sender, currentIndex[userID])
		}
	case "exit":
		bot.Send(c.Sender, "Выход из списка команд")
		delete(currentIndex, userID)
	}
	bot.Respond(c)
}
