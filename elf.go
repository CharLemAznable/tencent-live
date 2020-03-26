package main

import (
    "bytes"
    "github.com/CharLemAznable/gokits"
    "io"
    "mime"
    "net/http"
    "path/filepath"
    "strings"
)

func detectContentType(name string) (t string) {
    if t = mime.TypeByExtension(filepath.Ext(name)); t == "" {
        t = "application/octet-stream"
    }
    return
}

func serveFavicon(path string) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
        fi, _ := AssetInfo(path)
        buffer := bytes.NewReader(MustAsset(path))
        writer.Header().Set("Content-Type", detectContentType(fi.Name()))
        writer.Header().Set("Last-Modified", fi.ModTime().UTC().Format(http.TimeFormat))
        writer.WriteHeader(http.StatusOK)
        _, _ = io.Copy(writer, buffer)
    }
}

func serveResources(prefix string) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
        filename := request.URL.Path[len(gokits.PathJoin(appConfig.ContextPath, prefix)):]
        if strings.HasSuffix(filename, ".html") {
            writer.WriteHeader(http.StatusNotFound)
            return
        }
        fi, _ := AssetInfo(filename)
        if fi == nil {
            writer.WriteHeader(http.StatusNotFound)
            return
        }

        fileContent := string(MustAsset(filename))
        if strings.HasSuffix(filename, ".js") {
            fileContent = gokits.MinifyJs(fileContent, appConfig.DevMode)
        } else if strings.HasSuffix(filename, ".css") {
            fileContent = gokits.MinifyCSS(fileContent, appConfig.DevMode)
        }
        fileContent = strings.Replace(fileContent, "${contextPath}", appConfig.ContextPath, -1)
        buffer := bytes.NewReader([]byte(fileContent))
        writer.Header().Set("Content-Type", detectContentType(fi.Name()))
        writer.Header().Set("Last-Modified", fi.ModTime().UTC().Format(http.TimeFormat))
        writer.WriteHeader(http.StatusOK)
        _, _ = io.Copy(writer, buffer)
    }
}

func serveHtmlPage(htmlName string) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
        htmlContent := string(MustAsset(htmlName + ".html"))
        htmlContent = gokits.MinifyHTML(htmlContent, appConfig.DevMode)
        htmlContent = strings.Replace(htmlContent, "${contextPath}", appConfig.ContextPath, -1)

        modelCtx := gokits.ModelContext(request.Context())
        for key, value := range modelCtx.Model {
            valueStr, ok := value.(string)
            if !ok {
                valueStr = gokits.Json(value)
            }
            htmlContent = strings.Replace(htmlContent, "${"+key+"}", valueStr, -1)
        }

        gokits.ResponseHtml(writer, htmlContent)
    }
}
