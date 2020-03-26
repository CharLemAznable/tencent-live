package main

import (
    "github.com/CharLemAznable/amber"
    . "github.com/CharLemAznable/gokits"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    handleFunc(mux, "/favicon.ico", serveFavicon("favicon.ico"), DumpRequestDisabled)
    handleFunc(mux, "/res/", serveResources("/res/"), DumpRequestDisabled)

    handleFunc(mux, "/push", serveHtmlPage("push"))
    handleFunc(mux, "/push/query", ServeAjax(serveQueryPushUrl))

    handleFunc(mux, "/live", serveHtmlPage("live"))
    handleFunc(mux, "/live/query", ServeAjax(serveQueryLiveUrl))

    HandleFunc(mux, "/cocs", amber.ServeCocsHandler)

    server := http.Server{Addr: ":" + StrFromInt(appConfig.Port), Handler: mux}
    if err := server.ListenAndServe(); err != nil {
        LOG.Crashf("Start server Error: %s", err.Error())
    }
}

func handleFunc(mux *http.ServeMux, path string, handler http.HandlerFunc, opts ...HandleFuncOption) {
    if nil != amber.ConfigInstance {
        handler = amber.AuthAmber(handler)
    }
    HandleFunc(mux, path, handler, opts...)
}
