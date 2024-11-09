/*
 * @lc app=leetcode.cn id=142 lang=cpp
 *
 * [142] 环形链表 II
 */

// @lc code=start
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
    ListNode *detectCycle(ListNode *head) {
        // 快慢指针法
        ListNode *slow = head;
        ListNode *fast= head;
        while (fast != nullptr && fast->next != nullptr) {
            slow = slow->next;
            fast = fast->next->next;
            // 快慢节点相遇，但不是环形入口
            if (slow == fast) {
                // 直接用结论：从 链表起点 和 相遇点 各自按一步移动，相遇点即环形入口
                ListNode *tmp1 = head;
                ListNode *tmp2= slow;
                while (tmp1 != tmp2) {
                    tmp1 = tmp1->next;
                    tmp2 = tmp2->next;
                }
                return tmp1;
            }
        }
        return nullptr;
    }
};
// @lc code=end

