package main

import (
	"context"
	"fmt"
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
	// go a.generateRandomNumbers()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) generateRandomNumbers()  string{
    return fmt.Sprintf("Hello , It's show time!")
    // ticker := time.NewTicker(500 * time.Millisecond)
    // defer ticker.Stop()

    // for {
    //     select {
    //     case <-a.ctx.Done():
    //         return
    //     case <-ticker.C:
    //         randomNumber := rand.Intn(100) // Menghasilkan nilai acak antara 0 dan 99
    //         runtime.EventsEmit(a.ctx, "randomNumber", randomNumber)
    //     }
    // }
}