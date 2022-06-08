package main

import (
	"fmt"
	"net/http"
)

func requestView(id string, lang string, langCode string, client *http.Client) (resp *http.Response, err error) {

	url := fmt.Sprintf("https://krdict.korean.go.kr/%s/dicSearch/SearchView?ParaWordNo=%s&nation=%s&nationCode=%s&viewType=A&blockCount=10&viewTypes=on",
		lang,
		id,
		lang,
		langCode,
	)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	resp, err = client.Do(req)

	return resp, err
}
