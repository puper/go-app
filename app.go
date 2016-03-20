package goapp

var (
	AppInstance *App
)

func init() {
	AppInstance = NewApp()
}

func NewApp() *App {
	return &App{
		dbs:
	}
}

type App struct {
	dbs map[string]*DBMan
}

func (app *App) GetDB(name string) *DBMan {
}

func (app *App) GetCache(name string) {
	
}