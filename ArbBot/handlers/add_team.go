package handlers

import (
	"arbbot/api"
	"arbbot/structures"
	"log"
	"strconv"

	tb "github.com/tucnak/telebot"
)

var teamData = make(map[int]*structures.Team)
var step = make(map[int]int)

func AddTeamHandler(bot *tb.Bot, m *tb.Message) {
	teamData[m.Sender.ID] = &structures.Team{IsVerified: false}
	step[m.Sender.ID] = 0
	bot.Send(m.Sender, "Введите название команды:")
}

func CollectTeamData(bot *tb.Bot, m *tb.Message) {
	userID := m.Sender.ID
	currentStep := step[userID]

	switch currentStep {
	case 0:
		teamData[userID].Name = m.Text
		teamData[userID].Owner = "q"
		step[userID]++
		bot.Send(m.Sender, "Введите контактную информацию: ")
	case 1:
		teamData[userID].Contacts = m.Text
		step[userID]++
		bot.Send(m.Sender, "Введите тематику команды: ")
	case 2:
		teamData[userID].Topic = m.Text
		step[userID]++
		bot.Send(m.Sender, "Введите минимальную цену за подписчика: ")
	case 3:
		price, err := strconv.ParseFloat(m.Text, 64)
		if err != nil {
			bot.Send(m.Sender, "Пожалуйста, введите корректное числовое значение для минимальной цены.")
			return
		}
		teamData[userID].MinSubPrice = price
		step[userID]++
		bot.Send(m.Sender, "Введите максимальную цену за подписчика: ")
	case 4:
		price, err := strconv.ParseFloat(m.Text, 64)
		if err != nil {
			bot.Send(m.Sender, "Пожалуйста, введите корректное числовое значение для максимальной цены.")
			return
		}
		teamData[userID].MaxSubPrice = price
		step[userID]++
		bot.Send(m.Sender, "Введите описание команды: ")
	case 5:
		teamData[userID].Description = m.Text
		step[userID]++
		bot.Send(m.Sender, "Введите ссылку на бота: ")
	case 6:
		teamData[userID].BotLink = m.Text
		step[userID]++
		bot.Send(m.Sender, "Введите размер команды (число участников): ")
	case 7:
		teamSize, err := strconv.Atoi(m.Text)
		if err != nil || teamSize <= 0 {
			bot.Send(m.Sender, "Пожалуйста, введите корректное числовое значение для размера команды.")
			return
		}
		teamData[userID].TeamSize = teamSize
		step[userID]++
		bot.Send(m.Sender, "Введите количество спонсоров: ")
	case 8:
		sponsorCount, err := strconv.Atoi(m.Text)
		if err != nil || sponsorCount < 0 {
			bot.Send(m.Sender, "Пожалуйста, введите корректное числовое значение для количества спонсоров.")
			return
		}
		teamData[userID].SponsorCount = sponsorCount
		step[userID]++
		bot.Send(m.Sender, "Введите минимальную сумму для вывода средств: ")
	case 9:
		minWithdrawal, err := strconv.Atoi(m.Text)
		if err != nil || minWithdrawal < 0 {
			bot.Send(m.Sender, "Пожалуйста, введите корректное числовое значение для минимальной суммы для вывода.")
			return
		}
		teamData[userID].MinWithdrawalAmount = minWithdrawal

		if err := api.AddTeamToAPI(*teamData[userID]); err != nil {
			log.Printf("Ошибка при добавлении команды: %v", teamData)
			bot.Send(m.Sender, "Ошибка при добавлении команды, попробуйте позже.")
		} else {
			bot.Send(m.Sender, "Команда успешно отправлена на модерацию.")
		}

		delete(teamData, userID)
		delete(step, userID)
	}
}
