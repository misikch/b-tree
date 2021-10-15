package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type TreeNode struct {
	val   int
	meta  map[string]interface{}
	right *TreeNode
	left  *TreeNode
}

func (t *TreeNode) Insert(value int, meta map[string]interface{}) error {
	if t == nil {
		return fmt.Errorf("tree is nil")
	}

	if t.val == value {
		return fmt.Errorf("node with value %d already exists", value)
	}

	if t.val > value {
		if t.left == nil {
			t.left = &TreeNode{val: value, meta: meta}
			return nil
		}

		return t.left.Insert(value, meta)
	}

	if t.val < value {
		if t.right == nil {
			t.right = &TreeNode{val: value, meta: meta}
			return nil
		}

		return t.right.Insert(value, meta)
	}

	return nil
}

func (t *TreeNode) FindMin() (int, map[string]interface{}) {
	if t == nil || t.left == nil {
		return t.val, t.meta
	}

	return t.left.FindMin()
}

func (t *TreeNode) FindMax() (int, map[string]interface{}) {
	if t == nil || t.right == nil {
		return t.val, t.meta
	}

	return t.right.FindMax()
}

func (t *TreeNode) PrintInOrder() {
	if t == nil {
		return
	}
	t.left.PrintInOrder()
	fmt.Printf("%d ", t.val)
	t.right.PrintInOrder()
}

func main() {
	tree := &TreeNode{}

	rand.Seed(time.Now().UnixNano())
	var set sync.Map
	var i int

	for {
		if i >= 10 {
			break
		}

		v := rand.Intn(600) - 200
		if _, exists := set.LoadOrStore(v, struct{}{}); exists {
			continue
		}

		meta := make(map[string]interface{}, 0)
		err := tree.Insert(v, meta)
		if err != nil {
			fmt.Println(err)
		}
		i++
	}

	tree.PrintInOrder()

	min, _ := tree.FindMin()
	max, _ := tree.FindMax()
	fmt.Printf("\nMin: %d\nMax: %d\n", min, max)
}
