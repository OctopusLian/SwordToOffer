/*
 * @Description: 
 * @Author: neozhang
 * @Date: 2022-04-28 23:36:19
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-28 23:36:19
 */
class Solution {
    public boolean findNumberIn2DArray(int[][] matrix, int target) {
        //从数组的左下角位置开始去搜索整个二维数组
        //1，当发现当前遍历的元素大于target时，意味着这个元素后面的所有元素也都大于target，那么就不用去搜索这一行了
        //2，当发现当前遍历的元素小于target时，意味着这个元素上面的所有元素也都小于target，那么就不用去搜索这一列了
        
        //初始化i和j为数组左下角元素
        //最后一行
        int i = matrix.length - 1;
        
        //第0列
        int j = 0;
        
        //从数组的左下角开始出发，只要 i 和 j 没有越界继续判断
        while(i >= 0 && j < matrix[0].length) {
            //当发现当前遍历的元素大于target时，意味着这个元素后面的所有元素也都大于target
            if (matrix[i][j] > target) {
                //行索引向上移动，即i--，即消去矩阵第i行元素
                i--;
            //当发现当前遍历的元素小于target时，意味着这个元素的所有元素也都小于target
            } else if (matrix[i][j] < target) {
                //列索引向右移动一格，即j++，即消去矩阵第j列元素
                j++;
            //否则，说明找到目标值
            } else {
                return true;
            }
        }
        
        return false;
    }
}