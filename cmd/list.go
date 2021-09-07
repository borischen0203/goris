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

//Package cmd provides command for using
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

//GopherJson struct
type GopherJson struct {
	Sha  string `json:"sha"`
	URL  string `json:"url"`
	Tree []struct {
		Path string `json:"path"`
		Mode string `json:"mode"`
		Type string `json:"type"`
		Sha  string `json:"sha"`
		Size int    `json:"size"`
		URL  string `json:"url"`
	} `json:"tree"`
	Truncated bool `json:"truncated"`
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List gopher png list",
	Long:  `List gopher png list`,
	Run: func(cmd *cobra.Command, args []string) {
		URL := "https://api.github.com/repos/scraly/gophers/git/trees/main?recursive=1"

		fmt.Println("[Try to get Gopher list...]")

		// Get the data
		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			jsonDataFromHttp, err := ioutil.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			var gopher GopherJson
			err = json.Unmarshal([]byte(jsonDataFromHttp), &gopher)
			if err != nil {
				panic(err)
			}
			for i := 0; i < len(gopher.Tree); i++ {
				if strings.Contains(gopher.Tree[i].Path, ".png") {
					result := strings.ReplaceAll(gopher.Tree[i].Path, ".png", "")
					fmt.Println(result)
				}
			}
		} else {
			fmt.Println("Error: list not exists! :-(")
		}

	},
}

// func getList() {

// }

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
