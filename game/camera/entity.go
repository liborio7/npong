package camera

import (
    "github.com/faiface/pixel/pixelgl"
    "github.com/liborio7/npong/core"
)

type entityManager struct {
}

func newEntityManager() *entityManager {
    return &entityManager{}
}

func (em *entityManager) newEntity(m *Manager, win *pixelgl.Window, jointEntity *core.Entity) *core.Entity {
    joint := m.ComponentManager.newJoint(jointEntity)
    transform := m.ComponentManager.newTransform()
    window := m.ComponentManager.newWindow(win)

    entity := core.NewEntity()
    entity.AddTags("camera")
    entity.AddComponents(
        joint,
        transform,
        window,
    )

    return entity
}
