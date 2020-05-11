lista = list(map(int, input().split()))
aux = []
s = 0
for i in lista:
    if i > 0:
        aux.append(i)

for i in range(0, aux[1]):
    s += aux[0] + i
print(s)
