package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {

	fmt.Println("Checking mime type for .wasm : ", mime.TypeByExtension(".wasm"))

	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &hello{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	url := "localhost:8080"
	fmt.Println("Listening to ", url)
	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatal(err)
	}
}

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

func (h *hello) OnAppUpdate(ctx app.Context) {
	fmt.Println("Mounting app ...")
	//ctx.Async(tick)
	_ = tick
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Success ! "),
		app.Br(),
		app.Div().Text("Time is "+time.Now().Format("2 jan 2006 at 15h04m05s")),
		app.If(app.IsClient), app.Div().Text("Is client"),
		app.If(app.IsServer), app.Div().Text("Is server"),
	)
}

func tick() {
	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		fmt.Println("Tick ...")
	}
}
