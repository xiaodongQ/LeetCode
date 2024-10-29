/*
 * @lc app=leetcode id=203 lang=cpp
 *
 * [203] Remove Linked List Elements
 *
 * https://leetcode.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (50.15%)
 * Likes:    8422
 * Dislikes: 254
 * Total Accepted:    1.2M
 * Total Submissions: 2.5M
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * Given the head of a linked list and an integer val, remove all the nodes of
 * the linked list that has Node.val == val, and return the new head.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: head = [1,2,6,3,4,5,6], val = 6
 * Output: [1,2,3,4,5]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: head = [], val = 1
 * Output: []
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: head = [7,7,7,7], val = 7
 * Output: []
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in the list is in the range [0, 10^4].
 * 1 <= Node.val <= 50
 * 0 <= val <= 50
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
    ListNode* removeElements(ListNode* head, int val) {
        // 定义一个虚拟节点，指向链表头
        ListNode *dummyHead = new ListNode();
        dummyHead->next = head;
        ListNode *currNode = dummyHead;
        // 注意循环中判断next，而不是用 currNode != nullptr，并对currNode迭代
        while (currNode->next != nullptr) {
            if (currNode->next->val == val) {
                // 下一个节点和val相等，则移除下个节点
                ListNode *tmp_node = currNode->next;
                currNode->next = currNode->next->next;
                delete tmp_node;
            } else {
                // 没有匹配到相等节点时迭代下一个，上面匹配到时next已经指向下一个了
                currNode = currNode->next;
            }
        }

        // dummyHead 是临时申请的空间，需要释放掉
        ListNode *node = dummyHead->next;
        delete dummyHead;

        // 返回的节点指针，是原来链表中已有的节点
        return node;
    }
};
// @lc code=end

