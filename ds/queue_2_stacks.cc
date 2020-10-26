#include <iostream>
#include <stack>


class Queue {
    public:
        Queue() {};

        // add to tail of queue
        void push_back(int val);

        // return front value
        int top();

        // check if queue is empty
        bool empty();

    private:
        std::stack<int> s1;
        std::stack<int> s2;
};


// O(n)
void Queue::push_back(int val) {
    while (!s1.empty()) {
        s2.push(s1.top());
        s1.pop();
    }
    s1.push(val);
    while (!s2.empty()) {
        s1.push(s2.top());
        s2.pop();
    }
}


// O(1)
int Queue::top() {
    int val = s1.top();
    s1.pop();
    return val;
}

bool Queue::empty() {
    return s1.empty();
}

int main() {
    Queue* q = new Queue();
    for (int i = 0; i < 10; i++) {
        q->push_back(i);
    }
    while (!q->empty()) {
        std::cout << q->top() << '\t';
    }
    delete q;
}


