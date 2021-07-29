package goi

import (
	"os"

	"github.com/gotk3/gotk3/gtk"
)

type Window struct {
	win   *gtk.Window
	child Component
}

func init() {
	gtk.Init(&os.Args)
}

// NewWindow makes a new window. If you want to have multiple windows in your application, pass in a boolean to multiwindow
func NewWindow(title string, multiwindow ...bool) (*Window, error) {
	w := &Window{}
	var err error
	w.win, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}
	w.win.SetTitle(title)
	if len(multiwindow) == 0 {
		w.win.Connect("destroy", func() {
			gtk.MainQuit()
		})
	}
	return w, nil
}

func (w *Window) SetDefaultSize(width, height int) {
	w.win.SetDefaultSize(width, height)
}

func (w *Window) SetChild(child Component) {
	if w.child != nil {
		w.win.Remove(w.child.getWidget())
	}
	w.child = child
	w.win.Add(w.child.getWidget())
}

func (w *Window) Show() {
	w.win.ShowAll()
}

func Run() {
	gtk.Main()
}
