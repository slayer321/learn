class TreeNode:
    def __init__(self, data):
        self.data = data
        self.childrean = []
        self.parent = None

    def add_child(self, child):
        self.parent = self
        self.childrean.append(child)


def build_product():
    root = TreeNode("Electronics")

    laptop = TreeNode("Laptop")
    root.add_child(laptop)

    return root


if __name__ == "__main__":
    root = build_product()
    pass
