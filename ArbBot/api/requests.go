package api

import (
	"arbbot/structures"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	tb "github.com/tucnak/telebot"
)

var (
	Teams []structures.Team
)

/*func AddUserAPI(user structures.User) error {
	apiUrl := "http://localhost:8080/api/users"

	userData, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to encode userdata: %v", err)
	}

	resp, err := http.Post(apiUrl, "aplication/json", bytes.NewBuffer(userData))
	if err != nil {
		return fmt.Errorf("failed to send request to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("API returned status: %v", resp.Status)
	}

	return nil
}*/

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

func SendTeam(bot *tb.Bot, user *tb.User, index int) {
	if index < 0 || index >= len(Teams) {
		log.Printf("Invalid index %d for teams array of size %d", index, len(Teams))
		return
	}

	team := Teams[index]
	msgTxt := fmt.Sprintf("Команда: %s\nВладелец: %s", team.Name, team.Owner)

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

	if _, err := bot.Send(user, msgTxt, menu); err != nil {
		log.Printf("Error sending message to user %d: %v", user.ID, err)
	}
}
