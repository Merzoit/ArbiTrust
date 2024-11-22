package handlers

import (
	"arbbot/api"
	"arbbot/structures"

	tb "github.com/tucnak/telebot"
)

var currentIndex = make(map[int]int)

func TeamListHandler(bot *tb.Bot, m *tb.Message, batchSize int, teams []structures.Team) {

	if len(teams) == 0 {
		bot.Send(m.Sender, "Список команд пуст")
	}

	currentIndex[m.Sender.ID] = 0
	api.SendTeam(bot, m.Sender, currentIndex[m.Sender.ID], batchSize, teams)
}

func HandleNavigation(bot *tb.Bot, c *tb.Callback, batchSize int, teams []structures.Team) {
	userID := c.Sender.ID
	switch c.Data {
	case "next":
		if currentIndex[userID]+batchSize < len(teams)-1 {
			currentIndex[userID] += batchSize
			api.SendTeam(bot, c.Sender, currentIndex[userID], batchSize, teams)
		}
	case "prev":
		if currentIndex[userID]-batchSize >= 0 {
			currentIndex[userID] -= batchSize
			api.SendTeam(bot, c.Sender, currentIndex[userID], batchSize, teams)
		}
	case "exit":
		bot.Send(c.Sender, "Выход из списка команд")
		delete(currentIndex, userID)
	}
	bot.Respond(c)
}
