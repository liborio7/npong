package system

import (
    "fmt"
    "github.com/liborio7/npong/component"
    "github.com/liborio7/npong/core"
    "time"
)

const AutomataSignature = component.AutomatonSignature

type Automata struct {
}

func NewAutomata() *Automata {
    return &Automata{}
}

func (Automata) Signature() uint64 {
    return AutomataSignature
}

func (Automata) Init(*core.Entity) {
}

func (Automata) Update(entity *core.Entity, t time.Time) {
    automaton := &component.Automaton{}
    automaton = entity.Unpack(automaton).(*component.Automaton)

    for len(automaton.TransitionStack) > 0 {
        // get last element
        n := len(automaton.TransitionStack) - 1
        newState := automaton.TransitionStack[n]

        if automaton.Fsm.IsValidTransition(automaton.CurrentState, newState) {
            fmt.Printf("change state from %v to %v\n", automaton.CurrentState, newState)
            for _, comp := range newState.Components {
                entity.AddComponents(comp)
            }
            automaton.CurrentState = newState
        }

        // pop from stack
        automaton.TransitionStack = automaton.TransitionStack[:n]
    }
}

func (Automata) Render(*core.Entity, time.Time) {
}
