
def fg(arr):
    for i in range(1, 10):
        arr.append(i)
    return arr


def solve():
    arr = []
    fg(arr)
    return arr


print(solve())
