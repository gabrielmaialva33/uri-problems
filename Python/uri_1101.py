on = True
soma = 0
while on:
    m, n = map(int, input().split(" "))
    soma = 0
    if (m <= 0) or (n <= 0):
        break
    maxx = max(m, n)
    minn = min(m, n)
    for i in range(minn, maxx + 1):
        print("%d " % i, end="")
        soma += i
    print("Sum=%d" % soma)
