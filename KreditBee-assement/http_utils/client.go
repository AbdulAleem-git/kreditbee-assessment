package http_utils

import "net/http"

var Client *http.Client
func NewClient() *http.Client{
	if Client == nil {
		return &http.Client{}
	}else{
		return Client
	}

}