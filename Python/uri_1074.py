n = int(input())
v = []
for i in range(0, n):
    v.append(int(input()))
for i in range(0, len(v)):
    if v[i] == 0:
        print('NULL')
    elif v[i] % 2 == 0:
        if v[i] > 0:
            print('EVEN POSITIVE')
        else:
            print('EVEN NEGATIVE')
    else:
        if v[i] > 0:
            print('ODD POSITIVE')
        else:
            print('ODD NEGATIVE')