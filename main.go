package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var Token = "....-..........-..........-..........-......"

type Response struct {
	OK      bool   `json:"ok"`
	Error   string `json:"error"`
	Members []struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Deleted         bool   `json:"deleted"`
		Restricted      bool   `json:"is_restricted"`
		UltraRestricted bool   `json:"is_ultra_restricted"`
		Bot             bool   `json:"is_bot"`
		TwoFactor       bool   `json:"has_2fa"`
	} `json:"members"`
}

func init() {
	flag.StringVar(&Token, "token", Token, "see https://api.slack.com/web#authentication")
}

func main() {
	var response Response

	flag.Parse()

	resp, err := http.Get(fmt.Sprintf("https://slack.com/api/users.list?token=%s", url.QueryEscape(Token)))

	if err != nil {
		log.Fatal("Error fetching https://slack.com/api/users.list: " + err.Error())
	}

	var decoder = json.NewDecoder(resp.Body)

	if err = decoder.Decode(&response); err != nil {
		log.Fatal(err)
	}

	if response.OK == false {
		log.Fatal(response.Error)
	}

	var non2fa = []string{}

	for _, user := range response.Members {
		if user.Deleted {
			continue
		}
		if user.Bot {
			continue
		}
		if user.TwoFactor {
			continue
		}
		non2fa = append(non2fa, user.Name)
	}

	if len(non2fa) < 1 {
		fmt.Println("all non-bot users have 2FA enabled!")
	} else {
		fmt.Println("non-bots without 2FA:", "@"+strings.Join(non2fa, ", @"))
		os.Exit(2)
	}
}
