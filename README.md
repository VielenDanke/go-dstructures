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

Structure interface contains the common methods for structures
````
type Structure interface {
	Contains(val EqualHashRule) bool
	ToArray() []EqualHashRule
	Size() int
}
````
All next data structures contains those methods. Keep in mind

## Data structures

### Array Queue

Simple representation of FIFO based on array

#### BigO
Contains - O(N)
Enqueue - O(1)
Dequeue - O(1)

#### Creation and methods
Creation Queue:
````
q := structures.NewArrayQueue()
````

Queue methods:
````
Enqueue(val EqualHashRule)
Dequeue() (EqualHashRule, bool)
````
If queue is empty - Dequeue() return nil and false

---

### Array Stack

Simple representation of LIFO based on array

#### BigO
Contains - O(N)  
Push - O(1)  
Pop - O(1)

#### Creation and methods

Creation stack
````
s := structures.NewArrayStack()
````

Stack methods
````
Push(val EqualHashRule)
Pop() (EqualHashRule, bool)
````
If stack is empty - Pop() return nil and false

---

### Binary heap

Special structure to work with int

The main idea described here https://en.wikipedia.org/wiki/Binary_heap

#### BigO

Insertion - O(logN)
Removal - O(logN)
Search - O(N)

#### Creation and methods

Creation binary heap
````
bh := structures.NewBinaryHeap()
````

Binary heap methods
````
Insert(val int)
RemoveMax() (int, bool)
Contains(val int) bool
````
If heap is empty - RemoveMax() return -1, false

---