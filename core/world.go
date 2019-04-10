package core

import (
    "github.com/satori/go.uuid"
    "time"
)

type World struct {
    Entities  map[uuid.UUID]*Entity
    Systems   []System
    Observers map[uint64][]Observer
    // add mutex for entities, systems, observers
}

func NewWorld() *World {
    return &World{
        Entities:  make(map[uuid.UUID]*Entity, 0),
        Systems:   make([]System, 0),
        Observers: make(map[uint64][]Observer, 0),
    }
}

func (w *World) AddEntities(es ...*Entity) {
    for _, e := range es {
        e.World = w
        w.Entities[e.Id] = e
    }
}

func (w *World) RemoveEntities(es ...*Entity) {
    for _, e := range es {
        delete(w.Entities, e.Id)
    }
}

func (w *World) AddSystems(ss ...System) {
    for _, s := range ss {
        w.Systems = append(w.Systems, s)
    }
}

func (w *World) RemoveSystems(ss ...System) {
    // todo
}

func (w *World) AddObservers(os ...Observer) {
    for _, o := range os {
        w.Observers[o.Signature()] = append(w.Observers[o.Signature()], o)
    }
}

func (w *World) RemoveObservers(os ...Observer) {
    // todo
}

func (w *World) NotifyObservers(e *Entity, signature uint64) {
    for _, o := range w.Observers[signature] {
        o.OnEvent(e)
    }
}

func (w *World) Init() {
    // can be multithreaded?
    for _, s := range w.Systems {
        for _, e := range w.Entities {
            if s.Signature()&e.Signature == s.Signature() {
                s.Init(e)
            }
        }
    }
}

func (w *World) Update(t time.Time) {
    for _, s := range w.Systems {
        for _, e := range w.Entities {
            if s.Signature()&e.Signature == s.Signature() {
                s.Update(e, t)
            }
        }
    }
}

func (w *World) Render(t time.Time) {
    for _, s := range w.Systems {
        for _, e := range w.Entities {
            if s.Signature()&e.Signature == s.Signature() {
                s.Render(e, t)
            }
        }
    }
}
