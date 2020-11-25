package cli

import (
	"log"
	"net/http"
	"sync"

	"github.com/Jeffail/gabs"
)

// The request body data used for the request
type RequestBody struct {
	SourceLang string // The source language
	TargetLang string // The target language
	SourceText string // The text to be translated
}

// The translate api url
const translateUrl = "https://translate.googleapis.com/translate_a/single"

// RequestTranslate creates a request to the google translate api
func RequestTranslate(body *RequestBody, str chan string, wg *sync.WaitGroup) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", translateUrl, nil)

	query := req.URL.Query()
	query.Add("client", "gtx")
	query.Add("sl", body.SourceLang)
	query.Add("tl", body.TargetLang)
	query.Add("dt", "t")
	query.Add("q", body.SourceText)
	req.URL.RawQuery = query.Encode()

	if err != nil {
		log.Fatalf("1 There was a problem: %s", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("2 There was a problem: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		str <- "You have been rate limited, Try again later."
		wg.Done()
		return
	}

	parsedJson, err := gabs.ParseJSONBuffer(res.Body)
	if err != nil {
		log.Fatalf("3 There was a problem - %s", err)
	}

	nestOne, err := parsedJson.ArrayElement(0)
	if err != nil {
		log.Fatalf("4 There was a problem - %s", err)
	}

	nestTwo, err := nestOne.ArrayElement(0)
	if err != nil {
		log.Fatalf("5 There was a problem - %s", err)
	}

	translatedStr, err := nestTwo.ArrayElement(0)
	if err != nil {
		log.Fatalf("6 There was a problem - %s", err)
	}

	str <- translatedStr.Data().(string)
	wg.Done()
}
