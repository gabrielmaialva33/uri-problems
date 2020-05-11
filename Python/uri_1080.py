v = []
x = 0
p = 0
for i in range(0, 100):
    v.append(int(input()))
    if v[i] > x:
        x = v[i]
        p = i
print('{}'.format(x))
print('{}'.format(p + 1))