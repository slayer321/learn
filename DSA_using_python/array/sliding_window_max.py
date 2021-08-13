

def slide(nums, k):
    res = []
    queue = []
    n = len(nums)
    i = 0
    j = 0

    while j < n:
        # if len(queue) != 0:
        while (len(queue) > 0 and queue[-1] < nums[j]):
            queue.pop()
        queue.append(nums[j])

        if j - i + 1 < k:
            j += 1
        elif j - i + 1 == k:
            res.append(queue[0])

            if queue[0] == nums[i]:
                queue.pop(0)
            i += 1
            j += 1

    return res


#nums = [1, 3, -1, -3, 5, 3, 6, 7]
nums = [9, 11]
k = 2

print(slide(nums, k))
