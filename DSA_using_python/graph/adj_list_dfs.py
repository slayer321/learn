def dfs(i, vis, adj, res):
    res.append(i)
    vis[i] = True
    for neg in adj[i]:
        if vis[neg] == False:
            dfs(neg, vis, adj, res)


def dfsOfGraph(V, adj):
    # code here
    res = []
    vis = [False] * V
    for i in range(V):
        if vis[i] == False:
            dfs(i, vis, adj, res)
    return res


adj = [[1, 2, 3], [0], [0, 4], [0], [2]]
v = len(adj)
print(dfsOfGraph(v, adj))
