package api

type Tree interface {
	Map
	GetRoot() (key EqualHashRule, val interface{})
}
