def maxOfSubarray(arr, n, k):
    # write your code here
    # return a list denoting the answer.
    i, j = 0, 0
    out = []
    res = []
    while j < n:
        out.append(arr[j])
        if j - i + 1 < k:
            j += 1
        elif j - i + 1 == k:
            val = max(out)
            res.append(val)
            out.pop(0)
            i += 1
            j += 1

    return res


# arr = [10, 7, 8, 11]
# n = len(arr)
# k = 2

# print(maxOfSubarray(arr, n, k))
if __name__ == "__main__":
    arr = []
    T = int(input())
    for i in range(T):
        val = input().split()
    print(val, type(val))
    print(type(val[0]), type(val[1]))
    for i in range(int(val[0])):
        # print(i)
        arr.append(input())

    print(arr)
    print(type(arr[0]))
    arr = list(map(int, arr))
    print(arr)
    print(type(arr))

    print(maxOfSubarray(arr, int(val[0]), int(val[1])))
