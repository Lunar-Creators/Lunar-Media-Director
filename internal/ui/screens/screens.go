// Package that Lazy-Load and send ready app screens objects
package screens

import (
	"os"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"lunar-md/internal/config"
	"lunar-md/internal/lang"
	"lunar-md/internal/style"
	"lunar-md/internal/assets"
)

/*//////////
	MAIN SCREEN
*///////////

// Main Screen Data
type MainScreen struct {
	Container *fyne.Container
	Logo fyne.CanvasObject
	Title *widget.Label
	Desc *fyne.Container
	window fyne.Window
}

// Data pointer
var MainData *MainScreen

func GetMain(w fyne.Window) *MainScreen {
	if MainData == nil {
		// And init window if it is first click
		MainData = createMain(w)
	}
	return MainData
}

// Init Settings Screen
func createMain(w fyne.Window) *MainScreen {
	// Read current Window
	m := &MainScreen{window: w}

	m.Logo = assets.GetImg("LunarFull_Colored.webp", 200, 100)
	m.Title = widget.NewLabelWithStyle(lang.Get("main.mainmenu"), fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	m.Desc = container.NewHBox(
		layout.NewSpacer(),
		widget.NewLabel(lang.Get("main.welcome")),
		layout.NewSpacer(),
	)

	m.Container = container.NewVBox(layout.NewSpacer(), m.Logo, m.Title, m.Desc, layout.NewSpacer())
	return m
}

/*//////////
	SETTINGS SCREEN
*///////////

// Settings Screen Data
type SettingsScreen struct {
	Container *fyne.Container
	Input *widget.Entry
	SaveBtn *widget.Button
	window fyne.Window
}

// Data pointer
var SettingsData *SettingsScreen

// Get Settings screen
func GetSettings(w fyne.Window) *SettingsScreen {
	if SettingsData == nil {
		// And init window if it is first click
		SettingsData = createSettings(w)
	}
	return SettingsData
}

// Init Settings Screen
func createSettings(w fyne.Window) *SettingsScreen {
	// Read current Window
	s := &SettingsScreen{window: w}

	// Setup input field for FFmpeg path option
	s.Input = widget.NewEntry()
	s.Input.SetText(config.Current.FFmpegbin)
	s.Input.SetPlaceHolder(lang.Get("settings.ffmpegpath"))
	// It requires a string. Return error if validation failed.
	s.Input.Validator = func(_ string) error {
		i, err := os.Stat(s.Input.Text)
		if err != nil {
			return fmt.Errorf("FFmpeg binary: File not exists or Restricted!")
		}
		if i.IsDir() {
			return fmt.Errorf("FFmpeg binary: Found Directory. Expected File!")
		}
		return nil // If no err
	}

	// Language select widget
	selectlang := style.NewLanguageSelect()

	// Button to validate and save settings
	s.SaveBtn = widget.NewButton(lang.Get("button.save"), func() {
		// Validate entry
		if err := s.Input.Validate(); err != nil {
			// Dialog if error
			dialog.ShowError(err, w)
			return // Interrupt saving
		}
		config.Save()

		// Send modal
		go style.NewFlashModalPopUp("✅ Saved!", 1, w)
		
		// Function for button update
		go style.NewFlashButtonUpdate(s.SaveBtn, lang.Get("button.save"), "✅ Saved!", 2, true)
	})
	s.Container = container.NewVBox(widget.NewLabel(lang.Get("settings.settings")), widget.NewLabel(lang.Get("settings.ffmpegpath")), s.Input, widget.NewLabel(lang.Get("settings.selectlang")), selectlang, s.SaveBtn)
	return s
}