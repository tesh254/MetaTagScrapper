package main

import (
	"fmt"
	// "io/ioutil"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Meta define meta data type
type Meta struct {
	image       string
	description string
	url         string
	title       string
	site        string
}

// MetaType defines meta type and details
type MetaType struct {
	metaType    string
	metaDetails Meta
}

func callWebsite(listingURL string) []string {
	// var metas []string/

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("GET", listingURL, nil)

	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("pragma", "no-cache")
	request.Header.Set("cache-control", "no-cache")
	request.Header.Set("dnt", "1")
	request.Header.Set("upgrade-insecure-requests", "1")
	request.Header.Set("referer", "https://youtube.com/")

	resp, err := client.Do(request)

	if resp.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			fmt.Println(err)
		}

		var facebookMetaType MetaType = MetaType{
			metaType: "facebook",
			metaDetails: Meta{
				image:       "",
				description: "",
				url:         "",
				title:       "",
				site:        "",
			},
		}

		var twitterMetaType MetaType = MetaType{
			metaType: "twitter",
			metaDetails: Meta{
				image:       "",
				description: "",
				url:         "",
				title:       "",
				site:        "",
			},
		}

		doc.Find("meta").Each(func(i int, s *goquery.Selection) {
			metaProperty, _ := s.Attr("property")
			metaContent, _ := s.Attr("content")

			// facebook -> og:sitename
			// facebook -> og:url
			// facebook -> og:title
			// facebook -> og:description
			// facebook -> og:image

			// twitter -> twitter:site
			// twitter -> twitter:site,
			// twitter -> twitter:image

			if metaProperty == "og:site_name" {
				facebookMetaType.metaDetails.site = metaContent
			}

			if metaProperty == "og:url" {
				facebookMetaType.metaDetails.url = metaContent
			}

			if metaProperty == "og:title" {
				facebookMetaType.metaDetails.title = metaContent
			}

			if metaProperty == "og:description" {
				facebookMetaType.metaDetails.description = metaContent
			}

			if metaProperty == "og:image" {
				facebookMetaType.metaDetails.image = metaContent
			}

			if metaProperty == "twitter:site" {
				twitterMetaType.metaDetails.site = metaContent
			}

			if metaProperty == "twitter:title" {
				twitterMetaType.metaDetails.title = metaContent
			}

			if metaProperty == "twitter:description" {
				twitterMetaType.metaDetails.description = metaContent
			}

			if metaProperty == "twitter:image" {
				twitterMetaType.metaDetails.image = metaContent
			}

			fmt.Println(metaProperty)
		})

		fmt.Println(facebookMetaType)

		fmt.Println(twitterMetaType)
	}
	return nil
}

func main() {
	callWebsite("https://youtube.com")
}
