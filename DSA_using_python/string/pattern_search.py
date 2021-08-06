
'''
Search Pattern (KMP-Algorithm) 
'''


def compute_lps(pat):
    lps = [0] * len(pat)
    top = 0
    for i in range(1, len(pat)):
        while top and pat[i] != pat[top]:
            top = lps[top - 1]

        if pat[i] == pat[top]:
            top += 1
            lps[i] = top

    return lps


def search(patt, s):
    match_i = []
    pattern_lps = compute_lps(patt)

    patterni = 0
    for i, ch in enumerate(s):

        while patterni and patt[patterni] != ch:
            patterni = pattern_lps[patterni - 1]

        if patt[patterni] == ch:
            if patterni == len(patt) - 1:
                match_i.append((i - patterni) + 1)
                patterni = pattern_lps[patterni]
            else:
                patterni += 1
    # print(len(match_i))
    if len(match_i) == 0:
        return [-1]

    return match_i


patt = 'q'
s = 'hqg'

print(search(patt, s))
