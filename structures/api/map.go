package api

type Map interface {
	Structure
	Get(key EqualHashRule) interface{}
	Put(key EqualHashRule, val interface{})
	Contains(key EqualHashRule) bool
	Remove(key EqualHashRule) (EqualHashRule, interface{})
	KeySet() []EqualHashRule
}
