class Node:
    def __init__(self, data):
        self.val = data
        self.left = None
        self.right = None


def delete_val(root, val):
    if (root == None):
        return None
    if (root.val == val):
        return helper(root)
    dummy = root
    while root:
        if (root.val > val):
            if (root.left != None and root.left.val == val):
                root.left = helper(root.left)
                break
            else:
                root = root.left
        else:
            if (root.right != None and root.right.val == val):
                root.right = helper(root.right)
                break
            else:
                root = root.right

    return dummy


def helper(root):
    if (root.left == None):
        return root.right
    elif (root.right == None):
        return root.left

    right_child = root.right
    last_right = check_last_right(root.left)
    last_right.right = right_child
    return root.left


def check_last_right(root):
    if root.right == None:
        return root
    return check_last_right(root.right)


if __name__ == "__main__":
    root = Node(5)
    root.left = Node(3)
    root.right = Node(6)
    root.left.left = Node(2)
    root.left.right = Node(4)
    root.right.right = Node(7)

    print(delete_val(root, 5))
