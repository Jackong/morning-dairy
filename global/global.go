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
	"github.com/gosexy/db"
	_ "github.com/gosexy/db/mysql"
)

var (
	GoPath string
	Now func() string
	Time func() time.Time
	Project config.Config
	Log log.Logger
	Conn db.Database
	Router *mux.Router
)

func init() {
	baseEnv()
	loadConfig()
	openDb()

	Router = mux.NewRouter()

	debug := fileLog("debug.log", log.LEVEL_DEBUG)
	info := fileLog("info.log", log.LEVEL_INFO)
	error := fileLog("error.log", log.LEVEL_ERROR)
	Log = log.MultiLogger(debug, info, error)
}

func baseEnv() {
	Time = func() time.Time {
		return time.Now()
	}

	Now = func() string {
		return Time().Format("2006-01-02 15:04:05")
	}

	GoPath = os.Getenv("GOPATH")
}

func loadConfig() {
	Project = config.NewConfig(GoPath  + "/src/morning-dairy/config/project.json")
}

func fileLog(name string, level int) log.Logger {
	logFile, err := os.OpenFile(Project.String("log", "dir") + "/" + name, os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return log.NewLogger(logFile, Project.String("server", "name"), level)
}

func openDb() {
	settings := db.DataSource{
		Socket:   "/var/run/mysqld/mysqld.sock",
		Database: "test",
		User:     "root",
		Password: "123456",
	}

	var err error
	if Conn, err = db.Open("mysql", settings); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
