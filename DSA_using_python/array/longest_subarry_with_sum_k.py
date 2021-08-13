# Longest Sub-Array with Sum K


def lenOfLongSubarr(A, N, K):
    # Complete the function
    i = 0
    j = 0
    sum_ = 0
    max_val = 0

    while j < N:
        sum_ += A[j]
        if sum_ < K:
            j += 1
        elif sum_ == K:
            max_val = max(max_val, j-i+1)
            j += 1
        elif sum_ > K:
            while (sum_ > K):
                sum_ -= A[i]
                i += 1
            j += 1

    return max_val


nums = [-13, 0, 6, 15, 16, 2, 15, - 12, 17, - 16, 0, - 3, 19, - 3, 2, - 9, - 6]
N = len(nums)
K = 15

print(lenOfLongSubarr(nums, N, K))
