package api

type Sort func(leftKey interface{}, rightKey interface{}) bool

//type SortGenerics [T any] func(leftKey T, rightKey T) bool
