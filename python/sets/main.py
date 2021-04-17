x = {1, 1, '1', '1'}

print(x)

x = [1, 2, 3, 1, '1', 1]

x = set(x)
print(x)

x.add(22)

print(x)

x.remove('1')

22 in x

y = {3}

print(y & x)

print(y | x)

print(y.union(x))
print(y.issubset(x))