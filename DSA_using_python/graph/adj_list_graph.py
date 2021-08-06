def bfsOfGraph(V, adj):
    # code here
    queue = []
    vis = [False] * (V)
    res = []
    # print(type(adj))
    # print(adj)
    for i in range(V):
        if vis[i] == False:
            queue.append(i)
            vis[i] = True

            while queue:
                curr = queue.pop(0)
                res.append(curr)
                # print(adj[curr])
                for neg in adj[curr]:
                    if vis[neg] == False:
                        vis[neg] = True
                        queue.append(neg)

    return res


#adj = [[1, 2, 3], [], [4], [], []]
adj = [[2, 5], [5, 6, 8], [], [4, 5], [7], [7], [], [], []]
v = len(adj)

# 0257

#[0, 2, 5, 7, 1, 6, 8, 3, 4]

print(bfsOfGraph(v, adj))


'''
0 --> [2, 5]
1 --> [5, 6, 8]
2 --> []
3 --> [4, 5]
4 --> [7]
5 --> [7]
'''
