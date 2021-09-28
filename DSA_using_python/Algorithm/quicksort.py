def swap(array, first, second):
    array[first], array[second] = array[second], array[first]


def quicksort(array):
    helper(array, 0, len(array) - 1)
    return array


def helper(array, startIdx, endIdx):
    if startIdx >= endIdx:
        return
    pivotIdx = startIdx
    leftIdx = startIdx+1
    rightIdx = endIdx

    while leftIdx <= rightIdx:
        if array[leftIdx] > array[pivotIdx] and array[rightIdx] < array[pivotIdx]:
            swap(array, leftIdx, rightIdx)
        if array[leftIdx] <= array[pivotIdx]:
            leftIdx += 1
        if array[rightIdx] >= array[pivotIdx]:
            rightIdx -= 1
    swap(array, pivotIdx, rightIdx)

    leftSubarrayIsSmaller = rightIdx - 1 - startIdx < endIdx - (rightIdx + 1)
    if leftSubarrayIsSmaller:
        helper(array, startIdx, rightIdx-1)
        helper(array, rightIdx+1, endIdx)
    else:
        helper(array, rightIdx+1, endIdx)

        helper(array, startIdx, rightIdx-1)


array = [43, 1, 23, 5, 4, 99]
print(quicksort(array))
