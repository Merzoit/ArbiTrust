package api

import (
	"arbbot/constants"
	"arbbot/structures"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Public []structures.Public

func AddPublicToAPI(public structures.Public) error {
	log.Printf("AddPublicToAPI called...")

	err := PostToAPI("http://localhost:8080/api/public", public)
	if err != nil {
		log.Printf(constants.LogErrorAddingPublic, err)
		return err
	}

	log.Printf("Public %v added successfully", public.ID)
	return nil
}

func FetchPublicsAPI() ([]structures.Public, error) {
	apiUrl := "http://localhost:8080/api/public/all"
	log.Printf("FetchPublicsAPI called..")

	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Printf(constants.ErrFetchingResponse, err)
		return nil, fmt.Errorf(constants.ErrFetchingResponse, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf(constants.ErrFetchingPublic, err)
		return nil, fmt.Errorf(constants.ErrFetchingPublic, err)
	}

	err = json.NewDecoder(resp.Body).Decode(&Public)
	if err != nil {
		log.Printf(constants.ErrDecodingResponse, err)
		return nil, fmt.Errorf(constants.ErrDecodingResponse, err)
	}

	log.Printf("Public fetched successfylly")
	return Public, nil
}

/*func SendPublic(bot *tb.Bot, user *tb.User, index int, batchSize int, publics []structures.Public) {
	log.Printf("Send team for user: %v", user.ID)

	endIndex := index + batchSize
	if endIndex > len(publics) {
		endIndex = len(publics)
	}

	msgText := "Список пабликов:\n"

	for i := index; i < endIndex; i++ {
		public := publics[i]
		msgText += fmt.Sprintf(
			"%d. Название: %s \nВладелец: %s\nОписание: %s\n\n",
			i+1, public.Name, public.Name, public.Description,
		)
	}

	menu := &tb.ReplyMarkup{}
	btnNext := tb
}*/

func FormatPublic(public structures.Public, index int) string {
	return fmt.Sprintf("%d. Название: %s\nВладелец: %s\nОписание: %s\n\n", index, public.Name, public.Name, public.Description)
}
