package api

type EqualHashRule interface {
	Equal(p interface{}) bool
	Hash() int
}
