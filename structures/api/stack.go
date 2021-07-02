package api

type Stack interface {
	Structure
	Push(val EqualHashRule)
	Pop() (EqualHashRule, bool)
}
