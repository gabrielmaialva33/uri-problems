a, b, c = map(float, input().split(' '))
v = [a, b, c]
v.sort(reverse=True)
a, b, c = v
if a >= (b + c):
    print('NAO FORMA TRIANGULO')
else:
    if a ** 2 == (b ** 2 + c ** 2):
        print('TRIANGULO RETANGULO')
    if a ** 2 > b ** 2 + c ** 2:
        print('TRIANGULO OBTUSANGULO')
    if a ** 2 < b ** 2 + c ** 2:
        print('TRIANGULO ACUTANGULO')
    if a == b == c:
        print('TRIANGULO EQUILATERO')
    elif a == b or c == a or b == c:
        print('TRIANGULO ISOSCELES')
