package rollbar

import (
	"io"
	"net/http"
)

//Poster can be used to inject into Client custom POST handler
//This is very useful when using Google App engine and http.Post can not be used
type Poster interface {
	Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
}

type DefaultPoster struct{
}

func (defaultPoster DefaultPoster) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	resp, err = http.Post(url, contentType, body)

	return
}