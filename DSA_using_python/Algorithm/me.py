def merge(arr1, arr2, n, m):
    # code here
    sorted_arr = []
    i = 0
    j = 0
    #print(arr1 , n)
    #print(arr2, m)

    while i < n and i < m:
        if arr1[i] <= arr2[j]:
            print("arr1->", arr1[i])
            sorted_arr.append(arr1[i])
            i += 1
        elif arr1[i] >= arr2[j]:
            print("arr2->", arr2[j])
            sorted_arr.append(arr2[j])
            j += 1
    print("i-->", i)
    print("j-->", j)
    print(sorted_arr)
    while i < n:
        print(i)
        sorted_arr.append(arr1[i])
        i += 1
    while j < m:
        print(j)
        sorted_arr.append(arr2[j])
        j += 1

    return sorted_arr


a = [1, 3, 5, 7]
b = [0, 2, 6, 8, 9]
print(merge(a, b, len(a), len(b)))
