#  Binary Tree Preorder Traversal


class Solution:

    def pre_order(self, root, l):
        if root:
            l.append(root.val)
            self.pre_order(root.left, l)
            self.pre_order(root.right, l)

    def preorderTraversal(self, root: TreeNode) -> List[int]:
        l = []
        self.pre_order(root, l)
        return l


# Iterating method using stack
def preorderTraversal(self, root: TreeNode) -> List[int]:
    res = []
    stack = [root]
    while stack:
        node = stack.pop()
        if node:
            res.append(node.val)
            stack.append(node.right)
            stack.append(node.left)

    return res
