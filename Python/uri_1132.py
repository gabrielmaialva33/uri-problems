x = int(input())
y = int(input())
s = 0
aux = max(x, y)
x = min(x, y)
for i in range(x, aux + 1):
    if i % 13 != 0:
        s += i
print(s)
