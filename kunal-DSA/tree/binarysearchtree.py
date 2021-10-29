class BinarySearchTree:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None

    def add_child(self, val):
        if val == self.data:
            return

        if val < self.data:
            if self.left:
                self.left.add_child(val)
            else:
                self.left = BinarySearchTree(val)
        else:
            if self.right:
                self.right.add_child(val)
            else:
                self.right = BinarySearchTree(val)

    def in_order(self):
        elements = []
        if self.left:
            elements += self.left.in_order()

        elements.append(self.data)

        if self.right:
            elements += self.right.in_order()

        return elements

    def pre_order(self):
        elements = []

        elements.append(self.data)
        if self.left:
            elements += self.left.pre_order()

        if self.right:
            elements += self.right.pre_order()

        return elements

    def height(self, root):
        if root == None:
            return 0
        left = self.height(self.left)
        right = self.height(self.right)

        out = max(left, right) + 1
        return out


def build_tree(element):
    root = BinarySearchTree(element[0])
    for i in range(1, len(element)):
        root.add_child(element[i])

    return root


if __name__ == '__main__':
    arr = [17, 4, 1, 20, 9, 23, 18, 34]
    val = build_tree(arr)
    print(val.in_order())
    print(val.pre_order())
    print(val.height(val))
