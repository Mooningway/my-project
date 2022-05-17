package u_http

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	HEADER_CONTENT_TYPE string = `Content-Type`
	APP_FORM_URLENCODED string = `application/x-www-form-urlencoded`
	APP_JSON            string = `application/json`
	APP_JSON_UTF8       string = `application/json;charset=UTF-8`
)

type Method int

const (
	GET Method = iota
	HEAD
	POST
	PUT
	DELETE
	CONNECT
	OPTIONS
	TRACE
	PATCH
)

var (
	methodNames = []string{`GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`}
)

func (m Method) String() string {
	if GET <= m && m <= PATCH {
		return methodNames[m]
	}
	return ``
}

func parseRequestUrl(requestUrl string, params url.Values) (resultUrl string, err error) {
	if len(requestUrl) == 0 {
		err = errors.New(`requestUrl must be required`)
		return
	}
	refUrl, err := url.Parse(requestUrl)
	if err != nil {
		return
	}
	if params != nil && len(params) > 0 {
		refUrl.RawQuery = params.Encode()
	}
	resultUrl = refUrl.String()
	return
}

func setRequestHeader(request *http.Request, header map[string]interface{}) {
	if len(header) > 0 {
		for key, val := range header {
			request.Header.Add(key, fmt.Sprintf(`%v`, val))
		}
	}
}

func getResponse(request *http.Request) (responseBody []byte, err error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
