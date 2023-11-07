package router

import (
	"main.go/router/overwrite"
)

type RouterGroup struct {
	OverWrite overwrite.OverWrteRouterGroup
}

var RouterGroupApp = new(RouterGroup)
