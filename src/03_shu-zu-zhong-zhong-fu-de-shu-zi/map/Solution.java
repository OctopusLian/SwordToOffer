import java.util.Set;

/*
 * @Description: 
 * @Author: neozhang
 * @Date: 2022-04-28 22:55:42
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-28 23:00:17
 */
class Solution {
    public int findRepeatNumber(int[] nums) {
        //HashSet的特点是不会存储重复元素
        //所以可以用HashSet来查找出重复的元素
        Set<Integer> dic = new HashSet<>();
        //遍历数组，设置此时遍历的元素为num
        for (int num : nums) {
            //如果发现dic中已经存储了num
            //那么说明找到了重复的那个元素
            if (dic.contains(num)) {
                //把num这个结果进行返回
                return num;
            } else {
                //否则，说明dic中还没有存储num
                dic.add(num);
            }
        }
        return -1;
    }
}