visited = []


def dfs(graph, node):
    if node not in visited:
        print(node, end=" ")
    visited.append(node)
    for neg in graph[node]:
        dfs(graph, neg)


graph = {
    "A": ['B', 'C'],
    'B': ['D', 'E'],
    'C': ['F'],
    'D': [],
    'E': ['F'],
    'F': []
}

dfs(graph, "A")
