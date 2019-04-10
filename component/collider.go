package component

import (
    "github.com/golang/geo/r2"
)

const ColliderSignature = 1 << COLLIDER

const (
    // Static colliders are used for level geometry which always stays at the same place and never moves around
    StaticCollider = iota
    // Rigidbody colliders are fully simulated by the physics engine and can react to collisions and forces applied
    RigidbodyCollider
    // Kinematic rigidbodies should be used for colliders that can be moved or disabled/enabled occasionally
    KinematicRigidbodyCollider
    //  A collider configured as a Trigger does not behave as a solid object and will simply allow other colliders to pass through
    TriggerCollider
)

const (
    YUpCollision = iota
    XRightCollision
    YDownCollision
    XLeftCollision
)

type Collider struct {
    *r2.Rect
    Type       int
    Friction   float64
    Bounciness float64
    Collisions []*Collision
    OnEnter    func()
    OnStay     func()
    OnExit     func()
}

type Collision struct {
    Direction int
    Type      int
}

func (j *Collider) Signature() uint64 {
    return ColliderSignature
}
