package ui

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	

	"lunar-md/internal/assets"
	"lunar-md/internal/lunar"
	"lunar-md/internal/style"
	"lunar-md/internal/lang"
	"lunar-md/internal/config"
)

func MainWindow() {
	myApp := app.NewWithID("ru.lunarcreators.mediadirector")
	window := myApp.NewWindow("Lunar Media Director (Dev)")
	window.Resize(fyne.NewSize(1000, 500))
	window.SetIcon(assets.Get("LunarMini_ColoredLB.webp"))

	// Right side - Main content
	// Create empty stack that we can update on need
	mainContent := container.NewStack()

	// Show main screen
	showHome := func() {
		logo := assets.GetImg("LunarFull_Colored.webp", 200, 100)
		title := widget.NewLabelWithStyle(lang.Get("main.mainmenu"), fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
		desc := container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel(lang.Get("main.welcome")),
			layout.NewSpacer(),
		)

		/*
			card := widget.NewCard("FFmpeg", "Progress",
				container.NewPadded(widget.NewProgressBarInfinite()),
			)
		*/
		/* 
			If need to send OS notification:
			myApp.SendNotification(fyne.NewNotification("Info", "Config Saved"))
		*/

		mainContent.Objects = []fyne.CanvasObject{container.NewVBox(layout.NewSpacer(), logo, title, desc, layout.NewSpacer())}
		mainContent.Refresh()
	}

	showSettings := func() {
		input := widget.NewEntry()
		input.SetText(config.Current.FFmpegbin)
		input.SetPlaceHolder(lang.Get("settings.ffmpegpath"))
		input.Validator = func(s string) error {
			i, err := os.Stat(input.Text)
			if err != nil {
				return fmt.Errorf("FFmpeg binary: File not exists or Restricted!")
			}
			if i.IsDir() {
				return fmt.Errorf("FFmpeg binary: Found Directory. Expected File!")
			}
			return nil // If no err
		}
		selectlang := style.NewLanguageSelect()
		var saveBtn *widget.Button // Init button before to update its text at Success
		saveBtn = widget.NewButton(lang.Get("button.save"), func() {

			// Validate entry
			if err := input.Validate(); err != nil {
				// Dialog if error
				dialog.ShowError(err, window)
				return // Interrupt saving
			}
			config.Save()
			go style.NewFlashModalPopUp("✅ Saved!", 1, window)
			
			// Thread for button update
			go func() {
				fyne.DoAndWait( func() {
					// Change text and disable button
					saveBtn.SetText("✅ " + "Saved!")
					saveBtn.Disable()
				})

				// On 2 seconds
				time.Sleep(time.Second * 2)

				fyne.DoAndWait( func() {
					// Reset button state
					saveBtn.SetText(lang.Get("button.save"))
					saveBtn.Enable()
				})
			}()
		})
		mainContent.Objects = []fyne.CanvasObject{container.NewVBox(widget.NewLabel(lang.Get("settings.settings")), widget.NewLabel(lang.Get("settings.ffmpegpath")), input, widget.NewLabel(lang.Get("settings.selectlang")), selectlang, saveBtn)}
		mainContent.Refresh()
	}

	// Sidepanel
	mainBtn := style.NewIconButton(theme.HomeIcon(), lang.Get("button.home"), showHome)
	settingsBtn := style.NewOnlyIconButton(theme.SettingsIcon(), showSettings)

	// Buttons in a vertical row
	sidebar := container.NewVBox(
		widget.NewLabel(lang.Get("main.menu")),
		container.NewPadded(mainBtn),
		container.NewPadded(settingsBtn),
		widget.NewSeparator(),
		container.NewPadded(widget.NewButton(lang.Get("button.exit"), lunar.AppOnExit)),
	)

	// Drag & Drop Logic
	// Drop on full window
	window.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
		for _, uri := range uris {
			// Print on every file
			fmt.Println("Файл перетащен:", uri.Path())

			// Example logic when using specific extension
			if uri.Extension() == ".txt" {
				showHome() // If TXT then show homepage
			}
		}
	})

	// Building window
	// Left - sidebar, Center - mainContent
	layout := container.NewBorder(nil, nil, sidebar, nil, mainContent)

	window.SetContent(layout)

	// Start with default window
	showHome()
	window.ShowAndRun()
}
