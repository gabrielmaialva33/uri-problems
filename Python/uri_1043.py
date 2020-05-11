a, b, c = map(float, input().split(' '))
if a + b > c and a + c > b and b + c > a:
    p = a + b + c
    print('Perimetro = %0.1f' % p)
else:
    s = ((a + b) * c) / 2.0
    print('Area = %0.1f' % s)
