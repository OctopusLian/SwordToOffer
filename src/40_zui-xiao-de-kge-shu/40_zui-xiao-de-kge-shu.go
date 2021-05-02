import "sort"

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func getLeastNumbers(arr []int, k int) []int {
	sort.Sort(IntSlice(arr))
	return arr[:k]
}