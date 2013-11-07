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
	"github.com/Jackong/log/writer"
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
	fmt.Println("init env...")
	baseEnv()

	fmt.Println("loading config...")
	loadConfig()

	fmt.Println("opening db...")
	openDb()

	fmt.Println("init router...")
	Router = mux.NewRouter()

	fmt.Println("getting mail log...")
	mailLog := mailLog()

	fmt.Println("getting file log...")
	fileLog := fileLog("access.log")

	Log = log.MultiLogger(fileLog, mailLog)
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

func fileLog(name string) log.Logger {
	file, err := os.OpenFile(Project.String("log", "dir") + "/" + name, os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return log.NewLogger(file, Project.String("server", "name"), int(Project.Get("log", "file", "level").(float64)))
}

func mailLog() log.Logger {
	mail := &writer.Email{
		User: Project.String("log", "email", "user"),
		Password: Project.String("log", "email", "password"),
		Host : Project.String("log", "email", "host"),
		To : Project.String("log", "email", "to"),
		Subject: Project.String("log", "email", "subject"),
	}

	server := Project.String("server", "name")
	if err := mail.SendMail("starting server " + server + "..."); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return log.NewLogger(&asyncMail{mail}, server, int(Project.Get("log", "email", "level").(float64)))
}

func openDb() {
	settings := db.DataSource{
		Socket:   Project.String("db", "socket"),
		Database: Project.String("db", "database"),
		User:     Project.String("db", "user"),
		Password: Project.String("db", "password"),
	}

	var err error
	if Conn, err = db.Open("mysql", settings); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
