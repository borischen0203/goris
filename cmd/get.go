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
	"embed"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

//go:embed gophers
var embedGopherFiles embed.FS

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub repository in order to return the desired Gopher.`,
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		gopherName := "dr-who.png"
		l := len(args)

		switch l {
		case 1:
			gopherName = args[0]
			path, _ := getDefaultPath()
			err := getGopherImage(path, gopherName)
			if err != nil {
				return err
			}
			return nil
		case 2:
			if args[0] == "link" {
				gopherName = args[1]
				err := getGopherLink(gopherName)
				if err != nil {
					return err
				}
			}
		}
		return nil
	},
}

func getGopherLink(gopherName string) error {
	fmt.Println("Try to Generate gopher link...")
	URL := "https://api.github.com/repos/scraly/gophers/contents/" + gopherName + ".png"
	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode == 200 || response.StatusCode == 403 {
		link := "https://raw.githubusercontent.com/scraly/gophers/main/" + gopherName + ".png"
		result := "<img src=" + "\"" + link + "\"" + " alt=" + "\"" + gopherName + "\"" + ">"
		fmt.Println(result)
		fmt.Println(gopherName + "! I choose you! Paste above link in the readme!")
	} else {
		fmt.Println("Error: " + gopherName + " not exists! :-(")
	}
	return nil
}

func getLW(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	m, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	g := m.Bounds()

	// Get height and width
	height := g.Dy()
	width := g.Dx()

	// The resolution is height x width
	resolution := height * width

	fmt.Println("Height", height)
	fmt.Println("Width", width)
	// Print results
	fmt.Println(resolution, "pixels")
}

func getGopherImage(path string, gopherName string) error {
	URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"
	fmt.Println("Try to get '" + gopherName + "' Gopher...")

	// Get the data
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		// Create the file
		out, err := os.Create(path + gopherName + ".png")
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer out.Close()

		// Writer the body to file
		_, err = io.Copy(out, response.Body)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Perfect! Just saved in " + out.Name() + "!")
		gopherSay(gopherName)
	} else {
		fmt.Println("Error: " + gopherName + " not exists!")
		return err
	}
	return nil
}

func getDefaultPath() (string, error) {
	path, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	path += "/desktop/"
	return path, nil
}

func gopherSay(gopherName string) {
	nbChar := len(gopherName) + len("download!")

	line := " "
	for i := 0; i <= nbChar; i++ {
		line += "-"
	}

	fmt.Println(line)
	fmt.Println("< " + gopherName + " download!" + " >")
	fmt.Println(line)
	fmt.Println("        \\")
	fmt.Println("         \\")

	fileData, err := embedGopherFiles.ReadFile("gophers/gopher.txt")
	if err != nil {
		log.Fatal("Error during read gopher ascii file", err)
	}
	fmt.Println(string(fileData))
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
