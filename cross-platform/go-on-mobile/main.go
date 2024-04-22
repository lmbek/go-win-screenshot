package main

import (
	"fmt"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
)

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e.(type) {
			case lifecycle.Event:
				// Initialize your app here
			case paint.Event:
				// Paint your UI here
				fmt.Println("Hello, GoMobile!")
				a.Publish()
			}
		}
	})
}
