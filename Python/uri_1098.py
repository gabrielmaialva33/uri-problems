x = 0
y = 1
for i in range(0, 11):
    for j in range(0, 3):
        if x == 0 or x == 1.0 or x > 1.8:
            print('I=%0.0f J=%0.0f' % (x, y + j))
        else:
            print('I=%0.1f J=%0.1f' % (x, y + j))
    x += 0.2
    y += 0.2
