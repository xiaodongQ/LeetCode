/*
 * @Description: 
 * @Author: xd
 * @Date: 2019-10-29 08:37:08
 * @LastEditTime: 2019-11-07 22:26:50
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
    // ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    //     int inc = 0;
    //     ListNode* head = NULL;
    //     ListNode* tail = NULL;

    //     while (l1 || l2) {
    //         int v1 = l1 ? l1->val : 0;
    //         int v2 = l2 ? l2->val : 0;
    //         int sum = v1+v2+inc;
    //         inc = sum/10;
    //         if (!head) {
    //             head = tail = new ListNode(sum%10);
    //         }else{
    //             tail->next = new ListNode(sum%10);
    //             tail = tail->next;
    //         }
    //         if (l1) {
    //             l1 = l1->next;
    //         }
    //         if (l2) {
    //             l2 = l2->next;
    //         }
    //     }
    //     if (inc == 1) {
    //         tail->next = new ListNode(inc);
    //     }
    //     return head;
    // }
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int inc = 0;
        ListNode* head = NULL;
        ListNode* tail = NULL;

        while (l1 || l2) {
            int sum = 0;
            sum += inc;
            if (l1) {
                sum += l1->val;
                l1 = l1->next;
            }
            if (l2) {
                sum += l2->val;
                l2 = l2->next;
            }
            inc = sum/10;
            if (!head){
                head = tail = new ListNode(sum%10);
            } else {
                tail->next = new ListNode(sum%10);
                tail = tail->next;
            }
        }
        if (inc == 1) {
            tail->next = new ListNode(inc);
        }
        return head;
    }
};