# Roman to Integer


def rToi(s):
    dict_ = {'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
    res, prev = 0, 0
    for i in s[::-1]:
        print(dict_[i])
        if dict_[i] >= prev:
            res += dict_[i]
        else:
            res -= dict_[i]
        prev = dict_[i]

    return res


s = "IX"
print(rToi(s))
