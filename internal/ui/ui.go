// Main package for UI in Lunar Media Director
package ui

import (
	"fmt"
	// "runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"lunar-md/internal/assets"
	"lunar-md/internal/lang"
	"lunar-md/internal/lunar"
	"lunar-md/internal/style"
	"lunar-md/internal/ui/screens"
)

func MainWindow() {
	myApp := app.NewWithID("ru.lunarcreators.mediadirector")
	Window := myApp.NewWindow("Lunar Media Director (Dev)")
	Window.Resize(fyne.NewSize(1000, 500))
	Window.SetIcon(assets.Get("LunarMini_ColoredLB.webp"))

	// Right side - Main content
	// Create empty stack that we can update on need
	mainContent := container.NewStack()

/*
	card := widget.NewCard("FFmpeg", "Progress",
		container.NewPadded(widget.NewProgressBarInfinite()),
	)
*/
/* 
	If need to send OS notification:
	myApp.SendNotification(fyne.NewNotification("Info", "Config Saved"))
*/

	// Show main tab
	showHome := func() {
		m := screens.GetMain(Window)
		ContentRender(mainContent, []fyne.CanvasObject{m.Container})
	}

	// Show settings tab
	showSettings := func() {
		s := screens.GetSettings(Window)
		ContentRender(mainContent, []fyne.CanvasObject{s.Container})
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
	Window.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
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

	Window.SetContent(layout)

	// Start with default window
	showHome()
	Window.ShowAndRun()
}

// Render objects in content stack
func ContentRender(content *fyne.Container, objects []fyne.CanvasObject) {
	content.Objects = nil
	content.Objects = objects
	content.Refresh()
	// runtime.GC()
}