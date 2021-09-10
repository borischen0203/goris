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

	// "fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/dialog"
	// "fyne.io/fyne/v2/widget"
	// "github.com/spf13/cobra"
	"github.com/spf13/cobra"
	// "github.com/webview/webview"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "This command will pop up a window to preview Gopher image.",
	Long:  `This command will pop up a window to preview Gopher image.`,
	Run: func(cmd *cobra.Command, args []string) {
		var gopherName string
		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
			fmt.Println(gopherName)
			// previewGopher(gopherName)
			// viewGopher(gopherName)
		} else {
			fmt.Println("Error: empty input! :-(")
		}
	},
}

// func viewGopher(gopherName string) {
// 	myApp := app.New()

// 	myWindow := myApp.NewWindow("Gopher")

// 	// Main menu
// 	fileMenu := fyne.NewMenu("File",
// 		fyne.NewMenuItem("Quit", func() { myApp.Quit() }),
// 	)

// 	helpMenu := fyne.NewMenu("Help",
// 		fyne.NewMenuItem("About", func() {
// 			dialog.ShowCustom("About", "Close", container.NewVBox(
// 				widget.NewLabel("Welcome to Gopher, a simple Desktop app created in Go with Fyne."),
// 				widget.NewLabel("Version: v0.1"),
// 				widget.NewLabel("Author: Aurélie Vache"),
// 			), myWindow)
// 		}))
// 	mainMenu := fyne.NewMainMenu(
// 		fileMenu,
// 		helpMenu,
// 	)
// 	myWindow.SetMainMenu(mainMenu)

// 	// Define a welcome text centered
// 	text := canvas.NewText(gopherName, color.White) //gopherName
// 	text.Alignment = fyne.TextAlignCenter
// 	URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"
// 	// Define a Gopher image
// 	var resource, _ = fyne.LoadResourceFromURLString(URL)
// 	gopherImg := canvas.NewImageFromResource(resource)
// 	gopherImg.SetMinSize(fyne.Size{Width: 500, Height: 500}) // by default size is 0, 0

// 	// Display a vertical box containing text, image and button
// 	box := container.NewVBox(
// 		text,
// 		gopherImg,
// 	)

// 	// Display our content
// 	myWindow.SetContent(box)

// 	// Close the App when Escape key is pressed
// 	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

// 		if keyEvent.Name == fyne.KeyEscape {
// 			myApp.Quit()
// 		}
// 	})

// 	// Show window and run app
// 	myWindow.ShowAndRun()

// }

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
