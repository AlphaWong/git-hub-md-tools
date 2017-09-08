package main

import (
	"encoding/json"
	"fmt"
	"github.com/ajstarks/svgo"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	width  = 140
	height = 60

	propMap = map[string]string{
		STAR: "stargazers_count",
		FORK: "forks_count",
	}

	clientId     = os.Getenv("GITHUB_CLIENT_ID")
	clientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
)

const (
	STAR = "star"
	FORK = "fork"
)

func main() {
	r := httprouter.New()
	r.GET("/info/:username/:reponame/", repo)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func repo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache")

	githubResp := sendGitHubAPI(
		ps.ByName("username"),
		ps.ByName("reponame"),
	)
	msg := strconv.FormatFloat(githubResp[propMap[STAR]].(float64), 'f', -1, 64)

	s := svg.New(w)
	s.Start(width, height)
	msg = fmt.Sprintf("Star: %s", msg)
	s.Text(4, 14, msg, "text-anchor:start;font-family:monospace;font-size: larger;")

	msg = strconv.FormatFloat(githubResp[propMap[FORK]].(float64), 'f', -1, 64)

	msg = fmt.Sprintf("Fork: %s", msg)
	s.Text(4, 28, msg, "text-anchor:start;font-family:monospace;font-size: larger;")

	s.End()
}

func sendGitHubAPI(username string, reponame string) (githubResp map[string]interface{}) {
	resp, err := http.Get("https://api.github.com/repos/" + username + "/" + reponame + "?client_id=" + clientId + "&client_secret=" + clientSecret)
	if err != nil {
		log.Fatal("githubResp:", err)
	}

	if resp != nil {
		json.NewDecoder(resp.Body).Decode(&githubResp)
		defer resp.Body.Close()
		return
	}
	return
}
