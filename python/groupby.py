from typing import List, Set


# this is an interesting problem I have been asked

items = [
    {
        'name': 'A',
        'status': 'active',
        'value': 1,
    },
    {
        'name': 'A',
        'status': 'inactive',
        'value': 2,
    },
    {
        'name': 'B',
        'status': 'active',
        'value': 3,
    },
]


class TreeNode:
    def __init__(self, val: str, children=None):
        self.val = val
        self.children = children or set()
        self.items = []

    def __eq__(self, other):
        return self.val == other.val

    def __hash__(self):
        return hash(self.val)

    def add_item(self, item):
        self.items.append(item)


def build_tree(root: TreeNode, items: List[dict], groups: List[str], i: int):
    if i == len(groups):
        return
    for item in items:
        group_name = groups[i]
        node = TreeNode(item.get(group_name))
        if i == len(groups) - 1:
            node.add_item(item)
        root.children.add(node)

    for child in root.children:
        build_tree(child, items, groups, i+1)


def traverse(root: TreeNode):
    if root is None:
        return
    print(root.val)
    if root.items:
        print([i for i in root.items])
    for child in root.children:
        traverse(child)


def build_dict(root: TreeNode):
    if not root:
        return {}
    d = {}
    for child in root.children:
        if child.items:
            d[child.val] = child.items
        else:
            d[child.val] = build_dict(child)
    return d


if __name__ == '__main__':
    root = TreeNode("0")
    build_tree(root, items, ['name', 'status'], 0)
    traverse(root)
    print(build_dict(root))
