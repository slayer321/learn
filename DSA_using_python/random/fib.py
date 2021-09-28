# def solve(n):
#     a = 0
#     b = 1
#     sum_ = 0
#     count = 1
#     while count <= n:
#         print(sum_, end=" ")
#         count += 1
#         a = b
#         b = sum_
#         sum_ = a+b


def solve(n):
    if n <= 1:
        return n
    return solve(n-2) + solve(n-1)


n = 5
print(solve(n))


# def Fibonacci(n):

#     # Check if input is 0 then it will
#     # print incorrect input
#     if n < 0:
#         print("Incorrect input")

#     # Check if n is 0
#     # then it will return 0
#     elif n == 0:
#         return 0

#     # Check if n is 1,2
#     # it will return 1
#     elif n == 1 or n == 2:
#         return 1

#     else:
#         return Fibonacci(n-1) + Fibonacci(n-2)


# # Driver Program
# print(Fibonacci(5))
