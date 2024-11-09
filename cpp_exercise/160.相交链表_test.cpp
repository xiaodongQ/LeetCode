/*
 * @lc app=leetcode.cn id=160 lang=cpp
 *
 * [160] 相交链表
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
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        int lenA = 0;
        int lenB = 0;
        ListNode* tmpA = headA;
        ListNode* tmpB = headB;

        // 计算链表长度
        while (tmpA != nullptr) {
            tmpA = tmpA->next;
            lenA++;
        }
        while (tmpB != nullptr) {
            tmpB = tmpB->next;
            lenB++;
        }
        tmpA = headA;
        tmpB = headB;

        // 长链表先移动到两者长度相等的节点
        int sub = lenA >= lenB ? (lenA-lenB) : (lenB-lenA);
        if (lenA >= lenB) {
            while (sub--) {
                tmpA = tmpA->next;
            }
        } else {
            while (sub--) {
                tmpB = tmpB->next;
            }
        }

        // 比较是否相交
        while (tmpA != nullptr && tmpB != nullptr) {
            if (tmpA == tmpB) {
                // 相交
                return tmpA;
            }
            tmpA = tmpA->next;
            tmpB = tmpB->next;
        }

        return nullptr;
    }
};
// @lc code=end

