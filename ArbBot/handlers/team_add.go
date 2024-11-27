package handlers

import (
	"arbbot/api"
	"arbbot/structures"
	"fmt"
	"strconv"

	tb "github.com/tucnak/telebot"
)

var teamDatas = make(map[int]*structures.Team)
var stepss = make(map[int]int)

func AddTeamHandler(bot *tb.Bot, m *tb.Message) {
	teamData[m.Sender.ID] = &structures.Team{IsVerified: false}
	step[m.Sender.ID] = 0
	bot.Send(m.Sender,
		"*Шаг 1*\n*Введите название команды:*\n\n_Пример: My Team\nНазвание команды может содержать_ *буквы, цифры, символы*_. Максимальное количество символов -_ *155*.",
		&tb.SendOptions{
			ParseMode: tb.ModeMarkdown,
		})
}

func CollectTeamDatas(bot *tb.Bot, m *tb.Message) {
	userID := m.Sender.ID
	currentStep := step[userID]

	switch currentStep {
	case 0:
		teamData[userID].Name = m.Text
		teamData[userID].Owner = m.Sender.Username
		step[userID]++
		bot.Send(m.Sender,
			"*Шаг 2*\n*Введите контактную информацию:*\n\n_Пример: @mynickname, my@gmail.com\nВведите предпочтительные способы связи. Это может быть_ *номер телефона, почта, иные способы*_. Несколько способов вводите через запятую._ ",
			&tb.SendOptions{
				ParseMode: tb.ModeMarkdown,
			})
	case 1:
		teamData[userID].Contacts = m.Text
		step[userID]++
		bot.Send(m.Sender,
			"*Шаг 3*\n*Введите тематику команды:*\n\n_Пример: аниме, заработок, криптовалюта\nНе нужно вносить много тем, достаточно одной._ *Тематика команды будет назначена модератером* _после успешной верефикации._ ",
			&tb.SendOptions{
				ParseMode: tb.ModeMarkdown,
			})
	case 2:
		teamData[userID].Topic = m.Text
		step[userID]++
		bot.Send(m.Sender,
			"*Шаг 4*\n*Введите минимальную стоимость подписчика:*\n\n_Пример: 0.4, 1, 1.2\nВведите корректное целочисленное либо дробное значение. Вводите только одну цифру,_ *минимум который можно получить за подписчика.*",
			&tb.SendOptions{
				ParseMode: tb.ModeMarkdown,
			})
	case 3:
		price, err := strconv.ParseFloat(m.Text, 64)
		if err != nil {
			bot.Send(m.Sender, "Пожалуйста, введите корректное числовое значение для минимальной цены.")
			return
		}
		teamData[userID].MinSubPrice = price
		step[userID]++
		bot.Send(m.Sender,
			"*Шаг 5*\n*Введите максимальную стоимость подписчика:*\n\n_Пример: 0.4, 1, 1.2\nВведите корректное целочисленное либо дробное значение. Вводите только одну цифру,_ *максимум который можно получить за подписчика.*",
			&tb.SendOptions{
				ParseMode: tb.ModeMarkdown,
			})
	case 4:
		price, err := strconv.ParseFloat(m.Text, 64)
		if err != nil {
			bot.Send(m.Sender, "Пожалуйста, введите корректное числовое значение для максимальной цены.")
			return
		}
		teamData[userID].MaxSubPrice = price
		step[userID]++
		bot.Send(m.Sender,
			"*Шаг 6\nВведите описание команды:*\n\n_Пример: Это пример вашего описания\nВведите описание команды._ *Проявите себя, либо оставьте поле пустым.*",
			&tb.SendOptions{
				ParseMode: tb.ModeMarkdown,
			})
	case 5:
		teamData[userID].Description = m.Text
		step[userID]++
		bot.Send(m.Sender,
			"*Шаг 7\nВведите ссылку на бота:*\n\n_Пример: @myBot\nВведите ссылку на вашего бота в формате @myBot._ *Вводите только одну ссылку.*",
			&tb.SendOptions{
				ParseMode: tb.ModeMarkdown,
			})
	case 6:
		teamData[userID].BotLink = m.Text
		step[userID]++
		bot.Send(m.Sender,
			"*Шаг 8\nВведите размер команды:*\n\n_Пример: 4, 10, 100\n_*Введите целое, положительное число,* _котороt будет указывать количество участников в вашей команде_",
			&tb.SendOptions{
				ParseMode: tb.ModeMarkdown,
			})
	case 7:
		teamSize, err := strconv.Atoi(m.Text)
		if err != nil || teamSize <= 0 {
			bot.Send(m.Sender, "Пожалуйста, введите корректное числовое значение для размера команды.")
			return
		}
		teamData[userID].TeamSize = teamSize
		step[userID]++
		bot.Send(m.Sender,
			"*Шаг 8\nВведите количество спонсоров в вашей команде:*\n\n_Пример: 3,5,7\n_*Введите целое, положительное число,* _которое будет указывать количество спонсоров в вашей команде_",
			&tb.SendOptions{
				ParseMode: tb.ModeMarkdown,
			})
	case 8:
		sponsorCount, err := strconv.Atoi(m.Text)
		if err != nil || sponsorCount < 0 {
			bot.Send(m.Sender, "Пожалуйста, введите корректное числовое значение для количества спонсоров.")
			return
		}
		teamData[userID].SponsorCount = sponsorCount
		step[userID]++
		bot.Send(m.Sender,
			"*Шаг 8\nВведите минимальную сумму для вывода средств:*\n\n_Пример: 100, 1000, 10000\n_*Введите целое, положительное число,* _которое будет указывать на минимальную сумму доступную для вывода средств_",
			&tb.SendOptions{
				ParseMode: tb.ModeMarkdown,
			})
	case 9:
		minWithdrawal, err := strconv.Atoi(m.Text)
		if err != nil || minWithdrawal < 0 {
			bot.Send(m.Sender, "Пожалуйста, введите корректное числовое значение для минимальной суммы для вывода.")
			return
		}
		teamData[userID].MinWithdrawalAmount = minWithdrawal

		if err := api.AddTeamToAPI(*teamData[userID]); err != nil {
			fmt.Printf("Ошибка при добавлении команды: %v", teamData)
			bot.Send(m.Sender, "Ошибка при добавлении команды, попробуйте позже.")
		} else {
			bot.Send(m.Sender,
				"*Ваша заявка на добавление команды успешно отправлена на модерацию.*\n_Не подавайте новых заявок на эту команду, пока не получите результат текущей проверки. Среднее время проверки заявки 24 часа._",
				&tb.SendOptions{
					ParseMode: tb.ModeMarkdown,
				})
		}

		delete(teamData, userID)
		delete(step, userID)
	}
}
