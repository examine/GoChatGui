//package GoChatPractice

package main

import (
		"fmt"
		ui "github.com/VladimirMarkelov/clui"
)

func createView () *TextReader {

	view := ui.AddWindow(0, 0, 10, 7, "TableView")
	bch := ui.CreateTextReader(view, 45, 24, 1)
	ui.ActivateControl(view, bch)

	return bch
}


func mainLoop () {
	ui.InitLibrary()
	defer ui.DeinitLibrary()

	view := ui.AddWindow(0, 0, 10, 7, "Hello World!")
	view := SetPack(ui.Vertical)

	btnQuit := ui.CreateButton(view, 15, 4, "Exit", 1)
	btnQuit.OnClick(func(ev ui.Event) {
		go ui.Stop()
	})

	ui.MainLoop()
}

func main () {
	mainLoop()
}

