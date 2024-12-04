package handlers

import (
	"arbbot/api"
	"arbbot/constants"
	"arbbot/structures"
	"log"

	tb "github.com/tucnak/telebot"
)

var currentIndex = make(map[int]int)

func TeamListHandlers(bot *tb.Bot, m *tb.Message, batchSize int, teams []structures.Team) {
	log.Printf("Displaying team list to user %v", m.Sender.ID)

	if len(teams) == 0 {
		if _, err := bot.Send(m.Sender, "Список команд пуст"); err != nil {
			log.Printf(constants.ErrSendingTeamList, err)
		}
		return
	}

	currentIndex[m.Sender.ID] = 0
	api.SendTeam(bot, m.Sender, currentIndex[m.Sender.ID], batchSize, teams)
}

func HandleNavigation(bot *tb.Bot, c *tb.Callback, batchSize int, teams []structures.Team) {
	log.Printf("HandleNavigation process..")
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
		if _, err := bot.Send(c.Sender, "Выход из списка команд"); err != nil {
			log.Printf(constants.ErrNavigationHandler, err)
		}
		delete(currentIndex, userID)
	}
	bot.Respond(c)
}
