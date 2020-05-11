x = 1
y = 7
for i in range(0, 5):
    for j in range(0, 3):
        print('I=%d J=%d' % (x, y - j))
    x += 2
    y += 2
