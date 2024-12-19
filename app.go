package main

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
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

func (a *App) SplitMKV(inputPath, start, end string) (string, error) {
	if inputPath == "" {
		return "", fmt.Errorf("no file selected")
	}

	dir := filepath.Dir(inputPath)
	base := filepath.Base(inputPath)
	ext := filepath.Ext(base)
	nameOnly := strings.TrimSuffix(base, ext)
	outputPath := filepath.Join(dir, fmt.Sprintf("%s_%s-%s.mkv", nameOnly, strings.ReplaceAll(start, ":", ""), strings.ReplaceAll(end, ":", "")))

	// ffmpeg -i input.mkv -ss HH:MM:SS -to HH:MM:SS -c copy output.mkv
	cmd := exec.Command(".\\ffmpeg.exe", "-i", inputPath, "-ss", start, "-to", end, "-c", "copy", "-y", outputPath)

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error on ffmpeg: %w", err)
	}

	return outputPath, nil
}
