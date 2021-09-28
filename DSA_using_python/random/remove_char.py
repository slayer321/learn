def solve(s):
    #res = ""
    out = ""
    for i in s:
        # print(i)
        # print(i.isalpha())
        if i.isalpha() == False:
            pass
            #res = s.replace(i, '')
        else:
            out += i

    return out


s = "$Gee*k;s..fo, r'Ge^eks?"
print(solve(s))
