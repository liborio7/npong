package player

import (
    "github.com/faiface/pixel/pixelgl"
    "github.com/golang/geo/r2"
    "github.com/liborio7/npong/core"
    "github.com/liborio7/npong/component"
    "time"
)

type entityManager struct {
}

func newEntityManager() *entityManager {
    return &entityManager{}
}

func (em *entityManager) newEntity(m *Manager, win *pixelgl.Window, point *r2.Point) *core.Entity {
    player := m.ComponentManager.newPlayer()
    transform := m.ComponentManager.newTransform(point)
    collision := m.ComponentManager.newCollider()
    sprite := m.ComponentManager.newSprite()
    motion := m.ComponentManager.newMotion()
    animator := m.ComponentManager.newAnimator()
    joystick := m.ComponentManager.newJoystick()
    automaton := m.ComponentManager.newAutomaton(m.StateManager)
    script := m.ComponentManager.newScript(m.StateManager)

    entity := core.NewEntity()
    entity.AddComponents(
        player,
        transform,
        collision,
        sprite,
        motion,
        animator,
        joystick,
        automaton,
        script,
    )

    go func() {
        for !win.Closed() {
            if win.Pressed(pixelgl.KeyLeft) {
                em.publishInput(entity, pixelgl.KeyLeft)
            }
            if win.Pressed(pixelgl.KeyRight) {
                em.publishInput(entity, pixelgl.KeyRight)
            }
            if win.JustPressed(pixelgl.KeyUp) {
                em.publishInput(entity, pixelgl.KeyUp)
            }
            if win.JustPressed(pixelgl.KeyDown) {
                em.publishInput(entity, pixelgl.KeyDown)
            }
        }
    }()

    return entity
}

func (em *entityManager) publishInput(e *core.Entity, btn pixelgl.Button) {
    player := &component.Player{}
    player = e.Unpack(player).(*component.Player)

    ch := player.Channel[btn.String()]
    if ch == nil {
        ch = make(chan time.Time, 1)
        player.Channel[btn.String()] = ch
    }
    select {
    case ch <- time.Now():
    default:
    }
}
