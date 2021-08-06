class Solution:
    def __init__(self):
        pass

    def findAnagrams(self, s, p):
        dic = {}
        for letter in p:
            if letter not in dic:
                dic[letter] = 0
            dic[letter] += 1
        #print("dic", dic)
        k = len(p)
        count_dic = len(dic)
        i, j = 0, 0
        res = []
        while j < len(s):
            if s[j] in dic:
                dic[s[j]] -= 1
                if dic[s[j]] == 0:
                    count_dic -= 1

            if j - i + 1 < k:
                j += 1
            elif j - i + 1 == k:
                if count_dic == 0:
                    res.append(i)
                # print(dic)
                #print("out", res)
                # print(dic)
                if s[i] in dic:
                    dic[s[i]] += 1
                    # print(dic[s[i]])
                    if dic[s[i]] == 1:
                        count_dic += 1
                i += 1
                j += 1

        return res


obj = Solution()
s = "cbaebabacd"
p = "abc"
print(obj.findAnagrams(s, p))
