#include <iostream>
#include <list>
#include <unordered_map>
#include <memory>
#include <utility>


// suppport set, get value with max capacity, least recently used items will be remove when capacity reached

class LRUCache {
    public:
        LRUCache(int x): capacity(x) {};

        void put(int key, int val);

        int get(int key);

    private:
        typedef std::list<int> LI;
        typedef std::pair<int, LI::iterator> PI;
        typedef std::unordered_map<int, PI> MPI;

        int capacity;
        MPI kv;
        LI ll;
};


void LRUCache::put(int key, int val) {
    // check if key exists
    // if exists, update value and order of linked list
    if (kv.find(key) != kv.end()) {
        auto it = kv[key].second;
        ll.erase(it);
    }

    // check capacity
    // if exceeds, remove least used item
    if (ll.size() == capacity) {
        int k = ll.back();
        ll.pop_back();
        kv.erase(k);
    }
    ll.push_front(key);
    kv[key] = {val, ll.begin()};
}

int LRUCache::get(int key) {
    if (kv.find(key) == kv.end()) return -1;
    auto p = kv[key];
    auto it = p.second;
    ll.erase(it);
    ll.push_front(key);
    kv[key] = {p.first, ll.begin()};
    return p.first;
}


int main() {
    int capacity = 10;
    auto lru = new LRUCache(10);

    for (int i = 0; i < 20; i++) {
        lru->put(i, i);
    }

    for (int i = 0; i < 20; i++) {
        std::cout << lru->get(i) << "\t";
    }
    delete lru;

    return 0;
}
