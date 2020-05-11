v = []
x = 0
for i in range(0, 2):
    v.append(int(input()))
v.sort()
n = v[0] + 1
m = v[1]
for i in range(n, m):
    if i % 2 != 0:
        x += i
print(x)