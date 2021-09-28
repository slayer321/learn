

def solve(s):
    vo = ['a', 'e', 'i', 'o', 'u']
    for i in s:
        k = i.lower()
        if k in vo:
            res = s.replace(k, '')

    return res


s = "welcome to geeksforgeeks"
print(solve(s))
