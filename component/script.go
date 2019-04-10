package component

import (
    "github.com/liborio7/npong/core"
)

const ScriptSignature = 1 << SCRIPT

type Script struct {
    Init   []func(entity *core.Entity)
    Update []func(entity *core.Entity)
    Render []func(entity *core.Entity)
}

func (Script) Signature() uint64 {
    return ScriptSignature
}
