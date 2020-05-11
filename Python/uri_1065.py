v = []
x = 0
for i in range(0, 5):
    v.append(int(input()))
    if v[i] % 2 == 0:
        x += 1
print('{} valores pares'.format(x))