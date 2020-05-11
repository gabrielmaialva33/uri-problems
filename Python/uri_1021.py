n = float(input())
n *= 1000
print('NOTAS:')
print('%d nota(s) de R$ 100.00' % (n / 100000))
n %= 100000
print('%d nota(s) de R$ 50.00' % (n / 50000))
n %= 50000
print('%d nota(s) de R$ 20.00' % (n / 20000))
n %= 20000
print('%d nota(s) de R$ 10.00' % (n / 10000))
n %= 10000
print('%d nota(s) de R$ 5.00' % (n / 5000))
n %= 5000
print('%d nota(s) de R$ 2.00' % (n / 2000))
n %= 2000
print('MOEDAS:')
print('%d moeda(s) de R$ 1.00' % int(n / 1000))
n %= 1000
print('%d moeda(s) de R$ 0.50' % int(n / 500))
n %= 500
print('%d moeda(s) de R$ 0.25' % int(n / 250))
n %= 250
print('%d moeda(s) de R$ 0.10' % int(n / 100))
n %= 100
print('%d moeda(s) de R$ 0.05' % int(n / 50))
n %= 50
print('%d moeda(s) de R$ 0.01' % int(n / 10))
