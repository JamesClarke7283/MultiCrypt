package frontend

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/JamesClarke7283/MultiCrypt/src/shared"
)

func CreateMainContent(a fyne.App, w fyne.Window, config *shared.Config, encryptFunc, decryptFunc func(string, string) (string, error)) fyne.CanvasObject {
	logger := shared.GetLogger()

	keyEntry := widget.NewPasswordEntry()
	keyEntry.SetPlaceHolder("Enter key")

	inputEntry := widget.NewMultiLineEntry()
	inputEntry.Wrapping = fyne.TextWrapWord
	inputEntry.SetMinRowsVisible(10)

	outputEntry := widget.NewMultiLineEntry()
	outputEntry.Wrapping = fyne.TextWrapWord
	outputEntry.SetMinRowsVisible(10)

	inputLabel := widget.NewLabel("Message Input")
	outputLabel := widget.NewLabel("Ciphertext Output")

	encryptDecryptToggle := widget.NewSelect([]string{"Encrypt", "Decrypt"}, func(selected string) {
		if selected == "Encrypt" {
			inputLabel.SetText("Message Input")
			outputLabel.SetText("Ciphertext Output")
			inputEntry.SetPlaceHolder("Enter message")
			outputEntry.SetPlaceHolder("Ciphertext")
		} else {
			inputLabel.SetText("Ciphertext Input")
			outputLabel.SetText("Plaintext Output")
			inputEntry.SetPlaceHolder("Enter ciphertext")
			outputEntry.SetPlaceHolder("Plaintext")
		}
		outputEntry.SetText("") // Clear output when switching modes
	})
	encryptDecryptToggle.SetSelected("Encrypt")

	processButton := widget.NewButton("Process", func() {
		if encryptDecryptToggle.Selected == "Encrypt" {
			ciphertext, err := encryptFunc(keyEntry.Text, inputEntry.Text)
			if err != nil {
				logger.Errorf("Encryption failed: %v", err)
				outputEntry.SetText("Encryption failed: " + err.Error())
			} else {
				outputEntry.SetText(ciphertext)
			}
		} else {
			plaintext, err := decryptFunc(keyEntry.Text, inputEntry.Text)
			if err != nil {
				logger.Errorf("Decryption failed: %v", err)
				outputEntry.SetText("Decryption failed: " + err.Error())
			} else {
				outputEntry.SetText(plaintext)
			}
		}
	})

	copyButton := widget.NewButtonWithIcon("Copy to Clipboard", theme.ContentCopyIcon(), func() {
		w.Clipboard().SetContent(outputEntry.Text)
	})

	generateKeyButton := widget.NewButtonWithIcon("Generate Random Key", theme.MediaReplayIcon(), func() {
		showGenerateKeyDialog(w, keyEntry)
	})

	settingsButton := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
		ShowSettingsDialog(a, w, config)
	})

	keySection := container.NewVBox(
		container.NewHBox(widget.NewLabel("Key:"), layout.NewSpacer(), settingsButton),
		keyEntry,
		generateKeyButton,
	)

	content := container.NewVBox(
		keySection,
		encryptDecryptToggle,
		inputLabel,
		container.NewScroll(inputEntry),
		outputLabel,
		container.NewScroll(outputEntry),
		container.NewHBox(layout.NewSpacer(), processButton, copyButton, layout.NewSpacer()),
	)

	return container.NewPadded(container.NewPadded(content))
}

func showGenerateKeyDialog(w fyne.Window, keyEntry *widget.Entry) {
	lengthEntry := widget.NewEntry()
	lengthEntry.SetText("32")

	charsetEntry := widget.NewEntry()
	charsetEntry.SetText("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:,.<>?")

	content := container.NewVBox(
		widget.NewLabel("Key Length:"),
		lengthEntry,
		widget.NewLabel("Character Set:"),
		charsetEntry,
	)

	dialog.ShowCustomConfirm("Generate Random Key", "Generate", "Cancel", content, func(generate bool) {
		if generate {
			length := 32
			fmt.Sscanf(lengthEntry.Text, "%d", &length)
			charset := charsetEntry.Text
			key := generateRandomKey(length, charset)
			keyEntry.SetText(key)
		}
	}, w)
}

func generateRandomKey(length int, charset string) string {
	key := make([]byte, length)
	for i := range key {
		key[i] = charset[rand.Intn(len(charset))]
	}
	return string(key)
}
