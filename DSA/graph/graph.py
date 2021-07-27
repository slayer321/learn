edges = [("A", "B"), ("A", "C"), ("B", "C"),
         ("B", "D"), ("B", "E"), ("D", "E"), ("C", "D")]


class Graph:
    def __init__(self, Nodes, is_directed=False):
        self.nodes = Nodes
        self.adj_list = {}
        self.is_directed = is_directed

        for node in self.nodes:
            self.adj_list[node] = []

    def add_edge(self, v, e):
        self.adj_list[v].append(e)
        if not self.is_directed:
            self.adj_list[e].append(v)

    def degree_vertex(self, node):
        degree = len(self.adj_list[node])
        return degree

    def print_lst(self):
        for node in self.nodes:
            print(node, ":", self.adj_list[node])


if __name__ == "__main__":
    nodes = ["A", "B", "C", "D", "E"]
    graph = Graph(nodes, is_directed=False)
    for v, e in edges:
        graph.add_edge(v, e)
    graph.print_lst()
    # print(graph.degree_vertex("B"))
