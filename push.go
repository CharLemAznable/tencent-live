package main

import (
    . "github.com/CharLemAznable/gokits"
    "net/http"
)

func serveQueryPushUrl(writer http.ResponseWriter, _ *http.Request) {
    pushServerUrl, pushStreamQuery := GetCurrentPushUrl()
    ResponseJson(writer, Json(map[string]interface{}{
        "pushServerUrl":   pushServerUrl,
        "pushStreamQuery": pushStreamQuery,
        "pushUrl":         pushServerUrl + pushStreamQuery,
    }))
}
