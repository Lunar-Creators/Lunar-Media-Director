// Package for managing and loading .toml configuration strings
package config

import (
	"os"
	"runtime"
	"path/filepath"
	"github.com/BurntSushi/toml"
)

// Application options
type Cfg struct {
	Language string `toml:"language"`
	FFmpegbin string `toml:"ffmpegbin"`
}

// Current applied app configuration
var Current Cfg
var configfile string
// FFmpeg Executable
var Ffmpeg string
// Data about OS type
var OS = runtime.GOOS

// Load settings on start
func init() {
	configfile = getAppPath("config.toml")
	Load()
}

// Load settings from file
func Load() {
	// Default settings if not found
	if _, err := os.Stat(configfile); os.IsNotExist(err) {
		Current.Language = "ru"
		Current.FFmpegbin = ""
		Save()
		return
	}

	// Load settings
	toml.DecodeFile(configfile, &Current)

	// Map OS to FFmpeg binary name
	var osbinmap = map[string]string{
    	"linux": getAppPath("ffmpeg"),
    	"windows": getAppPath("ffmpeg.exe"),
	}
	
	// Decide path to FFmpeg binary
	if Current.FFmpegbin == "" {
		Ffmpeg = osbinmap[OS]
	} else {
		Ffmpeg = Current.FFmpegbin
	}
}

// Save Current settings
func Save() {
	f, _ := os.Create(configfile)
	defer f.Close()
	toml.NewEncoder(f).Encode(Current)
}

func getAppPath(filename string) string {
	// Get executable path
	exePath, err := os.Executable()
	if err != nil {
		return filename // Fallback to workdir on fail
	}

	// Get only directory (Exported)
	ExeDir := filepath.Dir(exePath)

	// Dir + Config file = Config Path
	return filepath.Join(ExeDir, filename)
}