/*
 * @lc app=leetcode id=19 lang=cpp
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* result;

        // 头节点可能被移除，定义一个虚拟头
        ListNode* dummyHead = new ListNode(0);
        dummyHead->next = head;
        // 双指针法
        ListNode* slow = dummyHead;
        ListNode* fast = dummyHead;
        
        // 题目中边界限定了 n 不会超过链表长度，可以不单独考虑n没结束但fast提前到nullptr
        // 先让快指针走n+1步，这里n+1步是核心，正好让slow在待删的前一个节点
        while (n-- > 0 && fast != nullptr) {
            fast = fast->next;
        }
        fast = fast->next;

        // 一起步进，fast到底则slow->next就是要移除的节点
        while (fast != nullptr) {
            slow = slow->next;
            fast = fast->next;
        }
        ListNode* tmp = slow->next;
        slow->next = slow->next->next;
        delete tmp;
        result = dummyHead->next;
        delete dummyHead;

        return result;
    }
};
// @lc code=end

