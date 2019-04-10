package player

import (
    "github.com/faiface/pixel"
    "github.com/golang/geo/r1"
    "github.com/golang/geo/r2"
    "github.com/liborio7/npong/core"
    "github.com/liborio7/npong/component"
    "image"
    "os"
    "time"
)

type componentManager struct {
    spritePic pixel.Picture
}

func newComponentManager() *componentManager {
    pic, err := loadPicture("game/player/assets/left_right.png")
    if err != nil {
        panic(err)
    }
    return &componentManager{spritePic: pic}
}

func loadPicture(path string) (pixel.Picture, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    img, _, err := image.Decode(file)
    if err != nil {
        return nil, err
    }
    return pixel.PictureDataFromImage(img), nil
}

func (m *componentManager) newPlayer() *component.Player {
    return &component.Player{
        Channel: make(map[string]chan time.Time, 0),
    }
}

func (m *componentManager) newTransform(point *r2.Point) *component.Transform {
    return &component.Transform{Point: point}
}

func (m *componentManager) newCollider() *component.Collider {
    return &component.Collider{
        Rect: &r2.Rect{
            X: r1.Interval{Hi: 50},
            Y: r1.Interval{Hi: 50},
        },
        Type: component.RigidbodyCollider,
    }
}

func (m *componentManager) newJoystick() *component.Joystick {
    // defined by state
    return &component.Joystick{
    }
}

func (m *componentManager) newSprite() *component.Sprite {
    return &component.Sprite{
        Pic:        nil,
        Frame:      &r2.Rect{},
        LastUpdate: time.Now(),
    }
}

func (m *componentManager) newMotion() *component.Motion {
    return &component.Motion{
        SpeedX: &component.MotionSpeed{
            Max: 15,
            Min: -15,
        },
        SpeedY: &component.MotionSpeed{
            Max: 15,
            Min: -15,
        },
        AccelerationX: &component.MotionAcceleration{
            Max: 15,
            Min: -15,
        },
        AccelerationY: &component.MotionAcceleration{
            Max: 15,
            Min: -15,
        },
    }
}

func (m *componentManager) newAnimator() *component.Animator {
    return &component.Animator{
        Pics: map[string]pixel.Picture{
            "idle":  m.spritePic,
            "left":  m.spritePic,
            "right": m.spritePic,
        },
        Frames: map[string]*r2.Rect{
            "idle": {
                X: r1.Interval{Hi: 54},
                Y: r1.Interval{Hi: 70},
            },
            "left": {
                X: r1.Interval{Hi: 54},
                Y: r1.Interval{Hi: 70},
            },
            "right": {
                X: r1.Interval{Hi: 54},
                Y: r1.Interval{Lo: 140, Hi: 70},
            },
        },
        SpritesNum: map[string]float64{
            "idle":  1,
            "left":  8,
            "right": 8,
        },
        Animate: func(e *core.Entity) string {
            motion := &component.Motion{}
            motion = e.Unpack(motion).(*component.Motion)

            if motion.SpeedX.Value > 0 {
                return "right"
            } else if motion.SpeedX.Value < 0 {
                return "left"
            } else {
                return "idle"
            }
        },
    }
}

func (m *componentManager) newAutomaton(fsmManager *stateManager) *component.Automaton {
    return &component.Automaton{
        Fsm: fsmManager.fsm,
    }
}

func (m *componentManager) newScript(stateManager *stateManager) *component.Script {
    jumpScript := func(entity *core.Entity) {
        automaton := &component.Automaton{}
        automaton = entity.Unpack(automaton).(*component.Automaton)

        motion := &component.Motion{}
        motion = entity.Unpack(motion).(*component.Motion)

        collider := &component.Collider{}
        collider = entity.Unpack(collider).(*component.Collider)

        if motion.AccelerationY.Value > 0 {
            automaton.TransitionStack = append(automaton.TransitionStack, stateManager.Jumping)
            return
        }

        if motion.AccelerationY.Value < 0 {
            automaton.TransitionStack = append(automaton.TransitionStack, stateManager.Falling)
            return
        }

        for _, collision := range collider.Collisions {
            if collision.Direction == component.YDownCollision && collision.Type == component.RigidbodyCollider {
                automaton.TransitionStack = append(automaton.TransitionStack, stateManager.Idle)
                return
            }
        }

        automaton.TransitionStack = append(automaton.TransitionStack, stateManager.Jumping)
        return
    }

    return &component.Script{
        Update: []func(entity *core.Entity){
            jumpScript,
        },
    }
}
