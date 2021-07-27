def in_sort(arr):
    for i in range(1, len(arr)):
        j = i - 1
        anchor = arr[i]

        while j >= 0 and anchor < arr[j]:
            arr[j+1] = anchor
            j = j - 1
        arr[j+1] = anchor


a = [11, 9, 29, 7, 2, 15, 28]
in_sort(a)
print(a)
