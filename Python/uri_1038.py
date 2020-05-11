cod, qtd = input().split(' ')
cod = int(cod) - 1
qtd = int(qtd)
v = [4.00, 4.50, 5.00, 2.00, 1.50]
total = v[cod] * qtd
print('Total: R$ %0.2f' % total)
