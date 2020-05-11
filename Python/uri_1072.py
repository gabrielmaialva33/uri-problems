n = int(input())
v = []
IN = 0
OUT = 0
for i in range(0, n):
    v.append(int(input()))
    if 10 <= v[i] <= 20:
        IN += 1
    else:
        OUT += 1
print('{} in'.format(IN))
print('{} out'.format(OUT))