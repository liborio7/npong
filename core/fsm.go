package core

type Fsm struct {
    States      map[string]*State
    Transitions map[string][]string
}

func NewFsm() *Fsm {
    return &Fsm{States: make(map[string]*State, 0), Transitions: make(map[string][]string, 0)}
}

func (fsm *Fsm) AddStates(states ...*State) {
    for _, state := range states {
        fsm.States[state.Name] = state
    }
}

func (fsm *Fsm) AddTransitions(from string, to ...string) {
    if m := fsm.Transitions[from]; m == nil {
        fsm.Transitions[from] = make([]string, 0)
    }
    fsm.Transitions[from] = append(fsm.Transitions[from], to...)
}

func (fsm *Fsm) IsValidTransition(from, to *State) bool {
    if from == nil {
        return true
    }
    if to == nil {
        return false
    }
    if from == to {
        return false
    }

    for _, transition := range fsm.Transitions[from.Name] {
        if transition == to.Name {
            return true
        }
    }
    return false
}
