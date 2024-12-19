package main

import (
	"context"
	"fmt"
	"os"
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

	outputMKV := filepath.Join(dir, fmt.Sprintf("%s_%s-%s.mkv", nameOnly, strings.ReplaceAll(start, ":", ""), strings.ReplaceAll(end, ":", "")))
	outputMP4 := filepath.Join(dir, fmt.Sprintf("%s_%s-%s.mp4", nameOnly, strings.ReplaceAll(start, ":", ""), strings.ReplaceAll(end, ":", "")))

	cmdCut := exec.Command(".\\ffmpeg.exe", "-i", inputPath, "-ss", start, "-to", end, "-c", "copy", "-y", outputMKV)
	if err := cmdCut.Run(); err != nil {
		return "", fmt.Errorf("error cutting mkv: %w", err)
	}

	cmdConvert := exec.Command(".\\ffmpeg.exe", "-i", outputMKV, "-c", "copy", "-y", outputMP4)
	if err := cmdConvert.Run(); err != nil {
		return "", fmt.Errorf("error converting to mp4: %w", err)
	}

	if err := os.Remove(outputMKV); err != nil {
		fmt.Printf("Can't remove temp MKV file: %v\n", err)
	}

	return outputMP4, nil
}
