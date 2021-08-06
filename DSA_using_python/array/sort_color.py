# Sort Colors


def sortColors(self, arr: List[int]) -> None:
    """
        Do not return anything, modify nums in-place instead.
        """
    i = 0
    j = len(arr) - 1
    mid = 0

    while mid <= j:
        if arr[mid] == 0:
            arr[i], arr[mid] = arr[mid], arr[i]
            i += 1
            mid += 1
        elif arr[mid] == 1:
            mid += 1
        elif arr[mid] == 2:
            arr[mid], arr[j] = arr[j], arr[mid]
            j -= 1

    return arr
