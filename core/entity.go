package core

import (
    "github.com/satori/go.uuid"
    "sync"
)

type Entity struct {
    Id              uuid.UUID
    Tags            []string
    tagsMutex       sync.RWMutex
    Signature       uint64
    Components      map[uint64]Component
    componentsMutex sync.RWMutex
    World           *World
}

func NewEntity() *Entity {
    return &Entity{
        Id:              uuid.NewV4(),
        Tags:            make([]string, 0),
        tagsMutex:       sync.RWMutex{},
        Signature:       uint64(0),
        Components:      make(map[uint64]Component, 0),
        componentsMutex: sync.RWMutex{},
        World:           nil,
    }
}

func (e *Entity) AddTags(tags ...string) {
    e.tagsMutex.Lock()
    e.Tags = append(e.Tags, tags...)
    e.tagsMutex.Unlock()
}

func (e *Entity) HasTag(tag string) bool {
    e.tagsMutex.RLock()
    res := false
    for _, t := range e.Tags {
        if t == tag {
            res = true
            break
        }
    }
    e.tagsMutex.RUnlock()
    return res
}

func (e *Entity) AddComponents(cs ...Component) {
    e.componentsMutex.Lock()
    for _, c := range cs {
        e.Components[c.Signature()] = c
        e.Signature |= c.Signature()
    }
    e.componentsMutex.Unlock()
}

func (e *Entity) Unpack(c Component) Component {
    e.componentsMutex.RLock()
    component := e.Components[c.Signature()]
    e.componentsMutex.RUnlock()
    return component
}
