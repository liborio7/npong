package component

import (
    "github.com/liborio7/npong/core"
)

const AutomatonSignature = 1 << AUTOMATON

type Automaton struct {
    Fsm             *core.Fsm
    CurrentState    *core.State
    TransitionStack []*core.State
}

func (Automaton) Signature() uint64 {
    return AutomatonSignature
}
