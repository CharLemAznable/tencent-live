package main

import (
    . "github.com/CharLemAznable/gokits"
    "net/http"
)

func serveQueryLiveUrl(writer http.ResponseWriter, _ *http.Request) {
    m3u8Url := GetCurrentLiveUrl(".m3u8")
    flvUrl := GetCurrentLiveUrl(".flv")
    ResponseJson(writer, Json(map[string]interface{}{
        "m3u8Url": m3u8Url,
        "flvUrl": flvUrl,
    }))
}
