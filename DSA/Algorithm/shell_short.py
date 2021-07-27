def shell_short(arr):
    size = len(arr)
    gap = size//2

    while gap > 0:
        for i in range(gap, size):
            anchor = arr[i]
            j = i
            while j >= gap and arr[j-gap] > anchor:
                arr[j] = arr[j-gap]
                j -= gap
            arr[j] = anchor
        gap = gap//2


a = [21, 38, 29, 17, 4, 25, 11, 32, 9]
shell_short(a)
print(a)
