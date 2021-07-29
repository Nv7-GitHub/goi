package goi

import "github.com/gotk3/gotk3/gtk"

type Component interface {
	getWidget() gtk.IWidget
}
