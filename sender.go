package logrus_rollbar

import (
	"net/http"

	rollbar "github.com/rollbar/rollbar-go"
)

var (
	severityMap = map[string]func(args ...interface{}){
		rollbar.CRIT:  rollbar.Critical,
		rollbar.ERR:   rollbar.Error,
		rollbar.WARN:  rollbar.Warning,
		rollbar.INFO:  rollbar.Info,
		rollbar.DEBUG: rollbar.Debug,
	}
)

type Sender interface {
	RequestError(string, *http.Request, error, map[string]interface{})
	Error(string, error, map[string]interface{})
}

type RollbarSender struct{}

func (s RollbarSender) RequestError(severity string, req *http.Request, err error, fields map[string]interface{}) {
	severityMap[severity](req, err, fields)
}

func (s RollbarSender) Error(severity string, err error, fields map[string]interface{}) {
	severityMap[severity](err, fields)
}
