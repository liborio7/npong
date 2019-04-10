package system

import (
    "github.com/golang/geo/r2"
    "github.com/liborio7/npong/core"
    "github.com/liborio7/npong/component"
    "time"
)

const PlayerPhysicsSignature = component.PlayerSignature | component.TransformSignature | component.ColliderSignature | component.MotionSignature
const groundPhysicsSignature = component.GroundSignature | component.TransformSignature | component.ColliderSignature

type PlayerPhysics struct {
}

func NewPlayerPhysics() *PlayerPhysics {
    return &PlayerPhysics{}
}

func (p *PlayerPhysics) Signature() uint64 {
    return PlayerPhysicsSignature
}

func (p *PlayerPhysics) Init(*core.Entity) {
}

func (p *PlayerPhysics) Update(e *core.Entity, t time.Time) {
    transform := &component.Transform{}
    transform = e.Unpack(transform).(*component.Transform)

    collision := &component.Collider{}
    collision = e.Unpack(collision).(*component.Collider)
    collision.Collisions = nil

    motion := &component.Motion{}
    motion = e.Unpack(motion).(*component.Motion)

    // check collisions with ground
    for _, e2 := range e.World.Entities {
        if groundPhysicsSignature&e2.Signature == groundPhysicsSignature {
            if e2.Id == e.Id {
                continue
            }

            transform2 := &component.Transform{}
            transform2 = e2.Unpack(transform2).(*component.Transform)

            collider2 := &component.Collider{}
            collider2 = e2.Unpack(collider2).(*component.Collider)

            motion2 := &component.Motion{}
            if m, ok := e2.Unpack(motion2).(*component.Motion); ok {
                motion2 = m
            }

            evaluateGroundCollision(transform, collision, motion, transform2, collider2)
        }
    }
}

func evaluateGroundCollision(t1 *component.Transform, c1 *component.Collider, m1 *component.Motion,
    t2 *component.Transform, c2 *component.Collider) {

    rect1 := r2.RectFromCenterSize(r2.Point{X: t1.X, Y: t1.Y}, r2.Point{X: c1.Size().X, Y: c1.Size().Y})
    rect2 := r2.RectFromCenterSize(r2.Point{X: t2.X, Y: t2.Y}, r2.Point{X: c2.Size().X, Y: c2.Size().Y})

    intersection := rect1.Intersection(rect2)

    if !intersection.IsEmpty() {
        collision := &component.Collision{}
        collision.Type = c2.Type
        c1.Collisions = append(c1.Collisions, collision)

        // collision function must depend on the collided entities types
        if intersection.Size().X > intersection.Size().Y {
            if rect1.Hi().Y == intersection.Hi().Y && m1.AccelerationY.Value > 0 {
                // collision a up
                collision.Direction = component.YUpCollision
                t1.Y -= intersection.Size().Y
                m1.AccelerationY.Value = 0

            } else if rect1.Lo().Y == intersection.Lo().Y && m1.AccelerationY.Value < 0 {
                // collision a down
                collision.Direction = component.YDownCollision
                t1.Y += intersection.Size().Y
                m1.AccelerationY.Value = 0
            }

        } else {
            if rect1.Hi().X == intersection.Hi().X && m1.AccelerationX.Value > 0 {
                // collision a right
                collision.Direction = component.XRightCollision
                t1.X -= intersection.Size().X
                m1.AccelerationX.Value = 0

            } else if rect1.Lo().X == intersection.Lo().X && m1.AccelerationX.Value < 0 {
                // collision a left
                collision.Direction = component.XLeftCollision
                t1.X += intersection.Size().X
                m1.AccelerationX.Value += 10
            }
        }
    }
}

func (p *PlayerPhysics) Render(*core.Entity, time.Time) {
}
