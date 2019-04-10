package system

import (
    "github.com/liborio7/npong/core"
    "github.com/liborio7/npong/component"
    "time"
)

const AnimationSignature = component.AnimatorSignature | component.SpriteSignature

type Animation struct {
}

func NewAnimation() *Animation {
    return &Animation{}
}

func (Animation) Signature() uint64 {
    return AnimationSignature
}

func (Animation) Init(*core.Entity) {
}

func (Animation) Update(*core.Entity, time.Time) {
}

func (Animation) Render(entity *core.Entity, t time.Time) {
    animator := &component.Animator{}
    animator = entity.Unpack(animator).(*component.Animator)

    sprite := &component.Sprite{}
    sprite = entity.Unpack(sprite).(*component.Sprite)

    picId := animator.Animate(entity)

    if picId != sprite.PicId {
        pic := animator.Pics[picId]
        frame := animator.Frames[picId]
        // handle durations

        sprite.Pic = pic
        sprite.PicId = picId
        sprite.Frame = frame
        sprite.FrameIndex = 0
        sprite.LastUpdate = t

    } else {
        n := animator.SpritesNum[picId]

        if t.Sub(sprite.LastUpdate) > time.Millisecond*50 {
            sprite.FrameIndex += 1
            if sprite.FrameIndex >= n {
                sprite.FrameIndex = 0
            }
            sprite.LastUpdate = t
        }
    }
}
