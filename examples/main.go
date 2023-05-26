package main

import (
	nhttp "net/http"
	"sync"

	"alvinr.ca/go-framework/base"
	"alvinr.ca/go-framework/http"
	"alvinr.ca/go-framework/lang"
	"alvinr.ca/go-framework/log"
)

var APP_NAME string = "core"
var VERSION string = "0.0.1"

var Libs map[string]base.Library
var Lang *lang.Lang = &lang.Lang{}
var Log *log.Log = &log.Log{}
var Conf *Config = &Config{}
var Http *http.Http = &http.Http{}

func Act(res nhttp.ResponseWriter, req *nhttp.Request) http.ResponseResult {
	LogInfoWithMessage("http", "ENLIBHTTP050", req.Method + " " + req.Host + req.RequestURI)
	return http.ResponseResult{
		ResponseCode: 200,
		Error: nil,
	}
}

func main() {

	var libsOrdered [4]string = [4]string{"lang", "log", "conf", "http"}

	// Module config
	logConfig := &log.LogConfig{}
	logConfig = logConfig.New()
	Log.Init(logConfig)

	Conf.Init(logConfig)

	Http.Init(Conf.Config.Http)

    newRoute := http.Route{
		Name: "helloworld",
		Method: http.GET,
		Path: "helloworld",
		PathType: http.EXACT,
	    Action: Act,
	}
	Http.AddRoute(newRoute)
	
	// Starting modules
	Libs = map[string]base.Library{
		"lang": Lang,
		"log":  Log,
		"conf": Conf,
		"http": Http,
	}
	for _, name := range libsOrdered {
		err := Libs[name].Start()

		// Skip logging for specific modules, since logger isn't setup yet and they're not expected to fail
		if name == "lang" {
			continue
		}
		if err != nil {
			if err.Error() == "ENLIBCOMM008" {
				LogInfo(name, err.Error())
			} else {
				LogFatalWithException(name, "LIBCOMM009", err)
			}
		} else {
			LogInfo(name, "LIBCOMM010")
		}
	}

	// Post-init modules
	Lang.Set(Conf.Config.Language)
	Log.Reload(Conf.Config.Logging)
	LogInfo("log", "LIBCOMM012")

	// Spawn routines
	spawnRoutines().Wait()
}

func spawnRoutines() *sync.WaitGroup {
	grwg := new(sync.WaitGroup)
	grwg.Add(1)
	go func() {
		LogFatalWithException("http", "ENLIBHTTP009", Http.AsyncStart())
		grwg.Done()
	}()
	return grwg
}
