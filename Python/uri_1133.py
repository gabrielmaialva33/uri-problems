x = int(input())
y = int(input())
maxx = max(x, y)
minn = min(x, y)
for i in range(minn + 1, maxx):
    if i % 5 == 2 or i % 5 == 3:
        print(i)
