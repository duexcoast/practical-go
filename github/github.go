package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer cancel()
	name, numRepos, err := githubInfo(ctx, "tebeka")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("name: %q, numRepos: %d", name, numRepos)
}

func githubInfo(ctx context.Context, login string) (string, int, error) {
	// url := fmt.Sprintf("https://api.github.com/users/%s", login)
	url := "https://api.github.com/users/" + url.PathEscape(login)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	// resp, err := http.Get(url)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, errors.New(resp.Status)
	}

	var r struct { // anonymous struct
		Name     string
		NumRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.NumRepos, nil

}

type Reply struct {
	Name string
	// Public_Repos int
	NumRepos int `json:"public_repos"`
}
