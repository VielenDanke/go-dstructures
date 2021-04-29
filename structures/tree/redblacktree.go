package tree

import (
	"fmt"
	"github.com/vielendanke/go-dstructures/structures/api"
	"github.com/vielendanke/go-dstructures/structures/queue"
)

const (
	RED = iota
	BLACK
)

type rbTreeNode struct {
	key    api.EqualHashRule
	value  interface{}
	parent *rbTreeNode
	left   *rbTreeNode
	right  *rbTreeNode
	color  int
}

func (rbt *rbTreeNode) Equal(p interface{}) bool {
	incTreeNode := p.(*rbTreeNode)
	return rbt.key.Equal(incTreeNode.key)
}

func (rbt *rbTreeNode) Hash() int {
	return rbt.key.Hash()
}

func (rbt *rbTreeNode) String() string {
	return fmt.Sprintf("Node: {%v:%v}", rbt.key, rbt.value)
}

type rbTree struct {
	root     *rbTreeNode
	sortFunc api.Sort
	length   int
}

func NewRBTree(sortFunc api.Sort) api.Tree {
	return &rbTree{sortFunc: sortFunc}
}

func (r *rbTree) GetRoot() (key api.EqualHashRule, val interface{}) {
	if r.root == nil {
		return nil, nil
	}
	return r.root.key, r.root.value
}

func (r *rbTree) Remove(key api.EqualHashRule) (api.EqualHashRule, interface{}) {
	if r.root == nil || key == nil {
		return nil, nil
	}
	var temp *rbTreeNode
	if r.Size() == 1 {
		if r.root.key.Equal(key) {
			r.length--
			temp = r.root
			r.root = nil
			return temp.key, temp.value
		} else {
			return nil, nil
		}
	}
	temp = r.findRBNode(r.root, key)
	if temp != nil {
		r.length--
		rNode := temp
		r.deleteRBNode(temp)
		return rNode.key, rNode.value
	}
	return nil, nil
}

func (r *rbTree) Put(key api.EqualHashRule, val interface{}) {
	if key == nil {
		return
	}
	if r.root == nil {
		r.root = &rbTreeNode{key: key, value: val}
	} else {
		toInsert := &rbTreeNode{key: key, value: val}
		r.insertNewRBNode(r.root, toInsert)
	}
	r.length++
}

func (r *rbTree) Contains(key api.EqualHashRule) bool {
	if r.root == nil || key == nil {
		return false
	}
	if r.Size() == 1 {
		return r.root.key.Equal(key)
	} else {
		nFind := r.findRBNode(r.root, key)
		if nFind != nil {
			return true
		}
		return false
	}
}

func (r *rbTree) Get(key api.EqualHashRule) interface{} {
	if r.root == nil || key == nil {
		return nil
	}
	if r.Size() == 1 {
		if r.root.key.Equal(key) {
			return r.root.value
		} else {
			return nil
		}
	} else {
		return r.findRBNode(r.root, key)
	}
}

func (r *rbTree) ToArray() []api.EqualHashRule {
	return r.breadthForSearchKeys()
}

func (r *rbTree) Size() int {
	return r.length
}

func (r *rbTree) String() string {
	return fmt.Sprintf("Nodes: {%v}", r.breadthForSearchNodes())
}

func (r *rbTree) findRBNode(n *rbTreeNode, key api.EqualHashRule) (res *rbTreeNode) {
	if n == nil {
		return
	}
	if n.key.Equal(key) {
		return n
	}
	if r.sortFunc(n.key, key) {
		res = r.findRBNode(n.left, key)
	} else {
		res = r.findRBNode(n.right, key)
	}
	return
}

func (r *rbTree) insertNewRBNode(n *rbTreeNode, toInsert *rbTreeNode) {
	if n.key.Equal(toInsert.key) {
		if toInsert.value != nil {
			n.value = toInsert.value
		}
		return
	}
	toInsert.parent = n
	if r.sortFunc(n.key, toInsert.key) {
		if n.left == nil {
			n.left = toInsert
		} else {
			r.insertNewRBNode(n.left, toInsert)
		}
	} else {
		if n.right == nil {
			n.right = toInsert
		} else {
			r.insertNewRBNode(n.right, toInsert)
		}
	}
	r.fixAfterInsertion(toInsert)
}

