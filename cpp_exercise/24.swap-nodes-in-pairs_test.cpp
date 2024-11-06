/*
 * @lc app=leetcode id=24 lang=cpp
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (65.53%)
 * Likes:    12141
 * Dislikes: 458
 * Total Accepted:    1.5M
 * Total Submissions: 2.3M
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head. You
 * must solve the problem without modifying the values in the list's nodes
 * (i.e., only nodes themselves may be changed.)
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: head = [1,2,3,4]
 * 
 * Output: [2,1,4,3]
 * 
 * Explanation:
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: head = []
 * 
 * Output: []
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: head = [1]
 * 
 * Output: [1]
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: head = [1,2,3]
 * 
 * Output: [2,1,3]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in the list is in the range [0, 100].
 * 0 <= Node.val <= 100
 * 
 * 
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
    ListNode* swapPairs(ListNode* head) {
        // 由于头节点也涉及交换，还是定义一个虚拟头便于处理
        ListNode* dummyHead = new ListNode(0);
        dummyHead->next = head;
        ListNode* cur = dummyHead;
        // cur初始值和退出条件是关键
        while (cur->next != nullptr && cur->next->next != nullptr) {
            // 记录下用于步进
            // ListNode* tmp = cur->next;
            // 错误做法：交换两节点 dummyHead -> a -> b ->，先让a->next指向b的下一个节点，再b->next指向a
            // cur->next->next = cur->next->next->next;
            // tmp->next->next = tmp;

            // 上述光交换两节点不够
            ListNode* tmp = cur->next;
            // 再定义一个临时节点，记录 dummyHead -> a -> b -> 中的下一个节点
            ListNode* tmp1 = cur->next->next->next;
            // 交换两节点，并修改其next指向
            cur->next = cur->next->next;
            cur->next->next = tmp;
            cur->next->next->next = tmp1;

            // 步进cur
            cur = cur->next->next;
        }
        cur = dummyHead->next;
        delete dummyHead;

        return cur;
    }
};
// @lc code=end

