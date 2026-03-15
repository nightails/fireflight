package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/joho/godotenv"
	"github.com/nightails/fireflight/internal/app"
)

const (
	infoURL   = "/elgato/accessory-info"
	lightsURL = "/elgato/lights"
)

func main() {
	// Disable loading env for now
	//if err := godotenv.Load(); err != nil {
	//	fmt.Println("Error loading .env file")
	//}

	// test IP from .env
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")

	// url example
	_ = fmt.Sprintf("http://%s:%s%s", ip, port, lightsURL)

	m := app.NewModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Panicf("Program exited with error: %v", err)
	}
}
