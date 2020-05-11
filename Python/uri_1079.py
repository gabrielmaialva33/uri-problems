n = int(input())
a = []
b = []
for i in range(n):
    a.append([float(j) for j in input().split()])

for i in range(0, n):
    b.append([])
    b[i] = a[i][0] * 2.0, a[i][1] * 3.0, a[i][2] * 5.0
for i in range(len(a)):
    print('%0.1f' % (sum(b[i]) / 10))