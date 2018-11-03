package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	_ = githubStars("go")
}

type items struct {
	Name        string `json:"full_name"`
	Description string `json:"description"`
	URL         string `json:"html_url"`
	Stars       int    `json:"stargazers_count"`
}
type respositorio struct {
	Items []items `json:"items"`
}

func writeFileJSON(caminho string, s []byte) {
	var identJSON bytes.Buffer
	_ = json.Indent(&identJSON, s, "", "\t")
	ioutil.WriteFile(caminho, identJSON.Bytes(), 0644)
}

func (r *items) MarshalJSON() ([]byte, error) {
	//http://choly.ca/post/go-json-marshalling/
	return json.Marshal(&struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Stars       int    `json:"stars"`
	}{
		Name:        r.Name,
		Description: r.Description,
		URL:         r.URL,
		Stars:       r.Stars,
	})
}

func githubStars(lang string) error {
	query := fmt.Sprintf("https://api.github.com/search/repositories?page=1&per_page=10&q=language:%s&sort=stars&order=desc", lang)
	res, err := http.Get(query)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var repos = respositorio{}
	json.Unmarshal(body, &repos)
	r, _ := json.Marshal(repos.Items)
	writeFileJSON("stars.json", r)
	return nil
}
