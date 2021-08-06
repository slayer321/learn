# Factorial of the number


def fact(n):
    if n == 0:
        return 1
    #print("first", n)
    val = n * fact(n-1)
    #print("second", n)
    return val


n = 5
print(fact(n))
