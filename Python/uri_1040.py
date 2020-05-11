n1, n2, n3, n4 = input().split(' ')
n1 = float(n1)
n2 = float(n2)
n3 = float(n3)
n4 = float(n4)
m = ((n1 * 2) + (n2 * 3) + (n3 * 4) + n4) / 10.0
print('Media: %0.1f' % m)
if m >= 7.0:
    print('Aluno aprovado.')
elif m < 5.0:
    print('Aluno reprovado.')
else:
    print('Aluno em exame.')
    n5 = float(input())
    print('Nota do exame: %0.1f' % n5)
    m = (m + n5) / 2.0
    if m >= 5.0:
        print('Aluno aprovado.')
    else:
        print('Aluno reprovado.')
    print('Media final: %0.1f' % m)
