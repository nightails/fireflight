package app

import "github.com/nightails/fireflight/internal/elgato"

type errMsg error

type deviceMsg struct {
	Device elgato.Device
}
