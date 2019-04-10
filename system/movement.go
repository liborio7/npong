package system

import (
    "github.com/liborio7/npong/component"
    "github.com/liborio7/npong/core"
    "time"
)

const MovementSignature = component.MotionSignature | component.TransformSignature

type Movement struct {
}

func NewMovement() *Movement {
    return &Movement{}
}

func (ms *Movement) Signature() uint64 {
    return MovementSignature
}

func (ms *Movement) Init(*core.Entity) {
}

func (ms *Movement) Update(entity *core.Entity, t time.Time) {
    transform := &component.Transform{}
    transform = entity.Unpack(transform).(*component.Transform)

    motion := &component.Motion{}
    motion = entity.Unpack(motion).(*component.Motion)

    ms.updateMotion(motion)
    ms.updateTransform(transform, motion)
}

func (ms *Movement) Render(*core.Entity, time.Time) {
}

// ----------

func (ms *Movement) updateMotion(motion *component.Motion) {

    ms.updateSpeed(motion.SpeedX, motion.AccelerationX)
    ms.updateSpeed(motion.SpeedY, motion.AccelerationY)

    // ms.updateAcceleration(motion.AccelerationX)
    // ms.updateAcceleration(motion.AccelerationY)
    motion.AccelerationX.Value = 0
}

func (ms *Movement) updateSpeed(speed *component.MotionSpeed, acceleration *component.MotionAcceleration) {
    if acceleration.Value > acceleration.Max {
        acceleration.Value = acceleration.Max
    }
    if acceleration.Value < acceleration.Min {
        acceleration.Value = acceleration.Min
    }

    // update this function to get a better movement
    speed.Value = acceleration.Value

    if speed.Value > speed.Max {
        speed.Value = speed.Max
    }
    if speed.Value < speed.Min {
        speed.Value = speed.Min
    }
}

func (ms *Movement) updateAcceleration(acceleration *component.MotionAcceleration) {
    // update this function to get a better movement
    acceleration.Value = 0
}

func (ms *Movement) updateTransform(transform *component.Transform, motion *component.Motion) {
    // update this function to get a better movement
    transform.X += motion.SpeedX.Value
    transform.Y += motion.SpeedY.Value
}
