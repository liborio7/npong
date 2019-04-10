package system

import (
    "github.com/liborio7/npong/component"
    "github.com/liborio7/npong/core"
    "time"
)

const ScriptingSignature = component.ScriptSignature

type Scripting struct {
}

func NewScripting() *Scripting {
    return &Scripting{}
}

func (Scripting) Signature() uint64 {
    return ScriptingSignature
}

func (Scripting) Init(e *core.Entity) {
    script := &component.Script{}
    script = e.Unpack(script).(*component.Script)

    for _, fn := range script.Init {
        fn(e)
    }
}

func (Scripting) Update(e *core.Entity, t time.Time) {
    script := &component.Script{}
    script = e.Unpack(script).(*component.Script)

    for _, fn := range script.Update {
        fn(e)
    }
}

func (Scripting) Render(e *core.Entity, t time.Time) {
    script := &component.Script{}
    script = e.Unpack(script).(*component.Script)

    for _, fn := range script.Init {
        fn(e)
    }
}
