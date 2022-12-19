package main

import (
	"log"
	"os"

	"github.com/anpucel/go-daily/actions"
	"github.com/urfave/cli"
)

// Make a daily fact, trivia, news, weather
func main() {
	app := cli.NewApp()
	app.Name = "Daily Riddles, Trivia, Facts & Weather"
	app.Usage = "Shows a new riddle, fact or trivia question daily"

	app.Commands = []cli.Command{
		{
			Name:      "news",
			ShortName: "news",
			Action:    actions.NewsAction,
		},
		{
			Name:      "weather",
			ShortName: "weather",
			Action:    actions.WeatherAction,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
