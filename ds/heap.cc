


class TreeNode {
    public:
        TreeNode(int x): val(x) {}; 

    private:
        int val;
        TreeNode* left;
        TreeNode* right;
};

class MaxHeap {
    public:
        MaxHeap() {};

        void push(int val);

        int top();

        void pop();

    private:
        TreeNode* root;
};

void MaxHeap::push(int val) {
    if (!root) {
        root = new TreeNode(val);
        return;
    }
}
