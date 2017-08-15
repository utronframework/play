// Package ui is an auto-generated package which exposes the Gu.NApp which
// can be created to use the constructed views if any. Edit as you see fit.

//go:generate go run settings_bundle.go
//go:generate go run public_bundle.go
//go:generate go generate ./driver/...

package ui

import (
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/router"
	"github.com/gu-io/gu/router/cache/memorycache"
)

// Contains the projects *NApp instance and *Router level instances.
var (
	AppRouter = router.NewRouter(nil, memorycache.New("ui"))
	App       = gu.App("Ui", AppRouter)
)
