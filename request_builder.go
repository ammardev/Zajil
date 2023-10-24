package main

import (
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"
)

func sendHttpRequest(zajil *Zajil) {
	request, _ := http.NewRequest(
		zajil.methodSelector.GetMethod(),
		zajil.urlInput.GetUrl(),
		nil,
	)

	parseHeaders(request, zajil.rc.HeadersTextInput.Value())

	var startTime time.Time
	var ttfb int

	trace := &httptrace.ClientTrace{
		DNSStart: func(httptrace.DNSStartInfo) { startTime = time.Now() },

		GotFirstResponseByte: func() {
			ttfb = int(time.Since(startTime).Milliseconds())
		},
	}

	request = request.WithContext(httptrace.WithClientTrace(request.Context(), trace))

	res, _ := http.DefaultClient.Do(request)

	http.DefaultClient.CloseIdleConnections()

	zajil.responseView.SetResponse(res, ttfb)
}

func parseHeaders(request *http.Request, headersBlock string) {
	// TODO: Add real parsing here.

	rawHeadersList := strings.Split(headersBlock, "\n")

	for _, rawHeader := range rawHeadersList {
		if rawHeader == "" {
			continue
		}

		keyValueSplit := strings.Split(rawHeader, ":")
		request.Header.Add(keyValueSplit[0], keyValueSplit[1])
	}
}
