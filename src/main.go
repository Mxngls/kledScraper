package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

// Initialize used variables
var Words []Result
var idArr []string

// Custom func that creates JSON from every data type and every rune encoding
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func main() {

	started := time.Now()

	// Initiate the http client
	client, err := createClient()
	if err != nil {
		panic(err)
	}

	// Read the word ids
	data, _ := ioutil.ReadFile("dict/wordIDs.json")
	err = json.Unmarshal(data, &idArr)
	if err != nil {
		panic(err)
	}

	// Open the file where to save our dictionary entries
	// If no such file exits yet create one
	f, err := os.OpenFile("dict/dict_FULL.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Create an initial slice to append to
	initArr := &[]Result{}

	// Set the number of goroutines that should be run
	start := 4200
	end := start + 100

	// Initiate Channel
	c := make(chan Result)

	fmt.Printf("\rFetched and parsed 0 from %d entries; running...", len(idArr))

	// Loop over all the ids and write them to file
	for {

		// Iinitialize the slice for the current loop
		var arr = &[]Result{}

		// Set the end point of the last loop the the number of word ids
		if end > len(idArr) {
			end = len(idArr)
		}

		// Start looping
		// In order to provide enough time to fetch a response wait for 50ms after
		// a new goroutine is run
		for i := start; i < end; i++ {
			id := string(idArr[i])
			time.Sleep(time.Millisecond * 20)
			go getView(i, id, c, client)

			fmt.Printf("\rFetched and parsed %d from %d entries; running...", i+1, len(idArr))
		}

		// Store the fetched and parsed response bodies to file
		for i := start; i < end; i++ {
			*arr = append(*arr, <-c)
		}

		// Sort the intermediate array according to its elements index
		sort.SliceStable(*arr, func(i, j int) bool {
			return (*arr)[i].Alpha < (*arr)[j].Alpha
		})

		// Concatenate the initial slice and the slice with the results of the current loop
		*initArr = append(*initArr, *arr...)

		// Break out of the loop when the end point is equal to the number of word ids
		if end == len(idArr) {

			// Write the initial slice to file
			JSON, _ := JSONMarshal(*initArr)
			_, err := f.Write(JSON)
			if err != nil {
				log.Fatal(err)
			}
			f.Close()

			fmt.Printf("\rWrote %d from %d entries to file: 'dict_FULL.JSON'; finished", len(*initArr), len(idArr))

			finished := time.Since(started)

			fmt.Printf("Process took %d", finished)
			break
		}

		// Increment the start point
		start += 100
		end += 100
	}
}
