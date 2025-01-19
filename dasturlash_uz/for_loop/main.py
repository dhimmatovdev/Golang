n = input()
d = len(n)
ls = []
for i in n:
    ls.append(pow(int(i),d))
print(sum(ls))
