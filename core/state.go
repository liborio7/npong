package core

type State struct {
    Name       string
    Components []Component
}

func NewState(name string, components []Component) *State {
    return &State{Name: name, Components: components}
}