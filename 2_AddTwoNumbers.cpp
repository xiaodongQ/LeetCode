/*
 * @Description: 
 * @Author: xd
 * @Date: 2019-10-29 08:37:08
 * @LastEditTime: 2019-11-03 22:23:47
 * @LastEditors: xd
 * @Note: 
 * source: https://leetcode-cn.com/problems/add-two-numbers
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
    如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
    您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

    示例：
    输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
    输出：7 -> 0 -> 8
    原因：342 + 465 = 807
 */
#include <iostream>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        //传入的事已经拆分好了的数字链表，遍历叠加即可，注意链表结束的边界条件
        //确保传入的是非空链表
        if (NULL == l1 || NULL == l2) {
            return NULL;
        }
        for 
    }
};