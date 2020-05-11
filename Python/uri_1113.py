while True:
    x, y = map(int, input().split(" "))
    if x == y:
        break
    else:
        if x < y:
            print('Crescente')
        else:
            print('Decrescente')
