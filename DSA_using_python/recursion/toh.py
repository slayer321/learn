'''
Tower Of Hanoi
'''


def toh(N, fromm, to, aux):
    val = (2**N) - 1
    if N == 1:
        print("move disk 1 from rod ", fromm, "to rod", to)

        return val

    toh(N-1, fromm, aux, to)
    print("move disk", N, " from rod ", fromm, "to rod ", to)
    toh(N-1, aux, to, fromm)

    return val


N = 3
print(toh(N, 1, 3, 2))
