package api

type Tree interface {
	StructureSize
	Get(key EqualHashRule) interface{}
	Put(key EqualHashRule, val interface{})
	Contains(key EqualHashRule) bool
	Remove(key EqualHashRule) (EqualHashRule, interface{})
	GetRoot() (key EqualHashRule, val interface{})
	ToArray() []EqualHashRule
}
