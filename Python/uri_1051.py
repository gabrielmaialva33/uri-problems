salary = float(input())
R = 0.0
if 0 <= salary < 2000.00:
    print('Isento')
else:
    if 2000.01 <= salary <= 3000.00:
        R = (salary - 2000.00) * 0.08
    elif 3000.01 <= salary <= 4500.00:
        R = (1000.00 * 0.08) + ((salary - 3000) * 0.18)
    else:
        R = (1000 * 0.08) + (1500 * 0.18) + ((salary - 4500) * 0.28)
    print('R$ %0.2f' % R)