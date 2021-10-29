# Python3 program to find the maximum depth of tree

# A binary tree node
class Node:

    # Constructor to create a new node
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None

# Compute the "maxDepth" of a tree -- the number of nodes
# along the longest path from the root node down to the
# farthest leaf node


def maxDepth(node):
    return balance(node) != -1


def balance(node):
    if node is None:
        return 0

    else:

        # Compute the depth of each subtree
        lDepth = balance(node.left)
        if (lDepth == -1):
            return -1
        rDepth = balance(node.right)
        if (rDepth == -1):
            return -1

        # Use the larger one
        # if (lDepth > rDepth):
        #     return lDepth+1
        # else:
        #     return rDepth+1
        if (abs(lDepth - rDepth) > 1):
            return -1
        res = max(lDepth, rDepth) + 1
        return res


# Driver program to test above function
root = Node(1)
root.left = Node(2)
root.right = Node(3)
root.left.left = Node(4)
root.left.right = Node(5)
#root.left.left.left = Node(15)


print("Height of tree is %d" % (maxDepth(root)))
