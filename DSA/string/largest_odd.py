def largestOddNumber(num):
    arr = list(num)
    for i in range(len(arr) - 1, -1, -1):
        if arr[i] == 0:
            arr.pop()
        elif int(arr[i]) % 2 != 0:
            out = ''.join(arr)
            return out
    max_out = ""
    for i in range(len(arr)):
        if int(arr[i]) % 2 != 0:
            max_out += max(str(arr[i]), max_out)
        return max_out


num = "592"
print(largestOddNumber(num))
