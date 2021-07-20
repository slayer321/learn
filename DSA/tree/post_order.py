# Binary Tree Postorder Traversal


class Solution:

    def postorder(self, root, l):
        if root:

            self.postorder(root.left, l)
            self.postorder(root.right, l)
            l.append(root.val)

    def postorderTraversal(self, root: TreeNode) -> List[int]:
        l = []
        self.postorder(root, l)
        return l


# Iterating method using stack
def postorderTraversal(self, root: TreeNode) -> List[int]:
    res = []
    stack = [root]
    while stack:
        node = stack.pop()
        if node:
            res.append(node.val)
            stack.append(node.left)
            stack.append(node.right)

    return res[::-1]
