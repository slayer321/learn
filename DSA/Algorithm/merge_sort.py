def merge_sort(arr, n):
    if n <= 1:
        return arr
    mid = len(arr) // 2
    left = arr[:mid]
    right = arr[mid:]

    merge_sort(left, len(left))
    merge_sort(right, len(right))

    merge_two_sorted_lists(left, right, arr)


def merge_two_sorted_lists(a, b, arr):
    #sorted_list = []
    len_a = len(a)
    len_b = len(b)
    i = 0
    j, k = 0, 0
    count = 0
    while i < len_a and j < len_b:
        if a[i] <= b[j]:
            arr[k] = a[i]
            i += 1
            count += 1
        elif a[i] >= b[j]:
            arr[k] = b[j]
            j += 1
            count += 1

        print(count)
        k += 1
    while i < len_a:
        arr[k] = a[i]
        i += 1
        k += 1

    while j < len_b:
        arr[k] = b[j]
        j += 1
        k += 1


if __name__ == "__main__":
    #arr = [10, 3, 15, 7, 8, 23, 98, 29]
    arr = [2, 4, 1, 3, 5]
    merge_sort(arr, len(arr))
    print(arr)
