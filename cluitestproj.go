package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)
// Main Function.
func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	// SetManagerFunc for using layouts
	g.SetManagerFunc(layout)
	g.SetManagerFunc(layout2)

	// Keybind ctrl+C to quit
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

//func layout(g *gocui.Gui) error {
//	maxX, maxY := g.Size()
//	if v, err := g.SetView("welcome", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
//		if err != gocui.ErrUnknownView {
//			return err
//		}
//		fmt.Fprintln(v, "Welcome to GoChat!")
//	}
//	return nil
//}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("welcome", maxX/2-7, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Welcome to GoChat!")
	}
	return nil
}

func layout2(g *gocui.Gui) error {
	if v, err := g.SetView("viewname", 2, 2, 22, 7); err != nil {
		if err != gocui.ErrUnknownView {
			// handle error
		}
		fmt.Fprintln(v, "This is a new view")
		// ...
	}
	return nil
}


func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
