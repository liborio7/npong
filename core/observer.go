package core

import (
    "github.com/satori/go.uuid"
)

type Observer interface {
    Id() uuid.UUID
    Signature() uint64
    OnEvent(*Entity)
}
