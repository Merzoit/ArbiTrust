package api

import (
	"arbbot/structures"
	"encoding/json"
	"fmt"
	"net/http"

	tb "github.com/tucnak/telebot"
)

var Teams []structures.Team

func AddTeamToAPI(team structures.Team) error {
	return PostToAPI("http://localhost:8080/api/team", team)
}

func FetchTeamsAPI() error {
	apiUrl := "http://localhost:8080/api/allteams/"

	resp, err := http.Get(apiUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch teams %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(&Teams)
}

// Обработчик отображения команд, batchSize - указывает на количество отображаемых команд за один раз.
func SendTeam(bot *tb.Bot, user *tb.User, index int, batchSize int, teams []structures.Team) {
	endIndex := index + batchSize

	if endIndex > len(teams) {
		endIndex = len(teams)
	}

	msgText := "Список команд:\n"

	for i := index; i < endIndex; i++ {
		team := teams[i]
		msgText += fmt.Sprintf(
			"%d. Название: %s \nВладелец: %s\nОписание: %s\n\n",
			i+1, team.Name, team.Owner, team.Description,
		)
	}

	menu := &tb.ReplyMarkup{}
	btnNext := tb.InlineButton{Text: "Вперед", Data: "next"}
	btnPrev := tb.InlineButton{Text: "Назад", Data: "prev"}
	btnExit := tb.InlineButton{Text: "Выход", Data: "exit"}

	var buttons [][]tb.InlineButton

	if index > 0 {
		buttons = append(buttons, []tb.InlineButton{btnPrev})
	}

	if index < len(Teams)-1 {
		if len(buttons) == 0 {
			buttons = append(buttons, []tb.InlineButton{btnNext})
		} else {
			buttons[0] = append(buttons[0], btnNext)
		}
	}

	buttons = append(buttons, []tb.InlineButton{btnExit})
	menu.InlineKeyboard = buttons

	if _, err := bot.Send(user, msgText, menu); err != nil {
		fmt.Printf("Error sending message to user %d: %v", user.ID, err)
	}
}
