// slog
package slog

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/robertkowalski/graylog-golang"
)

type glog struct {
	g        *gelf.Gelf
	address  string
	port     int
	host     string
	facility string
}

var grayLog glog

const (
	logLvlDebug = "[DEBUG] "
	logLvlInfo  = "[INFO] "
	logLvlWarn  = "[WARNING] "
	logLvlError = "[ERROR] "
	logLvlFatal = "[FATAL] "
)

var (
	debugLog   *log.Logger
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
	fatalLog   *log.Logger
)

func init() {

	debugLog = log.New(os.Stdout, "", log.LstdFlags)
	infoLog = log.New(os.Stdout, "", log.LstdFlags)
	warningLog = log.New(os.Stdout, "", log.LstdFlags)
	errorLog = log.New(os.Stderr, "", log.LstdFlags)
	fatalLog = log.New(os.Stderr, "", log.LstdFlags)
}

func ConfigLogServer(address string, port int, host string, facility string) {

	if address != "" && host != "" {

		grayLog.address = address
		grayLog.port = port
		grayLog.host = host
		grayLog.facility = facility

		grayLog.g = gelf.New(gelf.Config{
			GraylogHostname: address,
			GraylogPort:     port,
			Connection:      "lan",
		})
	}
}

func Debug(f string, v ...interface{}) {

	logContent := logLvlDebug + fmt.Sprintf(f, v...)

	debugLog.Println(logContent)

	sendToLogSvr(logContent)
}

func Info(f string, v ...interface{}) {

	logContent := logLvlInfo + fmt.Sprintf(f, v...)
	infoLog.Println(logContent)

	sendToLogSvr(logContent)
}

func Warning(f string, v ...interface{}) {

	logContent := logLvlWarn + fmt.Sprintf(f, v...)
	warningLog.Println(logContent)

	sendToLogSvr(logContent)
}

func Error(f string, v ...interface{}) {

	logContent := logLvlError + fmt.Sprintf(f, v...)
	errorLog.Println(logContent)

	sendToLogSvr(logContent)
}

func Fatal(f string, v ...interface{}) {

	logContent := logLvlFatal + fmt.Sprintf(f, v...)
	fatalLog.Println(logContent)

	sendToLogSvr(logContent)
}

func PreReqDebug(code int, r *http.Request) {
	OutputDebugLog(code, ">>> @%v @%v @%v", r.Method, r.URL.Path, r.RemoteAddr)
}

func PostReqDebug(startTime time.Time, code int, r *http.Request) {

	endTime := time.Now()

	ms := float64(endTime.Sub(startTime).Nanoseconds()/100000) / 10.0

	OutputDebugLog(code, "*** @%v @%v @%v @%v ms", r.Method, r.URL.Path, r.RemoteAddr, ms)
}

func ReqError(code int, id string, remoteAddr string, codeText string, err error) {

	if len(id) < 4 {
		OutputErrorLog(code, "@%v @%v @%v @%v", codeText, err, remoteAddr, id)
	} else {
		OutputErrorLog(code, "T(%v) @%v @%v @%v @%v", id[:4], codeText, err, remoteAddr, id)
	}
}

func ReqDebug(code int, id string, remoteAddr string, f string, v ...interface{}) {
	if len(id) < 4 {
		OutputDebugLog(code, "@%v @%v @%v", fmt.Sprintf(f, v...), remoteAddr, id)
	} else {
		OutputDebugLog(code, "T(%v) @%v @%v @%v", id[:4], fmt.Sprintf(f, v...), remoteAddr, id)
	}
}

func OutputErrorLog(code int, f string, v ...interface{}) {

	fmt.Printf("code:"+strconv.Itoa(code)+f+"\n", v...)

	logContent := logLvlError + "@" + strconv.Itoa(code) + " " + fmt.Sprintf(f, v...)
	errorLog.Println(logContent)

	sendToLogSvr(logContent)
}

func OutputDebugLog(code int, f string, v ...interface{}) {
	logContent := logLvlDebug + "@" + strconv.Itoa(code) + " " + fmt.Sprintf(f, v...)
	debugLog.Println(logContent)

	sendToLogSvr(logContent)
}

func sendToLogSvr(logContent string) {

	if grayLog.g != nil {
		func(logContent string) {

			type log struct {
				Version       string `json:"version"`
				Host          string `json:"host"`
				Facility      string `json:"facility"`
				Short_message string `json:"short_message"`
			}

			lg := &log{
				Version:       "1.0",
				Host:          grayLog.host,
				Facility:      grayLog.facility,
				Short_message: logContent,
			}

			mr, _ := json.Marshal(lg)
			grayLog.g.Log(string(mr))

		}(logContent)
	}
}
