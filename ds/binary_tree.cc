#include <iostream>

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;

    TreeNode(): val(0), left(nullptr), right(nullptr) {};
    TreeNode(int val, TreeNode* left, TreeNode* right): val(val), left(left), right(right) {};
};

int main() {
    TreeNode* root = new TreeNode();
    std::cout << root->val;
}
