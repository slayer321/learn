# Python3 program to find the maximum depth of tree

# A binary tree node
class Node:

    # Constructor to create a new node
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None


def solve(node, p, q):
    if (p.data > node.data) and (q.data > node.data):
        return solve(node.right, p, q)
    elif (p.data < node.data) and (q.data < node.data):
        return solve(node.left, p, q)
    else:
        return node.data


    # Driver program to test above function
root = Node(6)
root.left = Node(2)
root.right = Node(8)
root.left.left = Node(0)
root.left.right = Node(4)
root.right.left = Node(7)
root.right.left = Node(9)
root.left.right.left = Node(3)
root.left.right.right = Node(5)


#root.left.left.left = Node(15)


print((solve(root, root.left.left, root.right.left)))
