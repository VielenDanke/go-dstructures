package api

type LinkedList interface {
	Structure
	Push(val EqualHashRule)
	Pop() (EqualHashRule, bool)
	Shift() (EqualHashRule, bool)
	Unshift(val EqualHashRule)
	Get(idx int) (EqualHashRule, bool)
	Set(idx int, val EqualHashRule) bool
	Insert(idx int, val EqualHashRule) bool
	RemoveIdx(idx int) (EqualHashRule, bool)
	Remove(val EqualHashRule) (EqualHashRule, bool)
	Reverse()
}
