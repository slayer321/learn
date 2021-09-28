def maximumMeetings(n, start, end):
    # code here
    l = []
    for i in range(n):
        l.append([end[i], start[i]])
    # print(l)

    l.sort(key=lambda i: i[0])
    # print(l)
    count = 1
    e = l[0][0]
    for i in range(1, n):
        if l[i][1] > e:
            count += 1
            e = l[i][0]

    return count


#start = [1, 3, 0, 5, 8, 5]
#end = [2, 10,  4, 6, 7, 9, 9]

start = [75250, 50074, 43659, 8931, 11273, 27545, 50879, 77924]
end = [112960, 114515, 81825, 93424, 54316, 35533, 73383, 160252]
n = len(start)

print(maximumMeetings(n, start, end))
