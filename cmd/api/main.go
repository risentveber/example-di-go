package main

func main() {
	app, err := BuildApp()
	if err != nil {
		panic(err)
	}
	err = app.Start()
	if err != nil {
		panic(err)
	}
}
