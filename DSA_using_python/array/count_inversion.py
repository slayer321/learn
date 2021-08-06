def inversionCount(arr, n):
    # Your Code Here
    if len(arr) <= 1:
        return arr
    mid = len(arr) // 2
    left = arr[:mid]
    right = arr[mid:]

    inversionCount(left, len(left))
    inversionCount(right, len(right))

    i = j = k = 0
    count = 0
    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            arr[k] = left[i]
            i += 1
        else:
            arr[k] = right[j]
            j += 1
        count += 1
        k += 1

    while i < len(left):
        arr[k] = left[i]
        i += 1
        k += 1
        # count += 1

    while j < len(right):
        arr[k] = right[j]
        j += 1
        k += 1
        # count += 1
    return count


arr = [2, 3, 4, 5, 6]
n = len(arr)
print(inversionCount(arr, n))
