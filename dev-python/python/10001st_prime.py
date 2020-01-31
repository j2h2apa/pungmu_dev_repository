# 10001 번째 소수 
from sympy import isprime

prime_count = 0; x = 0

while True:
    if isprime(x):
        prime_count += 1
    
    if prime_count == 10001:
        break

    x += 1

print("10001st prime numner : %d" % x)
