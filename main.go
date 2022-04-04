package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

const (
	EWS_CUSTOM_SOURCE_GURU_ID               = ""
	EWS_API_KEY                             = ""
	EWS_BEARER                              = "Bearer " + EWS_API_KEY
	EWS_BASE_URL                            = "https://my-deployment-d14254.ent.eastus2.azure.elastic-cloud.com"
	EWS_CUSTOM_SOURCE_GURU_DOCS_URL         = EWS_BASE_URL + "/api/ws/v1/sources/" + EWS_CUSTOM_SOURCE_GURU_ID + "/documents"
	EWS_CUSTOM_SOURCE_GURU_BULK_CREATE_URL  = EWS_BASE_URL + "/api/ws/v1/sources/" + EWS_CUSTOM_SOURCE_GURU_ID + "/documents/bulk_create"
	EWS_CUSTOM_SOURCE_GURU_BULK_DESTROY_URL = EWS_BASE_URL + "/api/ws/v1/sources/" + EWS_CUSTOM_SOURCE_GURU_ID + "/documents/bulk_destroy"

	GURU_QUERY_URL       = "https://api.getguru.com/api/v1/search/query"
	GURU_USER            = "@bit.ly"
	GURU_PASS            = ""
	GURU_COLLECTIONS_URL = "https://api.getguru.com/api/v1/collections"
	GURU_CARD_URL        = "https://app.getguru.com/card"

	DT_LAYOUT = "2006-01-02T15:04:05.000"
)

func GetGuruCardsToImport() []string {
	return []string{"63ef7ed0-d38d-490f-a30f-43f034ae9b2e", "d07dbc11-6b04-49b0-adbe-7974ab94f579", "25d54fe9-cec3-4710-8649-8dcf1e70c825"}
}

func PostTeamCardsToEWS() (res string, err error) {
	guruCards := GetGuruCardsByIDs(GetGuruCardsToImport())

	EWSDocs, err := ConvertGuruCardsExtendedToEWSDocs(guruCards)

	if err != nil {
		return "", err
	}

	res, err = PostGuruDocsToEWS(EWSDocs)

	if err != nil {
		return "", err
	}

	return
}

func GetGuruCardsByIDs(slugs []string) []GuruCardExtended {
	if slugs == nil {
		return nil
	}

	var allCards []GuruCardExtended

	var card GuruCardExtended

	for _, s := range slugs {
		url := "https://api.getguru.com/api/v1/cards/" + s + "/extended"
		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("Accept", "application/json")

		req.SetBasicAuth(GURU_USER, GURU_PASS)

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			fmt.Println(err)
			return nil
		}

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		ioutil.WriteFile("test.json", body, 0644)

		if err := json.Unmarshal(body, &card); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}

		allCards = append(allCards, card)

	}

	return allCards
}

func ConvertGuruCardsExtendedToEWSDocs(gr []GuruCardExtended) (docs []EWSGuruDoc, err error) {

	for _, s := range gr {
		var t time.Time
		t, err = time.Parse(DT_LAYOUT, s.DateCreated[:23])
		if err != nil {
			return nil, err
		}
		var body string
		body, err = SanitizeHTML(s.Content)
		if err != nil {
			return nil, err
		}

		docs = append(docs, EWSGuruDoc{
			ID:        s.ID,
			Title:     s.PreferredPhrase,
			Body:      body,
			URL:       GURU_CARD_URL + "/" + s.Slug,
			CreatedAt: t,
			Type:      s.Collection.CollectionType,
		})
	}

	return
}

func GetGuruCards() (gr GuruCardsQueryResponse) {

	req, _ := http.NewRequest("GET", GURU_QUERY_URL, nil)

	req.Header.Add("Accept", "application/json")

	req.SetBasicAuth(GURU_USER, GURU_PASS)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	ioutil.WriteFile("guru_get_cards_response.json", body, 0644)

	fmt.Println(res.Status)

	if err := json.Unmarshal(body, &gr); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return

}

func SanitizeHTML(content string) (html string, err error) {
	p := bluemonday.StrictPolicy()

	html = p.Sanitize(
		content,
	)

	return
}

