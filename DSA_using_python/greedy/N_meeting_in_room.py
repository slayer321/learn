def maximumMeetings(n, start, end):
    # code here
    count = 0
    start.sort()
    # print(end)
    for i in range(n-1):
        if end[i] < start[i+1]:
            count += 2
    return count


#start = [1, 3, 0, 5, 8, 5]
#end = [2, 10,  4, 6, 7, 9, 9]

start = [75250, 50074, 43659, 8931, 11273, 27545, 50879, 77924]
end = [112960, 114515, 81825, 93424, 54316, 35533, 73383, 160252]
n = len(start)

print(maximumMeetings(n, start, end))
