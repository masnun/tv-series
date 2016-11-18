package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TVShow struct {
	Name   string
	Poster string
}

func GetShowInformation(title string) *TVShow {

	title = strings.ToLower(title)
	apiUrl := fmt.Sprintf("http://api.tvmaze.com/search/shows?q=%s", title)
	response, error := http.Get(apiUrl)
	if error != nil {
		fmt.Println(error)
	}
	defer response.Body.Close()
	jsonBytes, error := ioutil.ReadAll(response.Body)

	var data []map[string]interface{}
	json.Unmarshal(jsonBytes, &data)

	if len(data) > 0 {
		matchedShow := data[0]
		showDict := matchedShow["show"].(map[string]interface{})
		name := showDict["name"].(string)

		imgDict := showDict["image"].(map[string]interface{})
		originalImage := imgDict["original"].(string)

		show := &TVShow{Name: name, Poster: originalImage}

		return show
	}

	return nil
}
