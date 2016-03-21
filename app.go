package goapp

import (
	"github.com/gin-gonic/gin"
)

var (
	AppInstance *App
)

func init() {
	AppInstance = NewApp()
}

func NewApp() *App {
	gin.SetMode(gin.DebugMode)
	app := new(App)
	app.engine = gin.New()
	app.listen = ":12345"
	dbmanConf := &DBManConfig{
		write: &DBConfig{driver: "mysql", dsn: "root:pupumed@tcp(localhost:3306)/test?charset=utf8"},
		read:  []*DBConfig{&DBConfig{driver: "mysql", dsn: "root:pupumed@tcp(localhost:3306)/test?charset=utf8"}},
	}
	app.dbs = make(map[string]*DBMan)
	app.dbs["default"] = NewDBMan(dbmanConf)
	return app
}

type App struct {
	engine *gin.Engine
	dbs    map[string]*DBMan
	listen string
}

func (app *App) Engine() *gin.Engine {
	return app.engine
}

func (app *App) Run() {
	app.engine.Run(app.listen)
}

func (app *App) GetDB(name string) *DBMan {
	db, _ := app.dbs[name]
	return db
}

func (app *App) GetCache(name string) {

}
