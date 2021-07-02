package api

type Structure interface {
	Contains(val EqualHashRule) bool
	ToArray() []EqualHashRule
	Size() int
}
