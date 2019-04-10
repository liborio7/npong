package component

import (
    "github.com/liborio7/npong/core"
)

const JoystickSignature = 1 << JOYSTICK

type Joystick struct {
    Conf map[string]core.Command
}

func (j *Joystick) Signature() uint64 {
    return JoystickSignature
}
