package player

import (
    "github.com/faiface/pixel/pixelgl"
    "github.com/liborio7/npong/component"
    "github.com/liborio7/npong/core"
)

type stateManager struct {
    fsm     *core.Fsm
    Idle    *core.State
    Jumping *core.State
    Falling *core.State
}

func newStateManager(commandManager *commandManager) *stateManager {
    idle := core.NewState("idle", []core.Component{
        &component.Joystick{Conf: map[string]core.Command{
            pixelgl.KeyLeft.String():  commandManager.MoveLeft,
            pixelgl.KeyRight.String(): commandManager.MoveRight,
            pixelgl.KeyUp.String():    commandManager.Jump,
        }},
    })
    jumping := core.NewState("jumping", []core.Component{
        &component.Joystick{Conf: map[string]core.Command{
            pixelgl.KeyLeft.String():  commandManager.MoveLeft,
            pixelgl.KeyRight.String(): commandManager.MoveRight,
        }},
    })
    falling := core.NewState("falling", []core.Component{
        &component.Joystick{Conf: map[string]core.Command{
            pixelgl.KeyLeft.String():  commandManager.MoveLeft,
            pixelgl.KeyRight.String(): commandManager.MoveRight,
        }},
    })

    fsm := core.NewFsm()
    fsm.AddStates(idle, jumping, falling)

    fsm.AddTransitions(idle.Name,
        jumping.Name, falling.Name)
    fsm.AddTransitions(jumping.Name,
        falling.Name)
    fsm.AddTransitions(falling.Name,
        idle.Name)

    return &stateManager{fsm: fsm, Idle: idle, Jumping: jumping, Falling: falling}
}
