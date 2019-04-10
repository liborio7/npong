package core

import (
    "time"
)

type System interface {
    Signature() uint64
    Init(*Entity)
    Update(*Entity, time.Time)
    Render(*Entity, time.Time)
}
