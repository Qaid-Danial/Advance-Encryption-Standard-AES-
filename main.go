package main

import (
	subbytes "AES/SubBytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"unsafe"

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

func generateKey(plaintext string) string {

	hasher := md5.New()
	io.WriteString(hasher, plaintext)

	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	return hashString
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
			key := generateKey(password)

			subkey := subbytes.Substitude(key)

			keySize := unsafe.Sizeof(key) * 8
			subkeySize := unsafe.Sizeof(subkey) * 8

			fmt.Fprintln(textPanel, password)
			fmt.Fprintln(textPanel, key)
			fmt.Fprintln(textPanel, keySize)
			fmt.Fprintln(textPanel, subkey)
			fmt.Fprintln(textPanel, subkeySize)

			input.SetText("")
		}
	})

	flex.AddItem(input, 0, 1, true)
	flex.AddItem(textPanel, 0, 15, false)

	window.SetRoot(flex, true)

	return window
}
