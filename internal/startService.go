package internal

func StartService() {
	app := App{}
	app.Initialize()
	app.Run()
}