func ConvertGuruCardsToEWSDocs(gr GuruCardsQueryResponse) (docs []EWSGuruDoc, err error) {

	for _, s := range gr {
		var t time.Time
		t, err = time.Parse(DT_LAYOUT, s.DateCreated[:23])
		if err != nil {
			return nil, err
		}
		var body string
		body, err = SanitizeHTML(s.Content)
		if err != nil {
			return nil, err
		}

		docs = append(docs, EWSGuruDoc{
			ID:        s.ID,
			Title:     s.PreferredPhrase,
			Body:      body,
			URL:       GURU_CARD_URL + "/" + s.Slug,
			CreatedAt: t,
			Type:      s.Collection.CollectionType,
		})
	}

	return
}

func DeleteGuruDocsFromEWS(docIDsToBeDeleted []string) (bodyString string, err error) {

	docIDsJSON, err := json.Marshal(docIDsToBeDeleted)

	ioutil.WriteFile("ews_docs_to_delete.json", docIDsJSON, 0644)

	if err != nil {
		return
	}

	req, _ := http.NewRequest("POST", EWS_CUSTOM_SOURCE_GURU_BULK_DESTROY_URL, bytes.NewBuffer(docIDsJSON))

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", EWS_BEARER)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	bodyString = string(body)

	return
}

func DeleteAllGuruDocsFromEWS() (*http.Response, error) {

	req, _ := http.NewRequest("DELETE", EWS_CUSTOM_SOURCE_GURU_DOCS_URL, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", EWS_BEARER)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(body)

	return res, nil

}

func GetGuruCollection() {

	req, _ := http.NewRequest("GET", GURU_COLLECTIONS_URL, nil)

	req.Header.Add("Accept", "application/json")

	req.SetBasicAuth(GURU_USER, GURU_PASS)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res.Status)
	// fmt.Println(string(body))

	fmt.Printf("Got %d records", len(body))

	// os.WriteFile("response.json", body, 0644)

	var gr GuruCollectionsResponse

	if err := json.Unmarshal(body, &gr); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

}

func Post50GuruDocsToEWS() (res string, err error) {
	guruCards := GetGuruCards()

	EWSDocs, err := ConvertGuruCardsToEWSDocs(guruCards)

	if err != nil {
		return "", err
	}

	res, err = PostGuruDocsToEWS(EWSDocs)

	if err != nil {
		return "", err
	}

	return
}

func GetGuruDocsFromEWS(cursor string) (response *EWSGuruGetDocs, err error) {

	var postBody io.Reader

	if cursor != "" {
		cursorbody := CursorBody{Cursor: cursor}
		j, err := json.Marshal(cursorbody)
		if err != nil {
			return nil, err
		}
		postBody = bytes.NewBuffer(j)
	}

	req, _ := http.NewRequest("POST", EWS_CUSTOM_SOURCE_GURU_DOCS_URL, postBody)

	req.Header.Add("Accept", "application/json")

	req.Header.Add("Authorization", EWS_BEARER)

	var res *http.Response

	res, err = http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
		log.Println("Can not unmarshal JSON")
		return nil, err
	}

	return
}

func GetAllGuruDocsFromEWS() ([]*EWSGetDocsResults, error) {

	var results *EWSGuruGetDocs
	var err error

	allResults := make([]*EWSGetDocsResults, 0)

	results, err = GetGuruDocsFromEWS("")

	allResults = append(allResults, results.Results...)

	for results.Meta.Cursor.Next != "" {
		results, err = GetGuruDocsFromEWS(results.Meta.Cursor.Next)
		allResults = append(allResults, results.Results...)
	}

	return allResults, err

}

func PostGuruDocsToEWS(docs []EWSGuruDoc) (bodyString string, err error) {

	docsJSON, err := json.Marshal(docs)

	ioutil.WriteFile("ews_guru_docs_exported.json", docsJSON, 0644)

	if err != nil {
		return
	}

	req, _ := http.NewRequest("POST", EWS_CUSTOM_SOURCE_GURU_BULK_CREATE_URL, bytes.NewBuffer(docsJSON))

	// fmt.Println(EWS_CUSTOM_SOURCE_GURU_POST_DOCS_URL)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	req.Header.Add("Authorization", EWS_BEARER)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	bodyString = string(body)

	return
}

func main() {
	PostTeamCardsToEWS()

	// 	res, err := DeleteAllGuruDocsFromEWS()

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	fmt.Println(res)
}
