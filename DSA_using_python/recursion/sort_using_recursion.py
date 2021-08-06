'''
Sort using recursion
'''


def sort_(arr):
    if len(arr) == 1:
        return arr
    # hypothesis
    tmp = arr[-1]
    arr.pop()
    arr = sort_(arr)
    # induction
    arr = insertfun(arr, tmp)
    return arr


def insertfun(arr, tmp):
    if (len(arr) == 0 or arr[-1] <= tmp):
        arr.append(tmp)
        return arr
    val = arr[-1]
    arr.pop()
    arr = insertfun(arr, tmp)
    arr.append(val)
    return arr


arr = [5, 1, 1, 2, 0, 0]
print(sort_(arr))
