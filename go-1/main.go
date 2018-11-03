package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type arquivos struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func ParseToJson(v []arquivos) ([]byte, error) {

	s, error := json.Marshal(v)
	if error != nil {
		return []byte{}, error
	} else {
		return s, nil
	}
}

func WriteFileJson(caminho string, s []byte) {
	var identJSON bytes.Buffer
	_ = json.Indent(&identJSON, s, "", "\t")
	ioutil.WriteFile(caminho, identJSON.Bytes(), 0644)
}

func ScanDir(caminho string) ([]arquivos, error) {
	l := []arquivos{}
	z := "./"
	err := filepath.Walk(caminho, func(path string, info os.FileInfo, err error) error {
		f := arquivos{}
		if err != nil {
			return err
		} else {
			if !filepath.IsAbs(path) {
				dir, _ := filepath.Split(path)
				f.Path = z + dir
				f.Name = z + filepath.Base(path)

			} else {
				if path == "./" || path == "." {
					//pass
				} else {
					f.Path = z + path
					f.Name = path
				}
			}
			l = append(l, f)
		}
		return nil

	})
	return l, err
}

func jsonify(dir string) error {
	v, err := ScanDir(dir)
	if err != nil {
		return err
	} else {
		j, err := ParseToJson(v)
		if err != nil {
			return err
		} else {
			WriteFileJson("files.json", j)
		}
	}
	return nil
}

func main() {
	//@todo read dir name from stdin
	_ = jsonify("./")

}
