v = []
x = []
for i in range(0, 6):
    v.append(float(input()))
    if v[i] > 0:
        x.append(v[i])
print('{} valores positivos'.format(len(x)))
print('%0.1f' % (sum(x) / len(x)))