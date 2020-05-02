package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// TODO: rewrite basically all of this so it's pretty and makes sense
// this is powered by (now 3AM) code
// there is also a bug when a country is longer than 1 word (like United States)
type response struct {
	Global  global    `json:"Global"`
	Country []summary `json:"Countries"`
}

type global struct {
	TotalConfirmed int `json:"TotalConfirmed"`
	TotalDeaths    int `json:"TotalDeaths"`
	TotalRecovered int `json:"TotalRecovered"`
}

type summary struct {
	Name           string `json:"Country"`
	Code           string `json:"CountryCode"`
	TotalConfirmed int    `json:"TotalConfirmed"`
	TotalDeaths    int    `json:"TotalDeaths"`
	TotalRecovered int    `json:"TotalRecovered"`
}

var responseObject response

func init() {
	NewCommand("covid", false, covidstatus).SetHelp("get info on covid stats by country. search by country code, or leave blank for global stats." +
		" please reference this list for valid country codes: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements").Add()
}

func covidstatus(s *discordgo.Session, m *discordgo.MessageCreate, msgList []string) {
	if len(msgList) < 2 || msgList[1] == "global" {
		s.ChannelMessageSend(m.ChannelID, "no argument found, showing global stats.")
		_, r := getStats(s, m, "global")

		em := NewEmbed().
			SetTitle("covid stats for the world").
			AddField("total confirmed cases", strconv.Itoa(r.Global.TotalConfirmed)).
			AddField("total confirmed deaths", strconv.Itoa(r.Global.TotalDeaths)).
			AddField("total confirmed recoveries", strconv.Itoa(r.Global.TotalRecovered)).
			SetFooter("stats provided by Johns Hopkins CSSE").MessageEmbed

		s.ChannelMessageSendEmbed(m.ChannelID, em)
		return
	}

	if len(msgList) == 2 {
		i, r := getStats(s, m, msgList[1])
		if i == 0 {
			return
		}

		em := NewEmbed().
			SetTitle("covid stats for "+r.Country[i].Name).
			AddField("total confirmed cases", strconv.Itoa(r.Country[i].TotalConfirmed)).
			AddField("total confirmed deaths", strconv.Itoa(r.Country[i].TotalDeaths)).
			AddField("total confirmed recoveries", strconv.Itoa(r.Country[i].TotalRecovered)).
			SetFooter("stats provided by Johns Hopkins CSSE").MessageEmbed

		s.ChannelMessageSendEmbed(m.ChannelID, em)
		return
	}
}

func getStats(s *discordgo.Session, m *discordgo.MessageCreate, arg string) (int, response) {
	r, err := http.Get("https://api.covid19api.com/summary")
	if err != nil {
		fmt.Println("error:", err)
	}

	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &responseObject)

	if arg == "global" {
		return 0, responseObject
	}

	i := GetIndexOf(responseObject.Country, arg)
	if i == 0 {
		s.ChannelMessageSend(m.ChannelID, "couldnt get stats. make sure the name/country code is correct.")
		return 0, responseObject
	}
	return i, responseObject
}
