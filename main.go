package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gizak/termui"
)

// Reads file extension and give backs percentage
func (pt *FileReaderExtension) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	if n > 0 {

		// set percentage
		percentage := float64(pt.total) / float64(pt.length) * float64(100)

		// only log if progress is above 2
		if percentage-pt.progress > 2 {
			err := termui.Init()

			if err != nil {
				panic(err)
			}

			// determine
			g0 := termui.NewGauge()
			g0.Percent = int(percentage)
			g0.Width = 50
			g0.Height = 3
			g0.BorderLabel = "Downloading File:"
			g0.BarColor = termui.ColorRed
			g0.BorderFg = termui.ColorWhite
			g0.BorderLabelFg = termui.ColorCyan

			// only set percentage of higher than 98
			if percentage > 98 {
				g0.Percent = 100
				if g0.Percent == 100 {
					g0.BarColor = termui.ColorGreen
				}
			}

			// render termui
			termui.Render(g0)

		}
	}

	return n, err
}

func main() {

	// get URL + filename to download from command args
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s url\n", os.Args[:1])
		os.Exit(1)
	} else if len(os.Args) > 3 {

	}

	// arguments, url + filename
	url := os.Args[1]
	// fileName := os.Args[2]

	// get request from given url
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
	}

	// set
	readerpt := &FileReaderExtension{Reader: resp.Body, length: resp.ContentLength}

	// read body and content length with io reader
	body, err := ioutil.ReadAll(readerpt)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading %s: %v\n", url, err)
		os.Exit(1)
	}

	// write file to correct
	err = ioutil.WriteFile("download", body, 0644)

	if err != nil {
		fmt.Println("Error while writing downloaded", url, err)
	}

	// setting directory
	dir, err := os.Getwd()

	if err != nil {
		// fmt.Fprintf("Error setting directory: %v \n", err)
	}

	fmt.Println("\n Successfully downloaded to:", dir)

	// stop loop event
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	// keep looping to check if percentage is done
	termui.Loop()

}
