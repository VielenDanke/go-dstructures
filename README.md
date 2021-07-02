# Go data structures

## Important interfaces

The key interface to work with - EqualHashRule.

It needs for correct equality in unique structures,
also hash needed in HashSet, HashMap, TreeSet, TreeMap structures. 

Example of using EqualHashRule interfaces:
````
type cStr string

func (c cStr) Equal(p interface{}) bool {
	incStr := p.(cStr)
	return string(c) == string(incStr)
}

func (c cStr) Hash() int {
	i := 0
	for _, v := range c {
		i += int(v)
	}
	return i
}
````

## Data structures