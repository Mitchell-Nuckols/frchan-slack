package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Mitchell-Nuckols/splat"
)

var (
	signingSecret string
	tbaToken      string
	client        = &http.Client{}
)

func init() {
	flag.StringVar(&signingSecret, "slack", "", "App signing secret")
	flag.StringVar(&tbaToken, "tba", "", "TBA token")
	flag.Parse()
}

func main() {
	app := splat.New(signingSecret)

	app.RegisterCommand("first", func(p *splat.SlashRequest) {

		response := new(splat.Response)

		args := strings.Split(p.Text, " ")

		switch args[0] {
		case "team":
			teamCommand(args[1:], response)
		case "help":
			helpCommand(response)
		default:
			unknownCommand(response)
		}

		p.Write(response)
	})

	app.RegisterAction("test_action", "test", func(p *splat.ActionPayload) {
		log.Println("recieved action")
	})

	log.Fatal(app.Open(":3000", "/slack-first"))
}

func teamCommand(args []string, r *splat.Response) {
	var info TBATeam

	req, err := http.NewRequest("GET", "http://www.thebluealliance.com/api/v3/team/"+args[0], nil)
	if err != nil {
		r.Text = err.Error()
		return
	}

	req.Header.Set("X-TBA-Auth-Key", tbaToken)

	res, err := client.Do(req)
	if err != nil {
		r.Text = err.Error()
		return
	}

	err = json.NewDecoder(res.Body).Decode(&info)
	if err != nil {
		r.Text = err.Error()
		return
	}

	if (info == TBATeam{}) {
		r.Text = "Team not found"
		return
	}

	r.Attachments = splat.Attachments{
		{
			Title:     info.Nickname + " (" + strconv.Itoa(info.TeamNumber) + ")",
			TitleLink: info.Website,
			Text:      info.Name,
			Fields: splat.Fields{
				{
					Title: "Motto",
					Value: info.Motto,
					Short: false,
				},
				{
					Title: "Rookie Year",
					Value: strconv.Itoa(info.RookieYear),
					Short: true,
				},
				{
					Title: "Location",
					Value: info.City + ", " + info.SateProv,
					Short: true,
				},
			},
			Color:  "#00704A",
			Footer: "Info provided by <https://www.thebluealliance.com/|The Blue Alliance>",
		},
	}
}

func helpCommand(r *splat.Response) {
	r.Attachments = splat.Attachments{
		{
			Title: "Help",
			Text:  "This bot provides FRC information using <https://www.thebluealliance.com/|The Blue Alliance>",
			Color: "#00704A",
		},
	}
}

func unknownCommand(r *splat.Response) {
	r.Text = "That command doesn't exist!"
}
