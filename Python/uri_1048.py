n = float(input())
if 0 <= n <= 400.00:
    t = 0.15
elif 400.01 <= n <= 800.00:
    t = 0.12
elif 800.01 <= n <= 1200.00:
    t = 0.10
elif 1200.01 <= n <= 2000.00:
    t = 0.07
else:
    t = 0.04
print('Novo salario: %0.2f' % (n + (n * t)))
print('Reajuste ganho: %0.2f' % (abs(n - (n + (n * t)))))
print('Em percentual: %0.0f %%' % (t * 100))