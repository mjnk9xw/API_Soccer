package main

import (
	"SoccerBeego/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Arr Models

type Arr struct {
	Tours   []*models.Tour   `json:"tours"`
	Rounds  []*models.Round  `json:"rounds"`
	Tables  []*models.Table  `json:"tables"`
	Players []*models.Player `json:"players"`
	Matchs  []*models.Match  `json:"matchs"`
}

// OpenFile doc file co chua chuoi json
// func OpenFile(filename string) {
// 	counts := make(map[string]int)
// 	data, _ := ioutil.ReadFile(filename)
// 	for _, line := range strings.Split(string(data), "\n") {
// 		counts[line]++
// 	}
// 	if filename == "tour.txt" {
// 		for line, n := range counts {
// 			if n >= 1 {
// 				Req("http://localhost:8080/api/v1/tour/", line)
// 			}
// 		}
// 	}
// 	if filename == "round.txt" {
// 		for line, n := range counts {
// 			if n >= 1 {
// 				Req("http://localhost:8080/api/v1/tour/1/round/", line)
// 			}
// 		}
// 	}
// 	if filename == "table.txt" {
// 		for line, n := range counts {
// 			if n >= 1 {
// 				Req("http://localhost:8080/api/v1/tour/1/round/1/table", line)
// 			}
// 		}
// 	}
// 	if filename == "player.txt" {
// 		for line, n := range counts {
// 			if n >= 1 {
// 				Req("http://localhost:8080/api/v1/tour/1/round/1/table/1/player", line)
// 			}
// 		}
// 	}
// 	if filename == "match.txt" {
// 		for line, n := range counts {
// 			if n >= 1 {
// 				Req("http://localhost:8080/api/v1/tour/1/round/1/table/1/match", line)
// 			}
// 		}
// 	}
// }

// Req Thuc hien Post cac chuoi JSON
func ReqTour(url string, js []*models.Tour) {
	for i := 0; i < len(js); i++ {
		//var jsonStr = []byte(js[i])
		//req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		t := new(bytes.Buffer)
		json.NewEncoder(t).Encode(js[i])
		req, err := http.NewRequest("POST", url, t)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println("response Status:", resp.Status)
	}
}
func ReqRound(url string, js []*models.Round) {
	for i := 0; i < len(js); i++ {
		t := new(bytes.Buffer)
		json.NewEncoder(t).Encode(js[i])
		req, err := http.NewRequest("POST", url, t)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println("response Status:", resp.Status)
	}
}
func ReqTable(url string, js []*models.Table) {
	for i := 0; i < len(js); i++ {
		t := new(bytes.Buffer)
		json.NewEncoder(t).Encode(js[i])
		req, err := http.NewRequest("POST", url, t)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println("response Status:", resp.Status)
	}
}
func ReqPlayer(url string, js []*models.Player) {
	for i := 0; i < len(js); i++ {
		//var jsonStr = []byte(js[i])
		t := new(bytes.Buffer)
		json.NewEncoder(t).Encode(js[i])
		req, err := http.NewRequest("POST", url, t)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println("response Status:", resp.Status)
	}
}
func ReqMatch(url string, js []*models.Match) {
	for i := 0; i < len(js); i++ {
		t := new(bytes.Buffer)
		json.NewEncoder(t).Encode(js[i])
		req, err := http.NewRequest("POST", url, t)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println("response Status:", resp.Status)
	}
}

func main() {
	// OpenFile("tour.txt")
	// OpenFile("round.txt")
	// OpenFile("table.txt")
	// OpenFile("player.txt")
	// OpenFile("match.txt")
	var arr Arr
	// jsonFile, _ := os.Open("tour.json")
	// defer jsonFile.Close()
	//jsonTours, _ := ioutil.ReadAll(jsonFile)
	// jsonParse := json.NewDecoder(jsonFile)
	// jsonParse.Decode(&arr)
	// fmt.Println(arr)
	jsonTours, _ := ioutil.ReadFile("tour.json")
	json.Unmarshal(jsonTours, &arr.Tours)
	jsonRounds, _ := ioutil.ReadFile("round.json")
	json.Unmarshal(jsonRounds, &arr.Rounds)
	jsonTables, _ := ioutil.ReadFile("table.json")
	json.Unmarshal(jsonTables, &arr.Tables)
	jsonPlayers, _ := ioutil.ReadFile("player.json")
	json.Unmarshal(jsonPlayers, &arr.Players)
	jsonMatchs, _ := ioutil.ReadFile("match.json")
	json.Unmarshal(jsonMatchs, &arr.Matchs)

	ReqTour("http://localhost:8080/api/v1/tour/", arr.Tours)
	ReqRound("http://localhost:8080/api/v1/tour/1/round/", arr.Rounds)
	ReqTable("http://localhost:8080/api/v1/tour/1/round/1/table", arr.Tables)
	ReqPlayer("http://localhost:8080/api/v1/tour/1/round/1/table/1/player", arr.Players)
	ReqMatch("http://localhost:8080/api/v1/tour/1/round/1/table/1/match", arr.Matchs)

}
