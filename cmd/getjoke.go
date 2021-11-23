/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// getjokeCmd represents the getjoke command
var getjokeCmd = &cobra.Command{
	Use:   "getjoke",
	Short: "Get Joke can you a random joke.",
	Long:  `Get Joke is randome joke getter from internet built using go.`,
	Run: func(cmd *cobra.Command, args []string) {
		getJoke()
	},
}

func init() {
	rootCmd.AddCommand(getjokeCmd)
}

type Joke struct {
	ID     string `json:"id`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}


func getJoke(){
 url := "https://icanhazdadjoke.com/"
 respBytes := getJokeFromApi(url)
 joke := Joke{}

 if err := json.Unmarshal(respBytes, &joke); err != nil {
	fmt.Printf("Could not unmarshal reponseBytes. %v", err)
 }

 fmt.Println(string(joke.Joke))
}

func getJokeFromApi(baseApi string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseApi,
		nil,
	)

	if err != nil {
		log.Printf("Could not request a dadjoke. %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "clit CLI (https://github.com/techieasif/clit)")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes

}
