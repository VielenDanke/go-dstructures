package api

type Tree interface {
	StructureSize
	Structure
	Get(key EqualHashRule) interface{}
	Put(key EqualHashRule, val interface{})
	Remove(key EqualHashRule) (EqualHashRule, interface{})
	GetRoot() (key EqualHashRule, val interface{})
}
