package assets

import (
	"embed"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	_ "golang.org/x/image/webp"
)

//go:embed icons/*.webp *.webp
var fs embed.FS

// Search files and convert to Fyne resource.
func Get(name string) fyne.Resource {
	// Searching in ./icons
	path := "icons/" + name
	data, err := fs.ReadFile(path)
	
	// if not: Searching in root
	if err != nil {
		path = name
		data, err = fs.ReadFile(path)
	}

	if err != nil {
		fmt.Printf("Err: %s not found in binary\n", name)
		return nil
	}

	return fyne.NewStaticResource(name, data)
}

func GetImg(name string, w float32, h float32) fyne.CanvasObject {
	res := canvas.NewImageFromResource(Get(name))
	res.FillMode = canvas.ImageFillContain // Fill image in area preserving ratio
	res.SetMinSize(fyne.NewSize(w, h)) // Area size

	return res
}