package api

type Set interface {
	Structure
	Add(val EqualHashRule)
	Remove(val EqualHashRule) EqualHashRule
}
