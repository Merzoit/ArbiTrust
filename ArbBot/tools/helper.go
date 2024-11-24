package tools

import (
	"arbbot/constants"
	"fmt"
	"log"
	"strconv"

	tb "github.com/tucnak/telebot"
)

func SendMessage(bot *tb.Bot, user *tb.User, text string) error {
	_, err := bot.Send(user, text, &tb.SendOptions{
		ParseMode: tb.ModeMarkdown,
	})

	if err != nil {
		log.Printf(constants.ErrSendingStep, err)
		return err
	}
	return nil
}

func ParsePrice(text string) (float64, error) {
	price, err := strconv.ParseFloat(text, 64)
	if err != nil || price < 0 {
		return 0, fmt.Errorf("invalid price")
	}

	return price, nil
}
