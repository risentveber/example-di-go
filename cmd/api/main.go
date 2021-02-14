package main

func main() {
	app, err := BuildInRuntime()
	if err != nil {
		panic(err)
	}
	err = app.Start()
	if err != nil {
		panic(err)
	}
}
