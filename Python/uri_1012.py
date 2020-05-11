import math

a, b, c = input().split(' ')
a = float(a)
b = float(b)
c = float(c)
triangulo = (a * c) / 2.0
circulo = 3.14159 * math.pow(c, 2)
trapezio = ((a + b) / 2.0) * c
quadrado = b * b
retangulo = a * b
print('TRIANGULO: %0.3f' % triangulo)
print('CIRCULO: %0.3f' % circulo)
print('TRAPEZIO: %0.3f' % trapezio)
print('QUADRADO: %0.3f' % quadrado)
print('RETANGULO: %0.3f' % retangulo)
