class Node:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

    # def add_child(self, data):
    #     if data == self.data:
    #         return

    #     if data < self.data:
    #         if self.left:
    #             self.left.add_child(data)
    #         else:
    #             self.left = Node(data)
    #     else:
    #         if self.right:
    #             self.right.add_child(data)
    #         else:
    #             self.right = Node(data)


def insert_val(self, root, val):
    if root:
        # if root.val == val:
        #     return root.val
        if root.left == None and root.val > val:
            root.left = Node(val)
            return root
        elif root.right == None and root.val < val:
            root.right = Node(val)
            return root
        if root.val > val:
            return insert_val(root.left, val)
        elif root.val < val:
            return insert_val(root.right, val)
    return Node(val)


# def build_tree(elements):
#     if len(elements):
#         root = Node(elements[0])
#         for i in range(1, len(elements)):
#             root.add_child(elements[i])

#     return root

# def print_val(root, val):
#     tmp = root
#     insert_val(root, val)
#     return tmp.val


if __name__ == '__main__':
    # numbers = []
    # numbers_tree = build_tree(numbers)
    # print(numbers_tree.insert_val())
    root = Node(4)
    root.left = Node(2)
    root.right = Node(7)
    root.left.left = Node(1)
    root.left.right = Node(3)


#print((insert_val(root, 5)))
# print(print_val(root, 5))
