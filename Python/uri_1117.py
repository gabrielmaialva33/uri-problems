cont = 2
s = 0.0
while cont > 0:
    x = float(input())
    if 0 <= x <= 10:
        cont -= 1
        s += x
        if cont == 0:
            media = s / 2.0
            print('media = {}'.format(media))
    else:
        print('nota invalida')