func (r *rbTree) deleteRBNode(p *rbTreeNode) {
	if p.left != nil && p.right != nil {
		sNode := successor(p)
		p.key = sNode.key
		p.value = sNode.value
		p = sNode
	}
	var replacement *rbTreeNode

	if p.left != nil {
		replacement = p.left
	} else {
		replacement = p.right
	}
	if replacement != nil {
		replacement.parent = p.parent

		if p.parent == nil {
			r.root = replacement
		} else if nodeEqualRB(p, p.parent.left) {
			p.parent.left = replacement
		} else {
			p.parent.right = replacement
		}
		p.left = nil
		p.right = nil
		p.parent = nil

		if p.color == BLACK {
			r.fixAfterDeletion(replacement)
		}
	} else if p.parent == nil {
		r.root = nil
	} else {
		if p.color == BLACK {
			r.fixAfterDeletion(p)
		}
		if p.parent != nil {
			if nodeEqualRB(p, p.parent.left) {
				p.parent.left = nil
			} else if nodeEqualRB(p, p.parent.right) {
				p.parent.right = nil
			}
			p.parent = nil
		}
	}
}

func successor(node *rbTreeNode) *rbTreeNode {
	if node == nil {
		return nil
	} else if node.right != nil {
		pNode := node.right
		for pNode.left != nil {
			pNode = pNode.left
		}
		return pNode
	} else {
		pNode := node.parent
		chNode := node
		for pNode != nil && nodeEqualRB(chNode, pNode.right){
			chNode = pNode
			pNode = pNode.parent
		}
		return pNode
	}
}

func (r *rbTree) fixAfterInsertion(xNode *rbTreeNode) {
	for xNode != nil && !nodeEqualRB(xNode, r.root) && xNode.parent.color == RED {
		if nodeEqualRB(parentNodeOf(xNode), leftNodeOf(parentNodeOf(parentNodeOf(xNode)))) {
			yNode := rightNodeOf(parentNodeOf(parentNodeOf(xNode)))

			if colorNodeOf(yNode) == RED {
				setNodeColor(parentNodeOf(xNode), BLACK)
				setNodeColor(yNode, BLACK)
				setNodeColor(parentNodeOf(parentNodeOf(xNode)), RED)
				xNode = parentNodeOf(parentNodeOf(xNode))
			} else {
				if nodeEqualRB(xNode, rightNodeOf(parentNodeOf(xNode))) {
					xNode = parentNodeOf(xNode)
					r.rotateNodeLeft(xNode)
				}
				setNodeColor(parentNodeOf(xNode), BLACK)
				setNodeColor(parentNodeOf(parentNodeOf(xNode)), RED)
				r.rotateNodeRight(parentNodeOf(parentNodeOf(xNode)))
			}
		} else {
			yNode := leftNodeOf(parentNodeOf(parentNodeOf(xNode)))

			if colorNodeOf(yNode) == RED {
				setNodeColor(parentNodeOf(xNode), BLACK)
				setNodeColor(yNode, BLACK)
				setNodeColor(parentNodeOf(parentNodeOf(xNode)), RED)
				xNode = parentNodeOf(parentNodeOf(xNode))
			} else {
				if nodeEqualRB(xNode, leftNodeOf(parentNodeOf(xNode))) {
					xNode = parentNodeOf(xNode)
					r.rotateNodeRight(xNode)
				}
				setNodeColor(parentNodeOf(xNode), BLACK)
				setNodeColor(parentNodeOf(parentNodeOf(xNode)), RED)
				r.rotateNodeLeft(parentNodeOf(parentNodeOf(xNode)))
			}
		}
		r.root.color = BLACK
	}
}

