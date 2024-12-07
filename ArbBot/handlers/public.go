package handlers

import (
	"arbbot/api"
	"arbbot/constants"
	"arbbot/structures"
	"arbbot/tools"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tb "github.com/tucnak/telebot"
)

var publicData = make(map[int]*structures.Public)
var pubStep = make(map[int]int)
var currentPublicIndex = make(map[int]int)
var publicEntities []structures.Public

type PublicStepHandler func(bot *tb.Bot, m *tb.Message, userID int) error

var pubSteps = map[int]PublicStepHandler{
	0:  pubHandleStep0,
	1:  pubHandleStep1,
	2:  pubHandleStep2,
	3:  pubHandleStep3,
	4:  pubHandleStep4,
	5:  pubHandleStep5,
	6:  pubHandleStep6,
	7:  pubHandleStep7,
	8:  pubHandleStep8,
	9:  pubHandleStep9,
	10: pubHandleStep10,
}

func pubHandleStep0(bot *tb.Bot, m *tb.Message, userID int) error {
	publicData[userID] = &structures.Public{
		Name:       m.Text,
		Owner:      int64(m.Sender.ID),
		IsScammer:  false,
		IsVerified: false,
		RegDate:    time.Now(),
	}
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 2*\n*Введите тэг паблика.*")
}

func pubHandleStep1(bot *tb.Bot, m *tb.Message, userID int) error {
	publicData[userID].Tag = m.Text
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 3*\n*Введите предпочитаемый способ связи:*")
}

func pubHandleStep2(bot *tb.Bot, m *tb.Message, userID int) error {
	publicData[userID].Contacts = m.Text
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 4*\n*Введите тематику паблика:*")
}

func pubHandleStep3(bot *tb.Bot, m *tb.Message, userID int) error {
	publicData[userID].Topic = m.Text
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 5*\n*Введите цену за подписчика:*")
}

func pubHandleStep4(bot *tb.Bot, m *tb.Message, userID int) error {
	price, err := tools.ParsePrice(m.Text)
	if err != nil || price < 0 {
		return err
	}

	publicData[userID].SubcriberPrice = price
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 6*\n*Введите стоимость рекламного поста в паблике:*")
}

func pubHandleStep5(bot *tb.Bot, m *tb.Message, userID int) error {
	price, err := tools.ParsePrice(m.Text)
	if err != nil || price < 0 {
		return err
	}

	publicData[userID].AdPrice = price
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 7*\n*Желаете встать на ОП? (да/нет):*")
}

func pubHandleStep6(bot *tb.Bot, m *tb.Message, userID int) error {
	answer := strings.ToLower(m.Text)

	switch answer {
	case "да":
		publicData[userID].WantsOP = true
	case "нет":
		publicData[userID].WantsOP = false
	default:
		return fmt.Errorf("invalid answer")
	}

	return tools.SendMessage(
		bot, m.Sender, "*Шаг 8*\n*Введите описание паблика:*")
}

func pubHandleStep7(bot *tb.Bot, m *tb.Message, userID int) error {
	publicData[userID].Description = m.Text
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 9*\n*Ваш паблик продаётся? (да/нет):*")
}

func pubHandleStep8(bot *tb.Bot, m *tb.Message, userID int) error {
	answer := strings.ToLower(m.Text)

	switch answer {
	case "да":
		publicData[userID].IsSelling = true
	case "нет":
		publicData[userID].IsSelling = false
	default:
		return fmt.Errorf("invalid answer")
	}

	return tools.SendMessage(
		bot, m.Sender, "*Шаг 10*\n*Введите стоимость паблика (либо укажите 0):*")
}

func pubHandleStep9(bot *tb.Bot, m *tb.Message, userID int) error {
	price, err := tools.ParsePrice(m.Text)
	if err != nil || price < 0 {
		return fmt.Errorf("invalid price")
	}

	publicData[userID].SalePrice = price
	return tools.SendMessage(
		bot, m.Sender, "*Шаг 11*\n*Введите количество ежемесечного актива паблика:*")
}

func pubHandleStep10(bot *tb.Bot, m *tb.Message, userID int) error {
	active, err := strconv.Atoi(m.Text)
	if err != nil || active <= 0 {
		return fmt.Errorf("invalid count")
	}

	publicData[userID].MonthlyUsers = active
	log.Printf("Submiting public data for user: %v", userID)

	if err := api.AddPublicToAPI(*publicData[userID]); err != nil {
		return fmt.Errorf(constants.LogErrorAddingPublic, err)
	}

	return tools.SendMessage(
		bot, m.Sender, "*Ваша заявка успешно отправлена на модерацию*")
}

func PublicListHandler(bot *tb.Bot, m *tb.Message) {
	entities, err := api.FetchPublicsAPI()
	if err != nil {
		bot.Send(m.Sender, "Ошибка загрузки списка пабликов.")
		return
	}

	publicEntities = entities

	tools.UniversalListHandler[structures.Public](
		bot,
		m,
		api.FetchPublicsAPI,
		api.FormatPublic,
		5,
		currentPublicIndex,
	)
}

func PublicavigatorHandler(bot *tb.Bot, m *tb.Message) {
	tools.HandleNavigation(
		bot,
		m,
		currentPublicIndex,
		5,
		publicEntities,
		api.FormatPublic,
	)
}
