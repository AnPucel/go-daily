package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

type Weather struct {
	loc *Location
}

type Location struct {
	Name      string
	Region    string
	Country   string
	Latitude  float32
	Longitude float32
}

func NewsAction(c *cli.Context) error {
	fmt.Printf("Here")
	return nil
}

func WeatherAction(c *cli.Context) error {
	fmt.Printf("Weather Action")
	loc := os.Getenv("LOCATION")
	key := os.Getenv("WEATHER_API_KEY")

	client := &http.Client{}

	req, _ := http.NewRequest("GET", "http://api.weatherapi.com/v1/current.json", nil)
	q := req.URL.Query()
	q.Add("q", loc)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("key", key)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return err
	}

	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)

	var weather_data map[string]interface{}
	err = json.Unmarshal(resp_body, &weather_data)

	// Parse JSON

	return nil
}

// Make a daily fact, trivia, news, weather
func main() {
	app := cli.NewApp()
	app.Name = "Daily Riddles, Trivia, Facts & Weather"
	app.Usage = "Shows a new riddle, fact or trivia question daily"

	app.Commands = []cli.Command{
		{
			Name:      "news",
			ShortName: "news",
			Action:    NewsAction,
		},
		{
			Name:      "weather",
			ShortName: "weather",
			Action:    WeatherAction,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
