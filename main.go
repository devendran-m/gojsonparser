package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const url = "http://universities.hipolabs.com/search?country=United+States"

func main() {

	content := "This is a file created by this Go Program"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close()
	defer readFile("./fromString.txt")

	resp, err := http.Get(url)
	checkError(err)
	fmt.Printf("Response Type: %T\n", resp)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	webcontent := string(bytes)
	fmt.Print(webcontent)
	fmt.Println("\n")

	univs := uniFromJson(webcontent)
	for _, uni := range univs {
		fmt.Println(uni.Country, " : ", uni.Name)
	}

}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file >> ", string(data))
}

func uniFromJson(webcontent string) []Uni {
	univs := make([]Uni, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(webcontent))
	token, err := decoder.Token()
	checkError(err)
	fmt.Print("Token >> \n", token)
	fmt.Print(err)

	var uni Uni
	for decoder.More() {
		err := decoder.Decode(&uni)
		checkError(err)
		univs = append(univs, uni)
	}
	return univs
}

type Uni struct {
	Country, Name string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
