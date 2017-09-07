package main

import (
	"encoding/json"
	"github.com/ajstarks/svgo"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

var (
	width  = 100
	height = 20

	propMap = map[string]string{
		STAR: "stargazers_count",
		FORK: "forks_count",
	}

	clientId     = os.Getenv("github-client-id")
	clientSecret = os.Getenv("github-client-secret")
)

const (
	STAR = "star"
	FORK = "fork"
)

func main() {
	r := httprouter.New()
	r.GET("/info/:username/:reponame/:propname", repo)
	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func repo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "image/svg+xml")
	propname := ps.ByName("propname")
	githubResp := sendGitHubAPI(
		ps.ByName("username"),
		ps.ByName("reponame"),
	)
	f := githubResp[propMap[propname]].(float64)
	msg := strconv.FormatFloat(f, 'f', -1, 64)
	s := svg.New(w)
	s.Start(width, height)
	s.Text(0, 15, msg, "text-anchor:start;font-family:monospace;")
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
