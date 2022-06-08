package main

import (
	"bytes"
	"io"
	"net/http"
)

func recRequestView(id string, lang string, langCode string, client *http.Client) (resp *http.Response) {
	resp, err := requestView(id, lang, langCode, client)
	if err != nil {
		resp = recRequestView(id, lang, langCode, client)
	}
	return resp
}

func getView(index int, id string, channel chan Result, client *http.Client) (err error) {

	resp := recRequestView(id, "eng", "6", client)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	resp.Body.Close()

	reader := bytes.NewReader(body)

	data, err := ParseView(reader, index, id, "6")
	if err != nil {
		panic(err)
	}

	channel <- data

	return err
}
