package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/masatana/go-textdistance"
)

type Answer struct {
	Message     string
	Answer      string
	Attachments string
}

var (
	DBAnswer []*Answer
	port     string = "8080"
)

func main() {
	//log.Println(os.Args)
	for i, arg := range os.Args {
		if arg == "-port" {
			port = os.Args[i+1]
			break
		}
	}

	// Echo instance
	e := echo.New()

	loadDatabase()

	// Middleware
	/*
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
	*/

	// Routes
	e.GET("/", hello)
	e.GET("/getMessage/:message", getMessage)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Пусто!")
}

func getMessage(c echo.Context) error {
	mess := c.Param("message")
	mess = strings.ReplaceAll(mess, "%20", " ")

	lastSeq := float64(0)
	answer := ""
	attach := ""

	for _, phase := range DBAnswer {
		seq := textdistance.JaroWinklerDistance(mess, phase.Message)
		if seq > lastSeq {
			answer = phase.Answer
			attach = phase.Attachments
			lastSeq = seq
		}
	}

	return c.String(http.StatusOK, "{\"answer\": \""+answer+"\", \"attachments\": \""+attach+"\"}")
}

func loadDatabase() {
	content, err := ioutil.ReadFile("answer_database.bin")
	if err != nil {
		log.Fatal(err)
	}

	s := string(content)
	base := strings.Split(s, "\n")

	for i, phase := range base {
		phaseArr := strings.Split(phase, "\\")
		DBAnswer = append(DBAnswer, new(Answer))
		DBAnswer[i].Message = phaseArr[0]
		attach := ""
		answ := ""
		words := strings.Split(phaseArr[1], " ")
		for _, word := range words {
			if (strings.HasPrefix(word, "https://vk.com/") || strings.HasPrefix(word, "http://vk.com/")) && len(word) > 15 {
				link := ""

				if strings.HasPrefix(word, "https://vk.com/") {
					link = strings.Replace(word, "https://vk.com/", "", 1)
				} else if strings.HasPrefix(word, "http://vk.com/") {
					link = strings.Replace(word, "http://vk.com/", "", 1)
				}

				if attach == "" {
					attach = link
				} else {
					attach = strings.Join(
						[]string{
							attach,
							link,
						},
						",",
					)
				}
			} else {
				if answ == "" {
					answ = word
				} else {
					answ = strings.Join(
						[]string{
							answ,
							word,
						},
						" ",
					)
				}
			}
		}
		DBAnswer[i].Answer = answ
		DBAnswer[i].Attachments = attach
	}

}
