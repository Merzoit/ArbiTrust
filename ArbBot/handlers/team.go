package handlers

import (
	"arbbot/api"
	"arbbot/constants"
	"arbbot/structures"
	"arbbot/tools"
	"fmt"
	"log"
	"strconv"

	tb "github.com/tucnak/telebot"
)

var teamData = make(map[int]*structures.Team)
var step = make(map[int]int)
var currentTeamIndex = make(map[int]int)
var teamEntities []structures.Team

type StepHandler func(bot *tb.Bot, m *tb.Message, userID int) error

var steps = map[int]StepHandler{
	0: handleStep0,
	1: handleStep1,
	2: handleStep2,
	3: handleStep3,
	4: handleStep4,
	5: handleStep5,
	6: handleStep6,
	7: handleStep7,
	8: handleStep8,
	9: handleStep9,
}

func CollectTeamData(bot *tb.Bot, m *tb.Message) {
	userID := m.Sender.ID
	currentStep := step[userID]

	log.Printf("Collect team data for user: %v", userID)

	handler, exist := steps[currentStep]
	if !exist {
		log.Printf(constants.ErrHandlerNoFound, currentStep)
		return
	}

	if err := handler(bot, m, userID); err != nil {
		log.Printf(constants.ErrHandlerStep, currentStep, userID, err)
		bot.Send(m.Sender, "Произошла ошибка, попробуйте снова.")
		return
	}

	step[userID]++
	if currentStep == len(steps)-1 {
		delete(step, userID)
		delete(teamData, userID)
		log.Println("Team creation completed")
	}
}

func handleStep0(bot *tb.Bot, m *tb.Message, userID int) error {
	teamData[userID] = &structures.Team{
		Name:       m.Text,
		Owner:      int64(m.Sender.ID),
		IsVerified: false,
	}
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 2*\n*Введите контактную информацию:*")
}

func handleStep1(bot *tb.Bot, m *tb.Message, userID int) error {
	teamData[userID].Contacts = m.Text
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 3*\n*Введите тематику команды:*")
}

func handleStep2(bot *tb.Bot, m *tb.Message, userID int) error {
	teamData[userID].Topic = m.Text
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 4*\n*Введите минимальную стоимость подписчика:*")
}

func handleStep3(bot *tb.Bot, m *tb.Message, userID int) error {
	price, err := tools.ParsePrice(m.Text)
	if err != nil {
		return err
	}

	teamData[userID].MinSubPrice = price
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 5*\n*Введите максимальную стоимость подписчикка:*")
}

func handleStep4(bot *tb.Bot, m *tb.Message, userID int) error {
	price, err := tools.ParsePrice(m.Text)
	if err != nil {
		return err
	}

	teamData[userID].MaxSubPrice = price
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 6*\n*Введите описание команды:*")
}

func handleStep5(bot *tb.Bot, m *tb.Message, userID int) error {
	teamData[userID].Description = m.Text
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 7*\n*Введите ссылку на бота:*")
}

func handleStep6(bot *tb.Bot, m *tb.Message, userID int) error {
	teamData[userID].BotLink = m.Text
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 8*\n*Введите размер команды:*")
}

func handleStep7(bot *tb.Bot, m *tb.Message, userID int) error {
	teamSize, err := strconv.Atoi(m.Text)
	if err != nil || teamSize <= 0 {
		return fmt.Errorf("invalid team size")
	}

	teamData[userID].TeamSize = teamSize
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 9*\n*Введите количество спонсоров в вашей команде:*")
}

func handleStep8(bot *tb.Bot, m *tb.Message, userID int) error {
	sponsorCount, err := strconv.Atoi(m.Text)
	if err != nil || sponsorCount <= 0 {
		return fmt.Errorf("invalid sponsor count")
	}

	teamData[userID].SponsorCount = sponsorCount
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 10*\n*Введите минимальную сумму для вывода средств:*")
}

func handleStep9(bot *tb.Bot, m *tb.Message, userID int) error {
	minWithdrawal, err := strconv.Atoi(m.Text)
	if err != nil || minWithdrawal <= 0 {
		log.Printf("СОСИ: %v", minWithdrawal)
		return fmt.Errorf("invalid withdrawal amount")
	}

	teamData[userID].MinWithdrawalAmount = minWithdrawal
	log.Printf("Submiting team data for user: %v", userID)

	if err := api.AddTeamToAPI(*teamData[userID]); err != nil {
		return fmt.Errorf(constants.LogErrorAddingTeam, err)
	}

	return tools.SendMessage(
		bot, m.Sender, "*Ваша заявка успешно отправлена на модерацию*")
}

func TeamListHandler(bot *tb.Bot, m *tb.Message) {
	entities, err := api.FetchTeamsAPI()
	if err != nil {
		bot.Send(m.Sender, "Ошибка загрузки списка команд.")
		return
	}
	teamEntities = entities

	tools.UniversalListHandler[structures.Team](
		bot,
		m,
		api.FetchTeamsAPI,
		api.FormatTeam,
		5,
		currentTeamIndex,
	)
}

func TeamNavigatorHandler(bot *tb.Bot, m *tb.Message) {
	tools.HandleNavigation(
		bot,
		m,
		currentTeamIndex,
		5,
		teamEntities,
		api.FormatTeam,
	)
}
