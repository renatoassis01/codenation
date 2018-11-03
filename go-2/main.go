package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"sort"
	"strconv"
)

func main() {

}

func loadData() ([][]string, error) {
	csvFile, _ := os.Open("data.csv")
	r := csv.NewReader(bufio.NewReader(csvFile))
	defer csvFile.Close()
	records, err := r.ReadAll()
	return records, err

}

//Quantas nacionalidades (coluna `nationality`) diferentes existem no arquivo?
func q1() (int, error) {
	nationality := make(map[string]int)
	records, error := loadData()
	for _, v := range records {
		n := string(v[14])
		if _, ok := nationality[n]; ok == false && n != "" {
			nationality[n] = 1
		}
	}

	return len(nationality) - 1 /** -1 ignorando o cabeçalho**/, error
}

//Quantos clubes (coluna `club`) diferentes existem no arquivo?
func q2() (int, error) {
	club := make(map[string]int)
	records, error := loadData()
	for _, v := range records {
		n := string(v[3])
		if _, ok := club[n]; ok == false && n != "" {
			club[n] = 1
		}
	}
	return len(club) - 1 /** -1 ignorando o cabeçalho**/, error
}

//Liste o primeiro nome dos 20 primeiros jogadores de acordo com a coluna `full_name`.
func q3() ([]string, error) {
	records, error := loadData()
	if error != nil {
		panic(error)

	}
	var j = make([]string, 0, 20)
	for index := 1; index <= 20; index++ {
		x := records[index][2]
		j = append(j, x)
	}
	return j, error
}

//Quem são os top 10 jogadores que ganham mais dinheiro (utilize as colunas `full_name` e `eur_wage`)?
func q4() ([]string, error) {
	records, error := loadData()
	if error != nil {
		panic(error)
	}
	sort.Slice(records, func(i, j int) bool {
		z, _ := strconv.ParseFloat(records[i][17], 32)
		l, _ := strconv.ParseFloat(records[j][17], 32)
		return z > l
	})
	var w = make([]string, 0, 10)
	for index := 1; index < 11; index++ {
		w = append(w, records[index][2])
	}
	return w, error
}

//Quem são os 10 jogadores mais velhos?
func q5() ([]string, error) {
	records, error := loadData()
	if error != nil {
		panic(error)
	}
	var w = make([]string, 0, 10)
	f := records[1:]
	sort.Slice(f, func(i, j int) bool {
		z, _ := strconv.ParseInt(f[i][6], 10, 32)
		l, _ := strconv.ParseInt(f[j][6], 10, 32)
		w = append(w, f[i][2])
		return z > l
	})

	return w[0:10], error
}

//Conte quantos jogadores existem por idade. Para isso, construa um mapa onde as chaves são as idades e os valores a contagem.
func q6() (map[int]int, error) {
	idades := make(map[int]int)
	records, error := loadData()
	for index := 1; index < len(records); index++ {
		if i, err := strconv.Atoi(records[index][6]); err == nil {
			if _, ok := idades[i]; ok == false {
				idades[i] = 1
			} else {
				idades[i]++
			}
		}

	}
	return idades, error
}
