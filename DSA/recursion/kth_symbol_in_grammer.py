'''
Kth Symbol in Grammar

'''


def kSymbol(n, k):
    if n == 1 and k == 1:
        return 0
    mid = (pow(2, n-1)) // 2
    if k <= mid:
        return kSymbol(n-1, k)
    else:
        # bitwise XOR
        return (kSymbol(n-1, k - mid) ^ 1)


n = 2
k = 2
print(kSymbol(n, k))
