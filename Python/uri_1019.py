n = int(input())  # seconds
h = n // 3600
aux = n % 3600
m = aux // 60
s = aux % 60
print('{}:{}:{}'.format(h, m, s))
