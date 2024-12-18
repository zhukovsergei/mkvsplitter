package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) SelectMKVFile(name string) string {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Выберите файл",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "MKV files",
				Pattern:     "*.mkv",
			},
		},
	})
	if err != nil {
		return "Ошибка при выборе файла: " + err.Error()
	}
	if file == "" {
		return "Файл не был выбран."
	}
	return "Выбран файл: " + file
}
