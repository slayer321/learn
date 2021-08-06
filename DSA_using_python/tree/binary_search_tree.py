class BinarySearchTree:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None

    def add_child(self, data):
        if data == self.data:
            return

        if data < self.data:
            if self.left:
                self.left.add_child(data)
            else:
                self.left = BinarySearchTree(data)
        else:
            if self.right:
                self.right.add_child(data)
            else:
                self.right = BinarySearchTree(data)

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

    def post_order(self):
        elements = []

        if self.left:
            elements += self.left.post_order()

        if self.right:
            elements += self.right.post_order()

        elements.append(self.data)

        return elements


def build_tree(elements):
    root = BinarySearchTree(elements[0])
    for i in range(1, len(elements)):
        root.add_child(elements[i])

    return root


if __name__ == "__main__":
    numbers = [17, 4, 1, 20, 9, 23, 18, 34]
    numbers_tree = build_tree(numbers)
    print("Inorder: ", numbers_tree.in_order())
    print("Pre_order: ", numbers_tree.pre_order())
    print("Post_order: ", numbers_tree.post_order())
