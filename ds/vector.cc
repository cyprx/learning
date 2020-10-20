


class Vector {
    private:
        int arr[];
    public:
        Vector() = default;
        ~Vector() = default;
        void push_back(int val);
        void push_back(int val, int i);
        void pop();
};

