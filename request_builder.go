package main

import (
	"net/http"
)

func sendHttpRequest(zajil *Zajil) {
    request, _ := http.NewRequest(
        zajil.methodSelector.GetMethod(),
        zajil.urlInput.GetUrl(),
        nil,
    )
    res, _ := http.DefaultClient.Do(request)
    zajil.responseView.SetResponse(*res)
}
