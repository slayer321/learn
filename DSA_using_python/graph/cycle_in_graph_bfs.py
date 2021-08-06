def cycle(i, vis, adj):
    queue = []
    queue.append([i, -1])
    vis[i] = True
    # print("queue: ",queue)
    # node = queue[0].pop(0)
    # par = queue[0].pop(0)
    # print(node , "-->", par)
    while queue:
        node = queue[0].pop(0)
        par = queue[0].pop(0)
        queue.pop(0)
        # print(node)
        # print(par)
        for neg in adj[node]:
            if vis[neg] == False:
                queue.append([neg, par])
                vis[neg] = True
            elif (par != neg):
                return True
    return False

    # Function to detect cycle in an undirected graph.


def isCycle(V, adj):
    vis = [False] * V
    # print(adj)
    for i in range(V):
        if vis[i] == False:
            if (cycle(i, vis, adj)):
                return True
    return False


adj = [[4], [2, 4], [1, 3], [2, 4], [0, 1, 3]]
v = len(adj)

print(isCycle(v, adj))
