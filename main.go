package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// DockerPayload IoT buttom event
type DockerPayload struct {
	CallbackURL string `json:"callback_url"`
	PushData    struct {
		Images   []string `json:"images"`
		PushedAt int      `json:"pushed_at"`
		Pusher   string   `json:"pusher"`
		Tag      string   `json:"tag"`
	} `json:"push_data"`
	Repository struct {
		CommentCount    int    `json:"comment_count"`
		DateCreated     int    `json:"date_created"`
		Description     string `json:"description"`
		Dockerfile      string `json:"dockerfile"`
		FullDescription string `json:"full_description"`
		IsOfficial      bool   `json:"is_official"`
		IsPrivate       bool   `json:"is_private"`
		IsTrusted       bool   `json:"is_trusted"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		Owner           string `json:"owner"`
		RepoName        string `json:"repo_name"`
		RepoURL         string `json:"repo_url"`
		StarCount       int    `json:"star_count"`
		Status          string `json:"status"`
	} `json:"repository"`
}

func main() {

	http.HandleFunc("/hook", slackHook)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// func slackHook(payload DockerPayload) {
func slackHook(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var payload DockerPayload
	err := decoder.Decode(&payload)

	if err != nil {
		panic(err)
	}

	url := os.Getenv("SLACK_HOOK")

	now := time.Now()
	ts := now.Unix()
	timestamp := strconv.FormatInt(ts, 10)

	var jsonStr = []byte(`{
		"attachments": [
			{
				"color": "#36a64f",
				"pretext": "DockerHub automated build status",
				"author_name": "` + payload.PushData.Pusher + `",
				"title": "` + payload.Repository.RepoName + `",
				"title_link": "` + payload.Repository.RepoURL + `",
				"text": "Docker build finished with tag: ` + payload.PushData.Tag + `",
				"footer": "DockerHub Go hook",
				"footer_icon": "https://platform.slack-edge.com/img/default_application_icon.png",
				"ts": ` + timestamp + `,
			}
		]
	}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// print response
	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
}
