package main

import (
	keygen "AES/KeyGen"
	operation "AES/Operations"
	"fmt"

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

			password := input.GetText()
			cypherKey := keygen.GenerateKey(password)

			stringKey, intKey := operation.Substitude(cypherKey, false)

			fmt.Fprintln(textPanel, stringKey)
			fmt.Fprintf(textPanel, "%x\n", intKey)

			// password := input.GetText()
			// roundKeys := keygen.GenerateRoundKey(password)

			// for _, element := range roundKeys {
			// 	fmt.Fprintln(textPanel, element)
			// }

			// key := keygen.GenerateKey(password)
			// // hello := keygen.GenerateRoundKey()

			// rcon := keygen.RCon(10)

			// fmt.Fprintln(textPanel, password)
			// fmt.Fprintln(textPanel, key)
			// fmt.Fprintln(textPanel, hello)
			// fmt.Fprintf(textPanel, "%x\n", rcon)

			input.SetText("")
		}
	})

	flex.AddItem(input, 0, 1, true)
	flex.AddItem(textPanel, 0, 15, false)

	window.SetRoot(flex, true)

	return window
}
