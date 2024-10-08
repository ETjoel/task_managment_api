package bootstrap

import "go.mongodb.org/mongo-driver/mongo"

type Application struct {
	Env    Env
	Client *mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Client = ConnectMongoDB(app.Env)
	return *app
}

func (app *Application) CloseDatabase() {
	DisconnectMongoDB(app.Client)
}

func (app *Application) CreateEmailUniqueIndex() {
	CreateEmailUniqueIndex(app.Env, app.Client)
}
