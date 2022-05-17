package u_http

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
)

func PostForm(requestUrl string, header map[string]interface{}, data url.Values) (responseBody []byte, err error) {
	url, err := parseRequestUrl(requestUrl, nil)
	if err != nil {
		return
	}
	request, err := http.NewRequest(POST.String(), url, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	request.Header.Add(HEADER_CONTENT_TYPE, APP_FORM_URLENCODED)
	setRequestHeader(request, header)
	return getResponse(request)
}

func PostJson(requestUrl string, header map[string]interface{}, jsonBytes []byte) (responseBody []byte, err error) {
	url, err := parseRequestUrl(requestUrl, nil)
	if err != nil {
		return
	}
	request, err := http.NewRequest(POST.String(), url, bytes.NewReader(jsonBytes))
	if err != nil {
		return
	}
	request.Header.Add(HEADER_CONTENT_TYPE, APP_JSON_UTF8)
	setRequestHeader(request, header)
	return getResponse(request)
}
