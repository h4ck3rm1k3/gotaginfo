package main

import (
//	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	bow := surf.NewBrowser()
	err := bow.Open("https://taginfo.openstreetmap.org/download")
	if err != nil {
		panic(err)
	}

	// Outputs: "The Go Programming Language"
	fmt.Println(bow.Title())

	re := regexp.MustCompile(`\w+\.db\.bz2$`)
	for _, s := range bow.Links() {
		//fmt.Println("Link",i,s)
		if re.MatchString(s.Text) {
			fmt.Println(s.Text)

			filename := s.Text
			fout, err := os.Create(filename)
			if err != nil {
				fmt.Printf(
					"Error creating file '%s'.", filename)
				continue
			}
			defer fout.Close()

			resp, err := http.Get(s.Asset.Url().String())
			if err != nil {
				fmt.Printf(
					"Error downloading file '%s' %s.", filename, err)

			}
			defer resp.Body.Close()

	
			_, err = io.Copy(fout, resp.Body)
			if err != nil {
				fmt.Printf(
					"Error downloading file '%s'.", filename)
			}
			
		}
	}
}
