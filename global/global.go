/**
 * User: jackong
 * Date: 11/5/13
 * Time: 4:50 PM
 */
package global

import (
	"github.com/gorilla/mux"
	"github.com/Jackong/log"
	"os"
	"fmt"
	"morning-dairy/config"
	"time"
)

var (
	GoPath string
	Now func() string
	Time func() time.Time
	Project config.Config
	Log log.Logger
	Router *mux.Router
)

func init() {
	Now = func() string {
		return time.Now().Format("2006-01-02 15:04:05")
	}
	Time = func() time.Time {
		return time.Now()
	}

	GoPath = os.Getenv("GOPATH")
	Project = config.NewConfig(GoPath  + "/src/morning-dairy/config/project.json")
	Router = mux.NewRouter()

	debug := fileLog("debug.log", log.LEVEL_DEBUG)
	info := fileLog("info.log", log.LEVEL_INFO)
	error := fileLog("error.log", log.LEVEL_ERROR)
	Log = log.MultiLogger(debug, info, error)
}

func fileLog(name string, level int) log.Logger {
	logFile, err := os.OpenFile(Project.String("log", "dir") + "/" + name, os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return log.NewLogger(logFile, Project.String("server", "name"), level)
}
