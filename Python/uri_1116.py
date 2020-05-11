n = int(input())
for i in range(0, n):
    x, y = map(int, input().split(" "))
    if y == 0:
        print('divisao impossivel')
    else:
        print('%.1f' % float(x / y))
