package main

func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, i int, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	m := p
	for postorder[p] > postorder[j] {
		p++
	}

	return p == j && recur(postorder, i, m-1) && recur(postorder, m, j-1)
}
