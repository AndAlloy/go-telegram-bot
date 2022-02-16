package main

import (
	tele "gopkg.in/telebot.v3"
	"log"
)

type State struct {
	Name         string
	StatusActive bool
	Message      tele.Context
}

func SetStateContext(state *State, ctx tele.Context) {
	if IsOn(state) {
		state.Message = ctx
	} else {
		log.Fatalf("State %s is not activated to save context", state.Name)
	}

}

func IsOn(state *State) bool {
	return state.StatusActive
}

func Enable(state *State) {
	state.StatusActive = true
}

func Disable(state *State) {
	state.StatusActive = false
}

func CreateState(name string) *State {
	return &State{
		Name:         name,
		StatusActive: false,
	}
}
