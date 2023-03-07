package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/notes"
)

func main() {
	fmt.Print("[*] server-address:")
	address := Input()
	fmt.Print("[*] token(filename):")
	fn := Input()
	token := Readfile(fn)

	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig(address, token))
	client.LogLevel(logrus.DebugLevel)

	content := ""
	for {
		content = Input()
		response, err := client.Notes().Create(notes.CreateRequest{
			Text:       core.NewString(content),
			Visibility: models.VisibilityHome,
		})
		if err != nil {
			log.Printf("[Notes] Error happened: %s", err)
			return
		}

		if content == "" {
			return
		}

		log.Println(response.CreatedNote.ID)
	}
}

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func Readfile(fn string) string {
	f, err := os.Open(fn)
	if err != nil {
		log.Println("error")
	}
	defer f.Close()

	c, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("error")
	}
	return string(c)
}
