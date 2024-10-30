/*
 * @lc app=leetcode id=707 lang=cpp
 *
 * [707] Design Linked List
 *
 * https://leetcode.com/problems/design-linked-list/description/
 *
 * algorithms
 * Medium (28.54%)
 * Likes:    2706
 * Dislikes: 1633
 * Total Accepted:    363.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]\n' +
  '[[],[1],[3],[1,2],[1],[1],[1]]'
 *
 * Design your implementation of the linked list. You can choose to use a
 * singly or doubly linked list.
 * A node in a singly linked list should have two attributes: val and next. val
 * is the value of the current node, and next is a pointer/reference to the
 * next node.
 * If you want to use the doubly linked list, you will need one more attribute
 * prev to indicate the previous node in the linked list. Assume all nodes in
 * the linked list are 0-indexed.
 * 
 * Implement the MyLinkedList class:
 * 
 * 
 * MyLinkedList() Initializes the MyLinkedList object.
 * int get(int index) Get the value of the index^th node in the linked list. If
 * the index is invalid, return -1.
 * void addAtHead(int val) Add a node of value val before the first element of
 * the linked list. After the insertion, the new node will be the first node of
 * the linked list.
 * void addAtTail(int val) Append a node of value val as the last element of
 * the linked list.
 * void addAtIndex(int index, int val) Add a node of value val before the
 * index^th node in the linked list. If index equals the length of the linked
 * list, the node will be appended to the end of the linked list. If index is
 * greater than the length, the node will not be inserted.
 * void deleteAtIndex(int index) Delete the index^th node in the linked list,
 * if the index is valid.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input
 * ["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get",
 * "deleteAtIndex", "get"]
 * [[], [1], [3], [1, 2], [1], [1], [1]]
 * Output
 * [null, null, null, null, 2, null, 3]
 * 
 * Explanation
 * MyLinkedList myLinkedList = new MyLinkedList();
 * myLinkedList.addAtHead(1);
 * myLinkedList.addAtTail(3);
 * myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
 * myLinkedList.get(1);              // return 2
 * myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
 * myLinkedList.get(1);              // return 3
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= index, val <= 1000
 * Please do not use the built-in LinkedList library.
 * At most 2000 calls will be made to get, addAtHead, addAtTail, addAtIndex and
 * deleteAtIndex.
 * 
 * 
 */

// @lc code=start
class MyLinkedList {
public:
    // 定义链表结构
    struct LinkedNode {
        int val;
        LinkedNode *next;
        LinkedNode(int val):val(val), next(nullptr){}
    };

    MyLinkedList() {
        size = 0;
        dummyHead = new LinkedNode(0);
    }
    
    int get(int index) {
        if (index > size - 1) {
            return;
        }

        // 遍历到index前一个节点
        LinkedNode *cur = dummyHead;
        for (int i = 0; i < index - 1; i++) {
            cur = cur->next;
        }
        return cur->next->val;
    }
    
    void addAtHead(int val) {
        LinkedNode *node = new LinkedNode(val);
        node->next = dummyHead->next;
        dummyHead->next = node;
        size++;
    }
    
    void addAtTail(int val) {
        LinkedNode *cur = dummyHead;
        // 遍历到最后节点
        while (cur->next != nullptr) {
            cur = cur->next;
        }
        LinkedNode *node = new LinkedNode(val);
        cur->next = node;
        size++;
    }
    
    void addAtIndex(int index, int val) {
        if (index > size) {
            return;
        }
        // 遍历到index前一个节点
        LinkedNode *cur = dummyHead;
        for (int i = 0; i < index - 1; i++) {
            cur = cur->next;
        }
        LinkedNode *node = new LinkedNode(val);
        node->next = cur->next;
        cur->next = node;
        size++;
    }
    
    void deleteAtIndex(int index) {
        if (index > size) {
            return;
        }
        // 遍历到index前一个节点
        LinkedNode *cur = dummyHead;
        for (int i = 0; i < index - 1; i++) {
            cur = cur->next;
        }
        LinkedNode *node = cur->next;
        cur->next = cur->next->next;
        delete node;
        size--;
    }

private:
    // 链表长度，获取指定索引时需要校验
    int size;
    // 虚拟头节点，实际next为链表
    LinkedNode *dummyHead;
};

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * MyLinkedList* obj = new MyLinkedList();
 * int param_1 = obj->get(index);
 * obj->addAtHead(val);
 * obj->addAtTail(val);
 * obj->addAtIndex(index,val);
 * obj->deleteAtIndex(index);
 */
// @lc code=end

