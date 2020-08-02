package scrapper

import (
	"fmt"
	// "io/ioutil"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Scrapi house scrapper methods and logic
type Scrapi struct{}

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

// CallWebsite make an api call
func (srapper *Scrapi) CallWebsite(websiteURL string) []MetaType {
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

	client := &http.Client{
		// Set request timeout
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("GET", websiteURL, nil)

	if err != nil {
		fmt.Println(err)
	}

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Pragma
	request.Header.Set("pragma", "no-cache")
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control
	request.Header.Set("cache-control", "no-cache")
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/DNT
	request.Header.Set("dnt", "1")
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Upgrade-Insecure-Requests
	request.Header.Set("upgrade-insecure-requests", "1")
	request.Header.Set("referer", websiteURL)

	// Make API call
	resp, err := client.Do(request)

	// if we have a successful request
	if resp.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			fmt.Println(err)
		}

		// Map through all meta tags fetched
		doc.Find("meta").Each(func(i int, s *goquery.Selection) {
			metaProperty, _ := s.Attr("property")
			metaContent, _ := s.Attr("content")

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
		})
	}
	var rp []MetaType

	rp = append(rp, facebookMetaType)

	rp = append(rp, twitterMetaType)

	return rp
}
