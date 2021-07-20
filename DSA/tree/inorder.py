# Binary Tree Inorder Traversal


# recursive solution

class Solution:
    def inorder(self, root, l):
        if root:
            self.inorder(root.left, l)
            l.append(root.val)
        # if root.right:
            self.inorder(root.right, l)

    def inorderTraversal(self, root: TreeNode) -> List[int]:
        l = []
        self.inorder(root, l)
        return l


# Iterating method using stack


def inorderTraversal(self, root: TreeNode) -> List[int]:
    stack, res = [], []
    curr = root
    while curr or stack:
        while curr:
            stack.append(curr)
            curr = curr.left
        curr = stack.pop()
        res.append(curr.val)
        curr = curr.right
    return res
