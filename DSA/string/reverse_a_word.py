'''
Reverse Words in a String

'''


def reverse_word(s):
    words = s.split(' ')
    # print(words)
    string = []
    for word in words:
        if word != '':
            string.insert(0, word)
    out = " ".join(string)
    return out


s = "i     like this program very much"
print(reverse_word(s))
