

# A binary tree node
class Node:

    # Constructor to create a new node
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None


def solve(node, arr, p):
    if not node:
        # out.append(arr)
        return False
    var = f"{node.data}->"
    arr.append(var)
    if (node.data == p):
        return True

    if (solve(node.left, arr, p) or solve(node.right, arr, p)):
        return True

    arr.pop(-1)
    return False


def pp(arr):
    pass


def printPath(root, p):

    # vector to store the path
    arr = []
    if root == None:
        return arr
    solve(root, arr, p)
    return arr

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


print(printPath(root, 9))
