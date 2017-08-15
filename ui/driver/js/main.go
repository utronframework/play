// Package main generates the gopherjs output for the app into the assets directory of this app.

//go:generate go get -v github.com/gopherjs/gopherjs
//go:generate gopherjs build -o ../../public/js/ui_app_bundle.js

package main

import (
	"github.com/gu-io/gopherjs"
	"github.com/utronframework/play/ui"
)

func main() {
	gopherjs.NewJSDriver(ui.App)
}
