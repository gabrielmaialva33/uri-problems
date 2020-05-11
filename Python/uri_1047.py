hi, mi, hf, mf = map(int, input().split(' '))
ti = (hi * 60) + mi
tf = (hf * 60) + mf
if ti >= tf:
    t = abs(ti - tf - 1440)
else:
    t = abs(tf - ti)
print('O JOGO DUROU {} HORA(S) E {} MINUTO(S)'.format(t // 60, t % 60))