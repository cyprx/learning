#include <iostream>
#include <string>


struct ListNode {
    int val;
    ListNode* next;

    ListNode(): val(0), next(nullptr) {};
    ListNode(int val): val(val), next(nullptr) {};
    ListNode(int val, ListNode* next): val(val), next(next) {};
};


int main() {
    ListNode* root = new ListNode(0);
    ListNode* cur = root;
    for (int i = 1; i < 5; i++) {
        cur->next = new ListNode(i);
        cur = cur->next;
    }

    cur = root;
    while (cur != nullptr) {
        std::cout << cur->val << " ";
        cur = cur->next;
    }
}
