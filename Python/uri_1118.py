def main():
    count = 2
    soma = 0
    while True:
        x = float(input())
        if 0 <= x <= 10:
            count -= 1
            soma += x
            if count == 0:
                print('media = %.2f' % (soma / 2.0))
                while True:
                    print('novo calculo (1-sim 2-nao)')
                    aux = int(input())
                    if aux == 2:
                        exit()
                    elif aux == 1:
                        main()
                    else:
                        continue
        else:
            print('nota invalida')


main()
