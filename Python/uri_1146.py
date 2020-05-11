x = 1
while x != 0:
    x = int(input())
    for i in range(1, x + 1):
        if i <= (x - 1):
            print("%d " % i, end="")
        else:
            print('%d' % i)
