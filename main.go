package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/nightails/fireflight/internal/app"
)

const (
	port      = ":9123"
	infoURL   = "/elgato/accessory-info"
	lightsURL = "/elgato/lights"
)

func main() {
	// Disable loading env for now
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// test IP from .env
	ip := os.Getenv("IP")

	// url example
	url := fmt.Sprintf("http://%s%s%s", ip, port, infoURL)

	m := app.NewModel(url)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Panicf("Program exited with error: %v", err)
	}
}
