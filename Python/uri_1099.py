v = []
s = 0
n = int(input())
for i in range(0, n):
    s = 0
    x, y = map(int, input().split(" "))
    big = max(x, y)
    small = min(x, y)
    for j in range(small + 1, big):
        if j % 2 != 0:
            s += j
    v.append(s)
for i in range(0, n):
    print(v[i])
