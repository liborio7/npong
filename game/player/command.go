package player

import (
    "github.com/liborio7/npong/component"
    "github.com/liborio7/npong/core"
)

type commandManager struct {
    Jump      core.Command
    MoveLeft  core.Command
    MoveRight core.Command
}

func newCommandManager() *commandManager {
    return &commandManager{
        Jump:      &jump{},
        MoveLeft:  &moveLeft{},
        MoveRight: &moveRight{},
    }
}

type jump struct {
    Command *core.Command
}

func (ml *jump) Execute(e *core.Entity) {
    motion := &component.Motion{}
    motion = e.Unpack(motion).(*component.Motion)
    if motion == nil {
        return
    }

    motion.AccelerationY.Value += motion.AccelerationY.Max
}

type moveLeft struct {
}

func (ml *moveLeft) Execute(e *core.Entity) {
    motion := &component.Motion{}
    motion = e.Unpack(motion).(*component.Motion)

    if motion == nil {
        return
    }

    motion.AccelerationX.Value += motion.AccelerationX.Min
}

type moveRight struct {
}

func (ml *moveRight) Execute(e *core.Entity) {
    motion := &component.Motion{}
    motion = e.Unpack(motion).(*component.Motion)

    if motion == nil {
        return
    }

    motion.AccelerationX.Value += motion.AccelerationX.Max
}
