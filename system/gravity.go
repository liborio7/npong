package system

import (
    "github.com/liborio7/npong/component"
    "github.com/liborio7/npong/core"
    "time"
)

const GravitySignature = component.TransformSignature | component.ColliderSignature | component.MotionSignature

type Gravity struct {
}

func NewGravity() *Gravity {
    return &Gravity{}
}

func (p *Gravity) Signature() uint64 {
    return GravitySignature
}

func (p *Gravity) Init(*core.Entity) {
}

func (p *Gravity) Update(e *core.Entity, t time.Time) {
    transform := &component.Transform{}
    transform = e.Unpack(transform).(*component.Transform)

    collision := &component.Collider{}
    collision = e.Unpack(collision).(*component.Collider)

    motion := &component.Motion{}
    motion = e.Unpack(motion).(*component.Motion)

    motion.AccelerationY.Value -= 1
    if motion.AccelerationY.Value < motion.AccelerationY.Min {
        motion.AccelerationY.Value = motion.AccelerationY.Min
    }
}

func (p *Gravity) Render(*core.Entity, time.Time) {
}
