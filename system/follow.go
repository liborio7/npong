package system

import (
    "github.com/liborio7/npong/component"
    "github.com/liborio7/npong/core"
    "time"
)

const FollowSignature = component.JointSignature | component.TransformSignature

type Follow struct {
}

func (Follow) Signature() uint64 {
    return FollowSignature
}

func NewFollow() *Follow {
    return &Follow{}
}

func (Follow) Init(*core.Entity) {
}

func (Follow) Update(entity *core.Entity, t time.Time) {
    transform := &component.Transform{}
    transform = entity.Unpack(transform).(*component.Transform)

    joint := &component.Joint{}
    joint = entity.Unpack(joint).(*component.Joint)

    following := &component.Transform{}
    if following = joint.Entity.Unpack(following).(*component.Transform); following != nil {
        transform.X = following.X + joint.Anchor.X
        transform.Y = following.Y + joint.Anchor.Y
    }
}

func (Follow) Render(*core.Entity, time.Time) {
}
