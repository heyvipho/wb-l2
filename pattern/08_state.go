package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

func main() {
	mobile := NewMobileAlert()

	fmt.Println(mobile.Alert())
	fmt.Println(mobile.Alert())

	mobile.SetState(&MobileAlertSong{})

	fmt.Println(mobile.Alert())
}

type MobileAlertStater interface {
	Alert() string
}

type MobileAlert struct {
	state MobileAlertStater
}

func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

type MobileAlertVibration struct {
}

func (a *MobileAlertVibration) Alert() string {
	return "Brrr... Brrr... Brrr..."
}

// MobileAlertSong implements beep alert
type MobileAlertSong struct {
}

func (a *MobileAlertSong) Alert() string {
	return "До... Ре... Ми..."
}
