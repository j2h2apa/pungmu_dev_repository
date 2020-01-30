# find the sum of all the multiples 3 or 5 below 1000

# case 1
total = 0

# multiple1 = 0
# for i in range(1, 11):
#     if not i % 3 or not i % 5:
#         multiple1 += i
#     print('{0} / {1}'.format(i, multiple1))
    
# case 2
total = sum([x for x in range(1, 101) if not x % 3 or not x % 5])
print('{0}'.format(total))