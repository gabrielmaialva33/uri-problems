n = int(input())
y = n // 365
n = n % 365
m = n // 30
d = n % 30
print('{} ano(s)'.format(y))
print('{} mes(es)'.format(m))
print('{} dia(s)'.format(d))
