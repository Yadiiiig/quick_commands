package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	search := strings.Join(os.Args[1:], " ")

	if search == "" {
		fmt.Println("Please search for something.")
		os.Exit(1)
	}

	fmt.Printf("\nYou searched for:\n%s\n\n", search)
	search = strings.Replace(search, " ", "+", -1)

	url := fmt.Sprintf("https://api.stackexchange.com/2.2/search/advanced?order=desc&sort=relevance&q=%s&site=stackoverflow", search)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("You've used special characters or you have achieved your rate limit.")
		os.Exit(1)
	}

	respPayload := payload{}
	err = json.NewDecoder(resp.Body).Decode(&respPayload)
	if err != nil {
		fmt.Println("Please submit this error as an issue.")
		fmt.Println(err)
		os.Exit(1)
	}

	for i := range respPayload.Items {
		if respPayload.Items[i].IsAnswered {
			fmt.Printf("\033]8;;%s\033\\%s\033]8;;\033\\\n", respPayload.Items[i].Link, string(respPayload.Items[i].Title))
			fmt.Println("─────────────────────────────────────────────────────────────────────────────────────────────────────────")
		}
	}

}

type payload struct {
	Items []struct {
		AcceptedAnswerID int64  `json:"accepted_answer_id"`
		AnswerCount      int64  `json:"answer_count"`
		ContentLicense   string `json:"content_license"`
		CreationDate     int64  `json:"creation_date"`
		IsAnswered       bool   `json:"is_answered"`
		LastActivityDate int64  `json:"last_activity_date"`
		Link             string `json:"link"`
		Owner            struct {
			AcceptRate   int64  `json:"accept_rate"`
			DisplayName  string `json:"display_name"`
			Link         string `json:"link"`
			ProfileImage string `json:"profile_image"`
			Reputation   int64  `json:"reputation"`
			UserID       int64  `json:"user_id"`
			UserType     string `json:"user_type"`
		} `json:"owner"`
		QuestionID int64    `json:"question_id"`
		Score      int64    `json:"score"`
		Tags       []string `json:"tags"`
		Title      string   `json:"title"`
		ViewCount  int64    `json:"view_count"`
	} `json:"items"`
}
