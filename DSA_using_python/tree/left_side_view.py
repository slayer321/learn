# Binary Tree Left Side View

def leftSideView(self, root: TreeNode) -> List[int]:
    if root == None:
        return
    res = []
    queue = [root]
    while queue:
        size = len(queue)
        for i in range(size):
            node = queue.pop(0)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
            if i == 0:
                res.append(node.val)

    return res
