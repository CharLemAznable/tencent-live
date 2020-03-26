package main

import (
    "crypto/md5"
    "flag"
    "fmt"
    "github.com/BurntSushi/toml"
    "github.com/CharLemAznable/amber"
    "github.com/CharLemAznable/gokits"
    "strings"
    "sync"
    "time"
    "unsafe"
)

type AppConfig struct {
    gokits.HttpServerConfig
    DevMode bool

    PushServer    string
    PushSecretKey string
    LiveServer    string
    LiveSecretKey string

    StreamName string
    Timeout    int64

    CurrLock       sync.RWMutex
    CurrTxTime     int64
    CurrPushStream string
    CurrLiveStream string

    AmberLoginEnabled bool
    AmberAppId        string
    AmberEncryptKey   string
    AmberCookieName   string
    AmberLoginUrl     string
    AmberLocalUrl     string
    AmberForceLogin   bool
}

var appConfig AppConfig
var _configFile string

func init() {
    gokits.LOG.LoadConfiguration("logback.xml")

    flag.StringVar(&_configFile, "configFile", "appConfig.toml", "config file path")
    flag.Parse()

    if _, err := toml.DecodeFile(_configFile, &appConfig); err != nil {
        gokits.LOG.Crashf("config file decode error: %s", err.Error())
    }

    gokits.If(0 == appConfig.Port, func() {
        appConfig.Port = 20514
    })
    gokits.If(0 != len(appConfig.ContextPath), func() {
        gokits.Unless(strings.HasPrefix(appConfig.ContextPath, "/"),
            func() { appConfig.ContextPath = "/" + appConfig.ContextPath })
        gokits.If(strings.HasSuffix(appConfig.ContextPath, "/"),
            func() { appConfig.ContextPath = appConfig.ContextPath[:len(appConfig.ContextPath)-1] })
    })

    gokits.If(0 == appConfig.Timeout, func() {
        appConfig.Timeout = 3600
    })

    appConfig.CurrTxTime = 0
    appConfig.CurrPushStream = ""
    appConfig.CurrLiveStream = ""
    CalculateCurrentQuery()

    if appConfig.AmberLoginEnabled {
        amber.ConfigInstance = amber.NewConfig(
            amber.WithAppId(appConfig.AmberAppId),
            amber.WithEncryptKey(appConfig.AmberEncryptKey),
            amber.WithCookieName(appConfig.AmberCookieName),
            amber.WithAmberLoginUrl(appConfig.AmberLoginUrl),
            amber.WithLocalUrl(appConfig.AmberLocalUrl),
            amber.WithForceLogin(appConfig.AmberForceLogin),
        )
    }

    gokits.GlobalHttpServerConfig = (*gokits.HttpServerConfig)(unsafe.Pointer(&appConfig))
    gokits.LOG.Debug("appConfig: %s", gokits.Json(appConfig))
}

func CalculateCurrentQuery() {
    appConfig.CurrLock.Lock()
    defer appConfig.CurrLock.Unlock()

    if appConfig.CurrTxTime <= time.Now().Unix() {
        appConfig.CurrTxTime = time.Now().Unix() + appConfig.Timeout
        txTimeHex := strings.ToUpper(fmt.Sprintf("%x", appConfig.CurrTxTime))

        txPushSecret := fmt.Sprintf("%x", md5.Sum([]byte(appConfig.PushSecretKey+appConfig.StreamName+txTimeHex)))
        appConfig.CurrPushStream = "?txSecret=" + txPushSecret + "&txTime=" + txTimeHex

        txLiveSecret := fmt.Sprintf("%x", md5.Sum([]byte(appConfig.LiveSecretKey+appConfig.StreamName+txTimeHex)))
        appConfig.CurrLiveStream = "?txSecret=" + txLiveSecret + "&txTime=" + txTimeHex
    }
}

func GetCurrentPushUrl() (string, string) {
    CalculateCurrentQuery()
    return appConfig.PushServer,
        appConfig.StreamName + appConfig.CurrPushStream
}

func GetCurrentLiveUrl(protocol string) string {
    CalculateCurrentQuery()
    return appConfig.LiveServer +
        appConfig.StreamName + protocol + appConfig.CurrLiveStream
}
