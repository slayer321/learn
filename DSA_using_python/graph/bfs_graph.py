queue = []
visited = []


def bfs(graph, node):
    queue.append(node)
    visited.append(node)
    while queue:
        curr = queue.pop(0)
        print(curr, end=" ")

        for neg in graph[curr]:
            if neg not in visited:
                queue.append(neg)
                visited.append(neg)


graph = {
    "A": ['B', 'C'],
    'B': ['D', 'E'],
    'C': ['F'],
    'D': [],
    'E': ['F'],
    'F': []
}

bfs(graph, "A")
