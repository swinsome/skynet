package common

import (
	"io/ioutil"
	"net/http"
)

type (
	Kdata struct {
		Date  string  `json:"date"`
		Open  float64 `json:"open"`
		Close float64 `json:"close"`
		High  float64 `json:"high"`
		Lower float64 `json:"lower"`
	}
	Stock struct {
		Name string  `json:"name"`
		Data []Kdata `json:"data"`
	}

	UrlGet interface {
		genUrl() (string, error)
		preprocess(buffer string, head string) (string, error)
		unmarshal(buffer string) (Stock, error)
	}
)

func Get_k_data(u UrlGet) (Stock, error) {
	urlStr, err := u.genUrl()
	var buffer string
	var result Stock
	resp, err := http.Get(urlStr)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	buffer = string(body)
	buffer, err = u.preprocess(buffer, "sh600016")
	result, err = u.unmarshal(buffer)
	return result, err
}
