from typing import Dict


def isValid(s):
    st = []
    dict_ = {')': '(',
             ']': '[',
             '}': '{'}
    for i in s:
        if i == '(' or i == '[' or i == '{':
            st.append(i)
        else:
            if len(st) > 0 and st[-1] == dict_[i]:
                st.pop()
            else:
                return False

    if len(st) == 0:
        return True
    return False


s = "()[]{}"
print(isValid(s))
