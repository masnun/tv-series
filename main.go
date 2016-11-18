package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/masnun/tv-series/utils"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No directories passed!")
		os.Exit(0)
	}

	dirName := os.Args[1]

	fmt.Println("--------------------")
	fmt.Println(dirName)
	fmt.Println("--------------------")

	names := utils.GetTorrents(dirName)

	info := color.New(color.FgGreen)
	warning := color.New(color.FgRed)

	for _, mediaFile := range names {
		show := utils.GetShowInformation(mediaFile.CleanName())
		if show != nil {
			info.Printf(" => %s (Poster: %s )\n", show.Name, show.Poster)
		} else {
			warning.Printf("No matches found for: %s \n", mediaFile.BaseName)
		}
	}

}
