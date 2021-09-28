# def solve(n, k):
#     if n <= 1:
#         return "0"
#     res = "0"
#     out = s(n, k, res)
#     for i in range(len(out)):
#         if i == k-1:
#             return out[i]


# def s(n, k, res):
#     if n == 1:
#         return res
#     for i in res:
#         if i == "0":
#             res = "01"
#         else:
#             res += "10"

#     return s(n-1, k, res)


def solve(n, k):
    if n == 1 and k == 1:
        return 0

    mid = (pow(2, n-1) // 2)
    if k <= mid:
        return solve(n-1, k)
    else:
        return (solve(n-1, k-mid) ^ 1)


n = 2
k = 2

print(solve(n, k))
