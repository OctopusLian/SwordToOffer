// 递归
// 终止条件：p为空时，s为空返回true，s不空返回false
// first标志首位是否相等，p[0]是.也表示相等
// (1)如果p的第2位为*
//    ①首位相同 则等同于判断s[1:]和p （忽略s[0]）
//    ②首位不同 则等同于判断s和p[2:] （将p前两位x*和空值配对）
// (2)否则不要期待奇迹出现了，
//    首位如果相等就各向后移一位，判断s[1:]和p[1:]
//    如果不等就是false
// 这里利用了go字符串的截断符号
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	var first bool
	if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		first = true
	}
	if len(p) > 1 && p[1] == '*' {
		return (first && isMatch(s[1:], p)) || isMatch(s, p[2:])
	}
	return first && isMatch(s[1:], p[1:])
}

//链接：https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/solution/diao-ku-di-gui-dong-tai-gui-hua-by-cmatrix/
