package api

import (
	"arbbot/constants"
	"arbbot/structures"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	tb "github.com/tucnak/telebot"
)

var Teams []structures.Team

func AddTeamToAPI(team structures.Team) error {
	log.Printf(constants.CallAddTeam, team.ID)

	err := PostToAPI("http://localhost:8080/api/team", team)
	if err != nil {
		log.Printf(constants.LogErrorAddingTeam, err)
		return err
	}

	log.Printf(constants.LogTeamAddingSuccessfully, team.ID)
	return nil
}

func FetchTeamsAPI() error {
	apiUrl := "http://localhost:8080/api/allteams/"
	log.Printf(constants.CallFetchTeam, apiUrl)

	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Printf(constants.ErrFetchingResponse, err)
		return fmt.Errorf(constants.ErrFetchingResponse, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf(constants.ErrFetchingTeam, resp.Status)
		return fmt.Errorf(constants.ErrFetchingTeam, resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&Teams)
	if err != nil {
		log.Printf(constants.ErrDecodingResponse, err)
		return fmt.Errorf(constants.ErrDecodingResponse, err)
	}

	log.Println(constants.LogTeamFetchingSuccessfylly)
	return nil
}

func SendTeam(bot *tb.Bot, user *tb.User, index int, batchSize int, teams []structures.Team) {
	log.Printf(constants.CallShowTeamSender, user.ID)

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
		log.Printf(constants.ErrSendingMessage, user.ID)
	} else {
		log.Println(constants.LogTeamSendingSuccessfylly)
	}
}
