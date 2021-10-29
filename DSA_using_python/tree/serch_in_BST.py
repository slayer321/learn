
# A binary tree node
class Node:

    # Constructor to create a new node
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None


def solve(root, val):
    # if root == None:
    #     return "Not there"
    arr = []
    if root:
        if root.data == val:
            return get_all_node(root, arr)
        if root.data > val:
            return solve(root.left, val)
        elif root.data < val:
            return solve(root.right, val)
    return


def get_all_node(search_node, arr):
    if search_node:
        arr.append(search_node.data)
        get_all_node(search_node.left, arr)
        get_all_node(search_node.right, arr)

    return arr


# Driver program to test above function
root = Node(4)
root.left = Node(2)
root.right = Node(7)
root.left.left = Node(1)
root.left.right = Node(3)
#root.left.left.left = Node(15)


print((solve(root, 2)))
