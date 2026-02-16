package style

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"

	"lunar-md/internal/config"
	"lunar-md/internal/lang"
)

/*//////////
	Flash Modal PopUp
*///////////

// Creates a new modal with custom timeout. Usage: go NewFlashModalPopUp("Its a popup", 1, window)
func NewFlashModalPopUp(label string, timeout time.Duration, window fyne.Window) {
	var pop fyne.Widget
	fyne.DoAndWait( func() {
		pop = widget.NewModalPopUp(container.NewVBox(
			widget.NewLabelWithStyle(label, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		),
			window.Canvas(),
		)
		pop.Show()
	})
	
	time.Sleep(time.Second * timeout)
    
    fyne.DoAndWait( func() {
        pop.Hide()
    })
}

/*//////////
	Button with Icon and lable. 
*///////////

// Data for NewIconButton
type IconButton struct {
	widget.BaseWidget
	icon    *canvas.Image
	label   *widget.Label
	onTap   func()
	bg      *canvas.Rectangle
	isHover bool
}

// Create new button with icon and label (iconRes, "Label", action())
func NewIconButton(iconRes fyne.Resource, text string, tapped func()) *IconButton {
	i := &IconButton{
		onTap: tapped,
		bg: canvas.NewRectangle(color.Transparent), // Фон прозрачный по умолчанию
	}
	i.bg.CornerRadius = 8 // Закругление углов

	i.icon = canvas.NewImageFromResource(iconRes)
	i.icon.FillMode = canvas.ImageFillContain
	i.icon.SetMinSize(fyne.NewSize(40, 40)) // Размер иконки

	i.label = widget.NewLabelWithStyle(text, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	i.ExtendBaseWidget(i)
	return i
}

// Render
func (i *IconButton) CreateRenderer() fyne.WidgetRenderer {
	// Build vertical icon with label
	content := container.NewVBox(i.icon, i.label)
	// Render it on graybox
	return widget.NewSimpleRenderer(container.NewStack(i.bg, content))
}

// OnClick
func (i *IconButton) Tapped(_ *fyne.PointEvent) {
	if i.onTap != nil {
		i.onTap()
	}
}

// On mouse Hover Event
func (i *IconButton) MouseIn(_ *desktop.MouseEvent) {
	i.bg.FillColor = color.RGBA{R: 255, G: 255, B: 255, A: 30} // Подсветка (белый 10%)
	i.bg.Refresh()
}

// End mouse Hover Event
func (i *IconButton) MouseOut() {
	i.bg.FillColor = color.Transparent
	i.bg.Refresh()
}

var _ desktop.Hoverable = (*IconButton)(nil)

// Required for mouse hover event
func (i *IconButton) MouseMoved(event *desktop.MouseEvent) {
    // It Just need to exist. Delete if want compilation err
}

/*//////////
	Button with Icon only. 
	Copy of IconButton without Label
*///////////

type OnlyIconButton struct {
	widget.BaseWidget
	icon    *canvas.Image
	onTap   func()
	bg      *canvas.Rectangle
	isHover bool
}

func NewOnlyIconButton(iconRes fyne.Resource, tapped func()) *OnlyIconButton {
	i := &OnlyIconButton{
		onTap: tapped,
		bg: canvas.NewRectangle(color.Transparent), // Фон прозрачный по умолчанию
	}
	i.bg.CornerRadius = 8 // Закругление углов

	i.icon = canvas.NewImageFromResource(iconRes)
	i.icon.FillMode = canvas.ImageFillContain
	i.icon.SetMinSize(fyne.NewSize(40, 40)) // Размер иконки

	i.ExtendBaseWidget(i)
	return i
}

// Отрисовка
func (i *OnlyIconButton) CreateRenderer() fyne.WidgetRenderer {
	// Build icon
	content := container.NewVBox(i.icon)
	// Render it on graybox
	return widget.NewSimpleRenderer(container.NewStack(i.bg, content))
}

// OnClick
func (i *OnlyIconButton) Tapped(_ *fyne.PointEvent) {
	if i.onTap != nil {
		i.onTap()
	}
}

// On mouse Hover Event
func (i *OnlyIconButton) MouseIn(_ *desktop.MouseEvent) {
	i.bg.FillColor = color.RGBA{R: 255, G: 255, B: 255, A: 30} // Подсветка (белый 10%)
	i.bg.Refresh()
}

// End mouse Hover Event
func (i *OnlyIconButton) MouseOut() {
	i.bg.FillColor = color.Transparent
	i.bg.Refresh()
}

var _ desktop.Hoverable = (*OnlyIconButton)(nil)

// Required for mouse hover event
func (i *OnlyIconButton) MouseMoved(event *desktop.MouseEvent) {
    // It Just need to exist. Delete if want compilation err
}

/*//////////
	Language selector for settings
*///////////

// Options for settings
var languageOptions = []string{"🇷🇺 Русский", "🇺🇸 English"}

// Map languageOptions to config strings
var langMap = map[string]string{
    "🇷🇺 Русский": "ru",
    "🇺🇸 English": "en",
}
// Map config strings to languageOptions
var reverseLangMap = map[string]string{
    "ru": "🇷🇺 Русский",
    "en": "🇺🇸 English",
}

// Create Language Selector widget
func NewLanguageSelect() *widget.Select {
    // 2. Создаем сам виджет
    selectWidget := widget.NewSelect(languageOptions, func(selected string) {
        // Set selected LangCode
        langCode := langMap[selected]
        
        // Save selection
        config.Current.Language = langCode
        
        // Apply language
        lang.LoadLanguage(langCode)
    })

    // Set current Lang as selected
    selectWidget.SetSelected(reverseLangMap[config.Current.Language])

    return selectWidget
}