package main

import (
	"log"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/notes"
)

func main() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://misskey.io", "20ef5DpyUy4RoTCMP3H2h7qX5hCRtET5"))
	client.LogLevel(logrus.DebugLevel)

	response, err := client.Notes().Create(notes.CreateRequest{
		Text:       core.NewString("test"),
		Visibility: models.VisibilityHome,
		Poll: &notes.Poll{
			Choices: []string{"a", "b", "c"},
		},
	})
	if err != nil {
		log.Printf("[Notes] Error happened: %s", err)

		return
	}

	log.Println(response.CreatedNote.ID)
}
