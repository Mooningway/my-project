package u_http

import (
	"net/http"
	"net/url"
)

func Get(requestUrl string, header map[string]interface{}, params url.Values) (responseBody []byte, err error) {
	url, err := parseRequestUrl(requestUrl, params)
	if err != nil {
		return
	}
	request, err := http.NewRequest(GET.String(), url, nil)
	if err != nil {
		return
	}
	setRequestHeader(request, header)
	return getResponse(request)
}
