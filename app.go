package main

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) SelectMKVFile() string {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "MKV files",
				Pattern:     "*.mkv",
			},
		},
	})
	if err != nil {
		return "Error on file select: " + err.Error()
	}

	if file == "" {
		return "File was not selected."
	}

	return file
}

func (a *App) SplitMKV(inputPath string, startSec, endSec int) (string, error) {
	if inputPath == "" {
		return "", fmt.Errorf("no file selected")
	}

	dir := filepath.Dir(inputPath)
	base := filepath.Base(inputPath)
	ext := filepath.Ext(base)
	nameOnly := strings.TrimSuffix(base, ext)
	outputPath := filepath.Join(dir, fmt.Sprintf("%s_%d-%d.mkv", nameOnly, startSec, endSec))

	// ffmpeg -i input.mkv -ss startSec -to endSec -c copy output.mkv
	cmd := exec.Command(".\\ffmpeg.exe", "-i", inputPath, "-ss", strconv.Itoa(startSec), "-to", strconv.Itoa(endSec), "-c", "copy", outputPath)

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error on ffmpeg: %w", err)
	}

	return outputPath, nil
}
