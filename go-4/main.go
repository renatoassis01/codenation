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

type respositorio struct {
	Items []struct {
		Name        string `json:"full_name"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Stars       int    `json:"stargazers_count"`
	} `json:"items"`
}






type itemsJSON struct {
	Name        string
	Description string
	URL         string
	Stars       int
}



func writeFileJSON(caminho string, s []byte) {
	var identJSON bytes.Buffer
	_ = json.Indent(&identJSON, s, "", "\t")
	ioutil.WriteFile(caminho, identJSON.Bytes(), 0644)
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
	var items = make([]itemsJSON, 0, 10)
	for index := 0; index < len(repos.Items); index++ {
		i := itemsJSON{}
		i.Stars = repos.Items[index].Stars
		i.Description = repos.Items[index].Description
		i.Name = repos.Items[index].Name
		i.URL = repos.Items[index].URL
		items = append(items, i)
	}
	r, _ := json.Marshal(items)
	writeFileJSON("stars.json", r)
	return nil
}
