v = []
w = 0
x = 0
y = 0
z = 0
for i in range(0, 5):
    v.append(int(input()))
    if v[i] % 2 == 0:
        w += 1
    if v[i] % 2 != 0:
        x += 1
    if v[i] > 0:
        y += 1
    if v[i] < 0:
        z += 1
print('{} valor(es) par(es)'.format(w))
print('{} valor(es) impar(es)'.format(x))
print('{} valor(es) positivo(s)'.format(y))
print('{} valor(es) negativo(s)'.format(z))