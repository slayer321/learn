# Binary Tree Level Order Traversal


def levelOrder(self, root: TreeNode) -> List[List[int]]:
    if root == None:
        return
    queue = [root]
    res = []
    while queue:
        lev = []
        for _ in range(len(queue)):
            node = queue.pop(0)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
            lev.append(node.val)
        res.append(lev)

    return res
