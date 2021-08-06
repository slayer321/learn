def merge(intervals):
    intervals.sort()
    i = 1
    while i < len(intervals):
        if intervals[i][0] <= intervals[i-1][1]:
            intervals[i] = [intervals[i-1][0],
                            max(intervals[i][1], intervals[i-1][1])]
            del intervals[i-1]
        else:
            i += 1
    return intervals


arr = [[1, 3], [2, 6], [8, 10], [15, 18]]
print(merge(arr))
