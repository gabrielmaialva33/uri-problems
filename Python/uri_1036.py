import math

a, b, c = input().split(' ')
a = float(a)
b = float(b)
c = float(c)
x = (b ** 2 - 4 * a * c)
if x < 0 or a <= 0:
    print('Impossivel calcular')
else:
    d = math.sqrt(x)
    r1 = ((-b + d) / (2 * a))
    r2 = ((-b - d) / (2 * a))
    print('R1 = %0.5f' % r1)
    print('R2 = %0.5f' % r2)
