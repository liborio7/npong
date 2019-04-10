package system

import (
    "github.com/faiface/pixel"
    "github.com/faiface/pixel/pixelgl"
    "github.com/liborio7/npong/component"
    "github.com/liborio7/npong/core"
    "time"
)

const GraphicsSignature = component.TransformSignature | component.SpriteSignature
const cameraSignature = component.TransformSignature | component.WindowSignature

type Graphics struct {
}

func NewGraphics() *Graphics {
    return &Graphics{}
}

func (ms *Graphics) Signature() uint64 {
    return GraphicsSignature
}

func (ms *Graphics) Init(*core.Entity) {
}

func (ms *Graphics) Update(*core.Entity, time.Time) {
}

func (ms *Graphics) Render(e *core.Entity, t time.Time) {
    transform := &component.Transform{}
    transform = e.Unpack(transform).(*component.Transform)

    sprite := &component.Sprite{}
    sprite = e.Unpack(sprite).(*component.Sprite)

    frameMinX := sprite.Frame.Lo().X + (sprite.Frame.Hi().X * sprite.FrameIndex)
    frameMaxX := sprite.Frame.Hi().X + frameMinX
    spriteVec := pixel.R(frameMinX, sprite.Frame.Lo().Y, frameMaxX, sprite.Frame.Hi().Y)

    // cameras/windows must be separated from other entities? a player should be able to edit its own camera/window
    for _, e2 := range e.World.Entities {
        if cameraSignature&e2.Signature == cameraSignature {
            if e2.Id == e.Id {
                continue
            }

            cameraWindow := &component.Window{}
            cameraWindow = e2.Unpack(cameraWindow).(*component.Window)

            cameraTransform := &component.Transform{}
            cameraTransform = e2.Unpack(cameraTransform).(*component.Transform)

            windowVec := pixel.Vec{X: cameraTransform.X - cameraWindow.Bounds().W()/2, Y: cameraTransform.Y - cameraWindow.Bounds().H()/2}
            entityVec := pixel.Vec{X: transform.X - windowVec.X, Y: transform.Y - windowVec.Y}

            if sprite.Pic != nil {
                newSprite := pixel.NewSprite(sprite.Pic, spriteVec)
                newSprite.Draw(cameraWindow.Window, pixel.IM.Scaled(windowVec, cameraWindow.Scale).Moved(entityVec))

            } else if sprite.RGBA != nil {
                newCanvas := pixelgl.NewCanvas(spriteVec)
                newCanvas.Clear(sprite.RGBA)
                newCanvas.DrawColorMask(cameraWindow.Window, pixel.IM.Scaled(windowVec, cameraWindow.Scale).Moved(entityVec), sprite.RGBA)
            }
        }
    }
}