func (r *rbTree) fixAfterDeletion(x *rbTreeNode) {
	for !nodeEqualRB(x, r.root) && colorNodeOf(x) == BLACK {
		if nodeEqualRB(x, leftNodeOf(parentNodeOf(x))) {
			sib := rightNodeOf(parentNodeOf(x))

			if colorNodeOf(sib) == RED {
				setNodeColor(sib, BLACK)
				setNodeColor(parentNodeOf(x), RED)
				r.rotateNodeLeft(parentNodeOf(x))
				sib = rightNodeOf(parentNodeOf(x))
			}
			if colorNodeOf(leftNodeOf(sib)) == BLACK && colorNodeOf(rightNodeOf(sib)) == BLACK {
				setNodeColor(sib, RED)
				x = parentNodeOf(x)
			} else {
				if colorNodeOf(rightNodeOf(sib)) == BLACK {
					setNodeColor(leftNodeOf(sib), BLACK)
					setNodeColor(sib, RED)
					r.rotateNodeRight(sib)
					sib = rightNodeOf(parentNodeOf(x))
				}
				setNodeColor(sib, colorNodeOf(parentNodeOf(x)))
				setNodeColor(parentNodeOf(x), BLACK)
				setNodeColor(rightNodeOf(sib), BLACK)
				r.rotateNodeLeft(parentNodeOf(x))
				x = r.root
			}
		} else {
			sib := leftNodeOf(parentNodeOf(x))

			if colorNodeOf(sib) == RED {
				setNodeColor(sib, BLACK)
				setNodeColor(parentNodeOf(x), RED)
				r.rotateNodeRight(parentNodeOf(x))
				sib = leftNodeOf(parentNodeOf(x))
			}
			if colorNodeOf(rightNodeOf(sib)) == BLACK && colorNodeOf(leftNodeOf(sib)) == BLACK {
				setNodeColor(sib, RED)
				x = parentNodeOf(x)
			} else {
				if colorNodeOf(leftNodeOf(sib)) == BLACK {
					setNodeColor(rightNodeOf(sib), BLACK)
					setNodeColor(sib, RED)
					r.rotateNodeLeft(sib)
					sib = leftNodeOf(parentNodeOf(x))
				}
				setNodeColor(sib, colorNodeOf(parentNodeOf(x)))
				setNodeColor(parentNodeOf(x), BLACK)
				setNodeColor(leftNodeOf(sib), BLACK)
				r.rotateNodeRight(parentNodeOf(x))
				x = r.root
			}
		}
	}
	setNodeColor(x, BLACK)
}

func (r *rbTree) rotateNodeRight(rotateNode *rbTreeNode) {
	if rotateNode != nil {
		lNode := rotateNode.left
		rotateNode.left = lNode.right

		if lNode.right != nil {
			lNode.right.parent = rotateNode
		}
		lNode.parent = rotateNode.parent

		if rotateNode.parent == nil {
			r.root = lNode
		} else if nodeEqualRB(rotateNode.parent.right, rotateNode) {
			rotateNode.parent.right = lNode
		} else {
			rotateNode.parent.left = lNode
		}
		lNode.right = rotateNode
		rotateNode.parent = lNode
	}
}

func (r *rbTree) rotateNodeLeft(rotateNode *rbTreeNode) {
	if rotateNode != nil {
		rNode := rotateNode.right
		rotateNode.right = rNode.left

		if rNode.left != nil {
			rNode.left.parent = rotateNode
		}
		rNode.parent = rotateNode.parent

		if rotateNode.parent == nil {
			r.root = rNode
		} else if nodeEqualRB(rotateNode.parent.left, rotateNode) {
			rotateNode.parent.left = rNode
		} else {
			rotateNode.parent.right = rNode
		}
		rNode.left = rotateNode
		rotateNode.parent = rNode
	}
}

func colorNodeOf(node *rbTreeNode) int {
	if node != nil {
		return node.color
	}
	return BLACK
}

func setNodeColor(node *rbTreeNode, color int) {
	if node != nil {
		node.color = color
	}
}

func nodeEqualRB(fNode, sNode *rbTreeNode) bool {
	if fNode != nil && sNode != nil {
		return fNode.key.Equal(sNode.key)
	}
	if fNode == nil && sNode == nil {
		return true
	}
	return false
}

func leftNodeOf(node *rbTreeNode) *rbTreeNode {
	if node != nil {
		return node.left
	}
	return nil
}

func rightNodeOf(node *rbTreeNode) *rbTreeNode {
	if node != nil {
		return node.right
	}
	return nil
}

func parentNodeOf(node *rbTreeNode) *rbTreeNode {
	if node != nil {
		return node.parent
	}
	return nil
}

func (r *rbTree) breadthForSearchKeys() []api.EqualHashRule {
	res := make([]api.EqualHashRule, 0)
	if r.root == nil {
		return nil
	}
	q := queue.NewArrayQueue()
	q.Enqueue(r.root)

	for q.Size() != 0 {
		elem, _ := q.Dequeue()
		n := elem.(*rbTreeNode)
		res = append(res, n.key)
		if n.left != nil {
			q.Enqueue(n.left)
		}
		if n.right != nil {
			q.Enqueue(n.right)
		}
	}
	return res
}

func (r *rbTree) breadthForSearchNodes() (res []*rbTreeNode) {
	if r.root == nil {
		return
	}
	q := queue.NewArrayQueue()
	q.Enqueue(r.root)

	for q.Size() != 0 {
		elem, _ := q.Dequeue()
		n := elem.(*rbTreeNode)
		res = append(res, n)
		if n.left != nil {
			q.Enqueue(n.left)
		}
		if n.right != nil {
			q.Enqueue(n.right)
		}
	}
	return
}
