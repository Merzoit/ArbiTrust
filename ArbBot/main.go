package main

import (
	"arbbot/api"
	"arbbot/handlers"
	"log"
	"time"

	tb "github.com/tucnak/telebot"
)

const batchSize = 3

func main() {
	botToken := "7676099516:AAEIhkwxeE1CDm6_kBO8eKghrQ0Lft7zE9M"
	bot, err := tb.NewBot(tb.Settings{
		Token:  botToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	/*bot.Handle("/start", func(m *tb.Message) {
		user := structures.User{
			TID:  int64(m.Sender.ID),
			Name: m.Sender.Username,
		}

		err := api.AddUserAPI(user)
		if err != nil {
			bot.Send(m.Sender, "Неудалось добавить пользователя")
			showMainMenu(bot, m)
			log.Println("Error adding user", err)
			return
		}

		bot.Send(m.Sender, "Успешно добавлен")
		showMainMenu(bot, m)
	})

	bot.Handle(&tb.ReplyButton{Text: "Список команд"}, func(m *tb.Message) {
		err := api.FetchTeamsAPI()
		if err != nil {
			bot.Send(m.Sender, "Не удалось получить список команд")
			return
		}

		currentIndex[m.Sender.ID] = 0
		api.SendTeam(bot, m.Sender, currentIndex[m.Sender.ID])
	})

	bot.Handle(tb.OnCallback, func(c *tb.Callback) {
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
	})*/

	bot.Handle("/start", func(m *tb.Message) {
		handlers.StartHandler(bot, m)
	})

	bot.Handle(&tb.ReplyButton{Text: "Список команд"}, func(m *tb.Message) {
		handlers.TeamListHandler(bot, m, batchSize, api.Teams)
	})

	bot.Handle(&tb.ReplyButton{Text: "Добавить команду"}, func(m *tb.Message) {
		handlers.AddTeamHandler(bot, m)
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		handlers.CollectTeamData(bot, m)
	})

	bot.Handle(tb.OnCallback, func(c *tb.Callback) {
		handlers.HandleNavigation(bot, c, 3, api.Teams)
	})
	bot.Start()
	log.Panicln("OK!")
}
