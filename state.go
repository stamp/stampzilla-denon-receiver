package main

type State struct {
	MainZone MainZone
	Zone2    Zone2
}

type MainZone struct {
	Power  bool
	Source string
	Volume float64
}

type Zone2 struct {
	Power  bool
	Source string
	Volume float64
}

func NewState() *State {
	return &State{}
}
