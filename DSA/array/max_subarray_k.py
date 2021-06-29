'''
Maximum Sum Subarray of size K
'''


def max_k(arr, k, n):
    i = 0
    j = 0
    sum_ = 0
    out = 0
    while j < n:
        sum_ += arr[j]
        if j-i+1 < k:
            j += 1
        elif j-i+1 == k:
            out = max(out, sum_)
            sum_ -= arr[i]
            i += 1
            j += 1
    return out


arr = [100, 200, 300, 400]
n = 4
k = 2

print(max_k(arr, k, n))
