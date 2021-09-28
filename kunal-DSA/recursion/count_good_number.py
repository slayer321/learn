def check_prime(n):
    if n > 1:
        for i in range(2, n):
            if n % i == 0:
                return 0
    else:
        return 0
    return 1
