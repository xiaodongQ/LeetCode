/*
 * @lc app=leetcode id=59 lang=cpp
 *
 * [59] Spiral Matrix II
 *
 * https://leetcode.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (72.29%)
 * Likes:    6489
 * Dislikes: 267
 * Total Accepted:    626.6K
 * Total Submissions: 866.7K
 * Testcase Example:  '3'
 *
 * Given a positive integer n, generate an n x n matrix filled with elements
 * from 1 to n^2 in spiral order.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: n = 3
 * Output: [[1,2,3],[8,9,4],[7,6,5]]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: n = 1
 * Output: [[1]]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= n <= 20
 * 
 * 
 */

// @lc code=start
class Solution {
public:
/*
vector<vector<int>> generateMatrix(int n) {
        // 核心是保持循环不变量。左闭右开区间，每圈依次按 上->右->下->左 的顺序遍历
        int startx = 0;
        int starty = 0;
        // 各框取值，从1开始
        int num = 1;
        // 轮次
        int round = n/2;
        // 每轮边界
        int offset = 1;
        int i=0, j=0;
        // 结果
        vector< vector<int> > result(n, vector<int>(n, 0));
        while (round-- > 0) {
            i = startx;
            j = starty;

            // 上边，此处依次获取一行记录 result[0][j]
            for (; j < n-offset; j++) {
                result[i][j] = num++;
            }
            // 右边，result[i][上轮的j列]
            for (; i < n-offset; i++) {
                result[i][j] = num++;
            }
            // 下边，注意是从右到左，result[上轮的i行][j]
            // 此处理解难点：行固定，列从最右到左边，到底是x坐标还是y坐标，换成下面的top和left更易理解
            for (; j > starty; j--) {
                result[i][j] = num++;
            }
            // 左边，注意是从下到上，result[i][上轮的j列]
            for (; i > startx; i--) {
                result[i][j] = num++;
            }
            startx++;
            starty++;
            offset++;
        }
        if (n % 2 == 1) {
            result[n/2][n/2] = num;
        }
        return result;
    }
};
*/

// 上述的每轮起始坐标(startx, starty) 换成 (top, left) 更易理解
vector<vector<int>> generateMatrix(int n) {
        // 核心是保持循环不变量。左闭右开区间，每圈依次按 上->右->下->左 的顺序遍历
        // 每轮开始的行
        int top = 0;
        // 每轮开始的列
        int left = 0;
        // 各位置取值，从1开始
        int num = 1;
        // 轮次
        int round = n/2;
        // 每轮到达的右边界偏移（最右n-offset）
        // 由于是开区间，右边界到倒数第2列（ 即`[0, n-1)` ）
        int offset = 1;
        int i=0, j=0;
        // 结果
        vector< vector<int> > result(n, vector<int>(n, 0));
        while (round-- > 0) {
            // 每轮的起点，行、列
            i = top;
            j = left;

            // 上边，此处依次获取一行记录 result[0][j]
            for (; j < n-offset; j++) {
                result[i][j] = num++;
            }
            // 右边，result[i][上轮的j列]
            for (; i < n-offset; i++) {
                result[i][j] = num++;
            }
            // 下边，注意是从右到左，result[上轮的i行][j]
            for (; j > left; j--) {
                result[i][j] = num++;
            }
            // 左边，注意是从下到上，result[i][上轮的j列]
            for (; i > top; i--) {
                result[i][j] = num++;
            }

            // 表示从顶往下一行 
            top++;
            // 表示从左往右一列
            left++;
            // n-offset 宽度变小
            offset++;
        }
        if (n % 2 == 1) {
            result[n/2][n/2] = num;
        }
        return result;
    }
};
// @lc code=end

