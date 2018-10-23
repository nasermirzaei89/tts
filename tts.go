package tts

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func Say(text string) ([]byte, error) {
	response, err := http.Get(fmt.Sprintf("https://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=%s&tl=en", url.QueryEscape(text)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
