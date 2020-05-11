cont = 0
inter = 0
gremio = 0
empates = 0
x = 0


def main():
    global cont
    global inter
    global gremio
    global empates
    global x
    on_off = True
    while on_off:
        i, g = map(int, input().split(' '))
        cont += 1
        if i == g:
            empates += 1
        elif i > g:
            inter += 1
        else:
            gremio += 1
        print('Novo grenal (1-sim 2-nao)')
        x = int(input())
        if x == 2:
            on_off = False
        elif x == 1:
            on_off = True


main()
print('%d grenais' % cont)
print('Inter:%d' % inter)
print('Gremio:%d' % gremio)
print('Empates:%d' % empates)
if inter > gremio:
    print('Inter venceu mais')
elif inter < gremio:
    print('Gremio venceu mais')
else:
    print('Nao houve vencedor')
