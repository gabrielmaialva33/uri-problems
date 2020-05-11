i, f = map(int, input().split(' '))

if i >= f:
    t = abs(i - 24) + f
else:
    t = abs(i - f)
print('O JOGO DUROU {} HORA(S)'.format(t))
