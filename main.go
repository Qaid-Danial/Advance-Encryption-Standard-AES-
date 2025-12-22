package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var ()

func main() {

	app := createWindow()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func createWindow() *tview.Application {
	window := tview.NewApplication().EnableMouse(true)

	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	input := tview.NewInputField()
	input.SetLabel("Password: ")

	textPanel := tview.NewTextView()
	textPanel.SetBorder(true)

	input.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {

			// password := input.GetText()
			// key := keygen.GenerateKey(password)
			// // hello := keygen.GenerateRoundKey()

			// rcon := keygen.RCon(10)

			// fmt.Fprintln(textPanel, password)
			// fmt.Fprintln(textPanel, key)
			// fmt.Fprintln(textPanel, hello)
			// fmt.Fprintf(textPanel, "%x\n", rcon)

			// cypherKey, rotBytes, subBytes := keygen.GenerateRoundKey(password, 10)
			// fmt.Fprintln(textPanel, password)
			// fmt.Fprintln(textPanel, cypherKey)
			// fmt.Fprintln(textPanel, rotBytes)
			// fmt.Fprintln(textPanel, subBytes)

			input.SetText("")
		}
	})

	flex.AddItem(input, 0, 1, true)
	flex.AddItem(textPanel, 0, 15, false)

	window.SetRoot(flex, true)

	return window
}
