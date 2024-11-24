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
		log.Fatalf("Failed to initialize bot: %v", err)
	}

	bot.Handle("/start", func(m *tb.Message) {
		log.Printf("User %d issued /start command\n", m.Sender.ID)
		handlers.StartHandler(bot, m)
	})

	bot.Handle(&tb.ReplyButton{Text: "Список команд"}, func(m *tb.Message) {
		log.Printf("User %d requested team list\n", m.Sender.ID)
		if err := api.FetchTeamsAPI(); err != nil {
			log.Printf("Error fetching teams: %v\n", err)
			bot.Send(m.Sender, "Не удалось получить список команд. Попробуйте позже.")
			return
		}
		handlers.TeamListHandler(bot, m, batchSize, api.Teams)
	})

	bot.Handle(&tb.ReplyButton{Text: "Добавить команду"}, func(m *tb.Message) {
		log.Printf("User %d requested to add a team\n", m.Sender.ID)
		handlers.AddTeamHandler(bot, m)
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		log.Printf("User %d sent text: %s\n", m.Sender.ID, m.Text)
		handlers.CollectTeamData(bot, m)
	})

	bot.Handle(tb.OnCallback, func(c *tb.Callback) {
		log.Printf("User %d triggered callback with data: %s\n", c.Sender.ID, c.Data)
		handlers.HandleNavigation(bot, c, batchSize, api.Teams)
	})

	log.Println("Bot is running...")
	bot.Start()
}
