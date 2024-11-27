package tools

import (
	"fmt"
	"log"

	tb "github.com/tucnak/telebot"
)

func UniversalListHandler[T any](
	bot *tb.Bot,
	m *tb.Message,
	fetchEntitiesFunc func() ([]T, error),
	formatEntityFunc func(T, int) string,
	batchSize int,
	storage map[int]int,
) {
	log.Printf("User %d requested entity list\n", m.Sender.ID)

	entities, err := fetchEntitiesFunc()
	if err != nil {
		log.Printf("Error fetching entities: %v\n", err)
		bot.Send(m.Sender, "Не удалось получить список. Попробуйте позже.")
		return
	}

	if len(entities) == 0 {
		bot.Send(m.Sender, "Список пуст.")
		return
	}

	storage[m.Sender.ID] = 0
	SendBatch(bot, m.Sender, storage[m.Sender.ID], batchSize, entities, formatEntityFunc, storage)
}

func SendBatch[T any](
	bot *tb.Bot,
	user *tb.User,
	index int,
	batchSize int,
	entities []T,
	formatEntityFunc func(T, int) string,
	storage map[int]int,
) {
	total := len(entities)
	endIndex := index + batchSize
	if endIndex > total {
		endIndex = total
	}

	msgText := fmt.Sprintf("Список (стр. %d из %d):\n", (index/batchSize)+1, (total+batchSize-1)/batchSize)

}
