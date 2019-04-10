package system

import (
    "github.com/liborio7/npong/core"
    "github.com/liborio7/npong/component"
    "time"
)

const PlayerControlSignature = component.JoystickSignature | component.PlayerSignature

type PlayerControl struct {
}

func NewPlayerControl() *PlayerControl {
    return &PlayerControl{}
}

func (p *PlayerControl) Signature() uint64 {
    return PlayerControlSignature
}

func (p *PlayerControl) Init(e *core.Entity) {
}

func (p *PlayerControl) Update(e *core.Entity, t time.Time) {
    player := &component.Player{}
    player = e.Unpack(player).(*component.Player)

    joystick := &component.Joystick{}
    joystick = e.Unpack(joystick).(*component.Joystick)

    for k, ch := range player.Channel {
        if t := p.receiveInput(ch); t != nil {
            if cmd := joystick.Conf[k]; cmd != nil {
                cmd.Execute(e)
            }
        }
    }
}

func (p *PlayerControl) receiveInput(ch chan time.Time) *time.Time {
    select {
    case item := <-ch:
        return &item
    default:
        return nil
    }
}

func (p *PlayerControl) Render(*core.Entity, time.Time) {
}
