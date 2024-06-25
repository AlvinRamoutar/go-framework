package http

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

type HttpConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (h *HttpConfig) New() *HttpConfig {
	hc := HttpConfig{}

	hc.Host = ""
	hc.Port = 8080

	return &hc
}

type Http struct {
	Config             *HttpConfig
	Routes             []Route
	DefaultRoutes      []Route
	Closer             io.Closer
	DefaultRouteRegexp *regexp.Regexp
}

type Route struct {
	Name         string
	Method       Method
	Path         string
	PathType     PathType
	CompiledPath *regexp.Regexp
	Action       actionDef
}

type actionDef func(res http.ResponseWriter, req *http.Request) ResponseResult

func (h *Http) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		head, _ := ShiftPath(req.URL.Path)

		var rr ResponseResult
		for _, route := range h.Routes {
			if route.PathType == REGEX {
				if route.CompiledPath.MatchString(head) {
					rr = route.Action(res, req)
				}
			} else if route.PathType == EXACT {
				if route.Path == head {
					rr = route.Action(res, req)
				}
			}
		}

		// process response
		for _, route := range h.DefaultRoutes {
			if strconv.Itoa(rr.ResponseCode) == route.Path {
				route.Action(res, req)
			}
		}
	}
}

func (h *Http) AddRoute(route Route) error {
	if h.indexOfRoute(route.Name) == -1 {
		if route.PathType == REGEX {
			crexp, err := regexp.Compile(route.Path)
			if err != nil {
				return &HttpError{code: "ENLIBHTTP011"}
			}
			route.CompiledPath = crexp
		}
		h.Routes = append(h.Routes, route)
	}
	return &HttpError{code: "ENLIBHTTP010"}
}

func (h *Http) AddRoutes(routes []Route) error {
	var err error = nil
	for _, r := range routes {
		err = h.AddRoute(r)
		if err != nil {
			return err
		}
	}
	return err
}

func (h *Http) RemoveRoute(name string) error {
	index := -1
	for i, r := range h.Routes {
		if r.Name == name {
			index = i
			continue
		}
	}

	if index != -1 {
		h.Routes[index] = h.Routes[len(h.Routes)-1]
		h.Routes = h.Routes[:len(h.Routes)-1]
		return nil
	}
	return &HttpError{code: "ENLIBHTTP015"}
}

func (h *Http) indexOfRoute(name string) int {
	for i, r := range h.Routes {
		if r.Name == name {
			return i
		}
	}
	return -1
}

func (h *Http) listenAndServe() error {
	srv := &http.Server{Addr: fmt.Sprintf("%s:%d", h.Config.Host, h.Config.Port), Handler: h}
	return srv.ListenAndServe()
}

func (h *Http) Serve() error {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	
	go func() {
		h.listenAndServe()
		wg.Done()
	}()
	
	wg.Wait()
	return nil
}
