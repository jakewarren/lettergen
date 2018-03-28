package main

import (
	"fmt"
	"image/png"
	"os"
	"unicode/utf8"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/disintegration/letteravatar"
)

var version string

var (
	app        = kingpin.New("lettergen", "cmd line tool to create single-letter avatar images.")
	letter     = app.Arg("letter", "the letter to put in the logo.").Required().String()
	outputFile = app.Flag("output-file", "The name of the image file.").Short('o').String()
	size       = app.Flag("size", "the size in pixels for the image.").Short('s').Default("75").Int()
)

func main() {
	app.Version(version).VersionFlag.Short('V')
	app.HelpFlag.Short('h')
	app.UsageTemplate(kingpin.SeparateOptionalFlagsUsageTemplate)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	log.SetHandler(cli.New(os.Stderr))

	firstLetter, _ := utf8.DecodeRuneInString(*letter)

	img, err := letteravatar.Draw(*size, firstLetter, nil)
	if err != nil {
		log.WithError(err).Fatal("Error creating the avatar")
	}

	var imgFileName string

	if *outputFile == "" {
		// set a default image name
		imgFileName = string(firstLetter) + ".png"
		//log.Warnf("Defaulting the output file name to %s", imgFileName)
	} else {
		imgFileName = *outputFile
	}

	file, err := os.Create(imgFileName)
	if err != nil {
		log.WithError(err).Fatal("Error creating the PNG image")
	}

	err = png.Encode(file, img)
	if err != nil {
		log.WithError(err).Fatal("Error writing the image to the filesystem")
	}

	printSuccess("Successfully wrote the avatar to %s", imgFileName)

}

func printSuccess(format string, values ...interface{}) {
	message := fmt.Sprintf(format, values...)
	fmt.Fprintf(os.Stdout, "\033[32m%*s\033[0m %-25s\n", 4, "✓", message)
}
