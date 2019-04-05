package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const TOKEN = "48c62d18afa2e20201ff6c7f6130857ba7f6a5f1"

type answer struct {
	NumeroCasas         byte   `json:"numero_casas"`
	Token               string `json:"token"`
	Cifrado             string `json:"cifrado"`
	Decifrado           string `json:"decifrado"`
	ResumoCriptografico string `json:"resumo_criptografico"`
}

func writeFileJSON(caminho string, s []byte) {
	var identJSON bytes.Buffer
	_ = json.Indent(&identJSON, s, "", "\t")
	ioutil.WriteFile(caminho, identJSON.Bytes(), 0644)
}

func escapaLetras(n int) bool {

	escape := map[int]bool{
		32: true,
		46: true,
		48: true,
		49: true,
		50: true,
		51: true,
		52: true,
		53: true,
		54: true,
		55: true,
		56: true,
		57: true,
	}
	if _, ok := escape[n]; ok {
		return true
	} else {
		return false
	}

}
func criptografa(s string, salto byte) string {
	v := []byte(strings.ToLower(s))
	for i := 0; i < len(v); i++ {
		if ok := escapaLetras(int(v[i])); ok {

		} else {
			v[i] = v[i] + salto
		}

	}

	return string(v)

}
func descriptografa(s string, salto byte) string {
	v := []byte(strings.ToLower(s))
	for i := 0; i < len(v); i++ {
		if ok := escapaLetras(int(v[i])); ok {
		} else {
			v[i] = v[i] - salto
		}

	}

	return string(v)

}
func PostDados(arquivo, url string) {
	file, _ := os.Open(arquivo)
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("answer", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()
	r, _ := http.NewRequest("POST", url, body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		fmt.Printf("Error : %s", err)
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string([]byte(body)))
	}
}

func get() answer {
	url := fmt.Sprintf("https://api.codenation.dev/v1/challenge/dev-ps/generate-data?token=%s", TOKEN)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var answer = answer{}
	json.Unmarshal(body, &answer)
	r, _ := json.Marshal(answer)
	writeFileJSON("answer.json", r)
	return answer

}

func main() {
	j := get()
	j.Decifrado = descriptografa(j.Cifrado, j.NumeroCasas)
	h := sha1.New()
	h.Write([]byte(j.Decifrado))
	sha := h.Sum(nil)
	j.ResumoCriptografico = hex.EncodeToString(sha)
	r, _ := json.Marshal(j)
	writeFileJSON("answer.json", r)
	url := fmt.Sprintf("https://api.codenation.dev/v1/challenge/dev-ps/submit-solution?token=%s", TOKEN)
	PostDados("answer.json", url)

}
