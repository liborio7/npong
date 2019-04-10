package core

type Command interface {
    Execute(*Entity)
}
