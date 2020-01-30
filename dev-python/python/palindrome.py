# 팰린드롬 : 좌우 대칭 숫자 혹은 문자가 같은 것
# 3 * 3 숫자 중 팰린드롬 찾기

import time

# case 1
# def isPalindrome(string):

#     if string == string[::-1]:
#         return True
    
#     return False

# if __name__ == "__main__":
#     max_value = 0
#     for i in range(900, 1000):
#         for j in range(900, 1000):
#             product = i * j
#             if isPalindrome(str(product)):
#                 max_value = product if product > max_value else max_value

#     print(max_value)

# case 2 
# itertools 사용으로 중복된 숫자는 포함하지 않음 (순열)
import itertools as itertool

start = time.time()

max_value = 0
for n1, n2 in itertool.combinations(range(999, 0, -1), 2):
    product = n1 * n2
    if (str(product) == str(product)[::-1]) and product > max_value:
        max_value = product

print(max_value)

print("running time : {0}".format(time.time() - start))