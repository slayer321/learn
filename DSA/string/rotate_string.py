'''
Rotate String

'''


'''
Using slicing method
'''


def fun(s, g):
    for i in range(0, len(s)):
        l = len(s)
        x = s[0]
        s = s[1:l]
        s += x
        if s == g:
            return True
    return False


'''
BruteForce solution 
time Complexity : O(n^2)
space Complexity : O(n)
'''


def rotate(v):
    for i in range(len(v) - 1):
        v[i] = v[i + 1]
    return v


def stringR(s, goal):
    n = len(s)
    v = list(s)
    for i in range(n):
        end = v[0]

        out = rotate(v)
        out.insert(n-1, end)
        out.pop()
        val = ''.join(out)
        if val == goal:
            print(val)
            return val


'''
Using KMP string matching Algo
'''


def compute_lps(g):
    lps_out = [0] * len(g)
    top = 0
    for i in range(1, len(g)):
        while top and g[top] != g[i]:
            top = lps_out[top - 1]
        if g[top] == g[i]:
            top += 1
            lps_out[i] = top

    return lps_out


def rotateString(s, goal):
    val = s+s
    pat_val = compute_lps(goal)
    print(pat_val)
    pati = 0
    for i, ch in enumerate(val):
        while pati and val[pati] != ch:
            pati = pat_val[pati - 1]
        if goal[pati] == ch:
            if pati == len(goal) - 1:
                return True
            else:
                pati += 1

    return False


s = 'abcde'
#v = list(s)
goal = 'abced'

print(rotateString(s, goal))
