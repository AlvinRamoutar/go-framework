package lang

import (
	"encoding/json"
	"errors"
	"strconv"

	"alvinr.ca/go-framework/lang/data"
)

var VERSION string = "0.0.1"

// Key format is [AABBBCCCCDDD] where
//   - AA is language 2-char shortcode (e.g. EN, FR)
//   - BBB is type 3-char shortcode (e.g. SVC, LIB)
//   - CCCC is type name 4-char shortcode (e.g. LANG, CORE)
//   - DDD is phrase unique 3-digit numerical id (e.g. 001, 999)
type Lang struct {
	Language string
	Data     map[string]string
}

func (l *Lang) Init() error {
	l.Language = "EN"
	l.Data = map[string]string{}
	return nil
}

func (l *Lang) Start() error {
	l.Init()
	langPacks := []data.Pack{data.ENPack{}}
	for _, pack := range langPacks {
		for k, v := range pack.New() {
			l.Data[k] = v
		}
	}
	return nil
}

func (l *Lang) AsyncStart() error {
	return errors.New("ENLIBCOMM007")
}

func (l *Lang) Restart() error {
	l.Stop()
	return l.Start()
}

func (l *Lang) Stop() error {
	for k, _ := range l.Data {
		delete(l.Data, k)
	}
	return nil
}

func (l *Lang) Status() (string, error) {
	status := map[string]string{
		"loaded": strconv.Itoa(len(l.Data)),
	}
	j, _ := json.Marshal(status)
	return string(j), nil
}

func (l *Lang) Version() string {
	return VERSION
}

func (l *Lang) Get(key string) string {
	if len(key) == 10 {
		return l.Data[l.Language+key]
	} else {
		return l.Data[key]
	}
}

func (l *Lang) Set(language string) {
	l.Language = language
}

func (l *Lang) Extend(ext map[string]string) {
	for k, v := range ext {
		l.Data[k] = v
	}
}
