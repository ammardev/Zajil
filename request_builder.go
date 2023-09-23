package main

import (
	"net/http"
	"net/http/httptrace"
	"time"
)

func sendHttpRequest(zajil *Zajil) {
    request, _ := http.NewRequest(
        zajil.methodSelector.GetMethod(),
        zajil.urlInput.GetUrl(),
        nil,
    )

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
    zajil.responseView.SetResponse(res, ttfb)
}
