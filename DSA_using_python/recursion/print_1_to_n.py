'''
print the value from 1 to n / n to 1 using recursion

'''

# Print the value from 1 to n


def valout(n):
    if n == 0:
        return
    print("first", n)
    valout(n-1)
    print("second", n)
    #print("Third", n)


# Print value from n to 1

# def valout(n):

#     if n == 0:
#         return
#     print(n)
#     valout(n-1)


n = 5
valout(5)
