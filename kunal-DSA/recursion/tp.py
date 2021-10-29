#arr = [15, 23, 33, 45, 56, 15, 23, 33]
arr = ['3z4', '3z4', '3Z4', '3Z4', '3z4', '3z4', '3Z4', '3Z4']
dic = {}
for i in range(len(arr)):
    if arr[i] not in dic:
        dic[i] = arr[i]
    else:
        arr.pop(i)
print(arr)
