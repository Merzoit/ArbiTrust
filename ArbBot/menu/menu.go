package menu

import (
	tb "github.com/tucnak/telebot"
)

func ShowMainMenu(bot *tb.Bot, m *tb.Message) {
	menu := &tb.ReplyMarkup{}
	btnList := tb.ReplyButton{Text: "Список команд"}
	btnAddTeam := tb.ReplyButton{Text: "Добавить команду"}
	menu.ReplyKeyboard = [][]tb.ReplyButton{{btnList}, {btnAddTeam}}
	bot.Send(m.Sender, "Меню:", menu)
}
