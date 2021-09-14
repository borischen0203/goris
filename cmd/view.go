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
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"

	// "github.com/webview/webview"
	"github.com/zserge/lorca"
)

//test commit
//test commit2
// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "This command will pop up a window to preview Gopher image.",
	Long:  `This command will pop up a window to preview Gopher image.`,
	Run: func(cmd *cobra.Command, args []string) {
		var gopherName string
		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
			found, _ := gopherExist(gopherName)
			if found {
				fmt.Println("Pop-up " + gopherName + " window!")
				// previewGopher(gopherName)
				gopherPreview(gopherName)
			} else {
				fmt.Println("Error: " + gopherName + " not exists!")
			}

		}
	},
}

// check gopher image exist or not
func gopherExist(gopherName string) (bool, error) {
	URL := "https://raw.githubusercontent.com/scraly/gophers/main/" + gopherName + ".png"
	response, err := http.Get(URL)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		return true, nil
	}
	return false, nil
}

// Pop-up a window for preview gopher image
func gopherPreview(gopherName string) {
	// Create UI with data URI
	ui, _ := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
	 <head><title>`+gopherName+`</title></head>
	 <style>
	 .gopher {
		width:100%;
		height:100%;
		text-align: center;
	}
	img {
		max-width: auto;
        max-height:100%;
        width: auto;
        object-fit:container;
	}
	</style>
	 <body>
	 <div class ="gopher">
	 <img src="https://raw.githubusercontent.com/scraly/gophers/main/`+gopherName+`.png" alt="friends"/>
	 </div>
	 </body>
	</html>
	`), "", 800, 600)
	defer ui.Close()
	<-ui.Done()
}

//webview lib version: has issue with deploy
// func previewGopher(gopherName string) {
// 	URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"
// 	debug := true
// 	w := webview.New(debug)
// 	defer w.Destroy()
// 	w.SetTitle(gopherName)
// 	w.SetSize(800, 600, webview.HintNone)
// 	w.Navigate(URL)
// 	w.Run()
// }

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
