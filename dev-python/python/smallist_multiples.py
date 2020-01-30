# 최소공배수 : 두 수를 곲해 최소공약수로 나누어 준다

# 최대공약수 case 1
def get_gcd(number1, number2):
    gcd = 0
    condition = number1 if number1 < number2 else number2

    for x in range(1, condition):
        if number1 % x == 0 and number2 % x == 0:
            gcd = x
    
    return gcd

# 최대공약수 case 2
import math
def get_gcd_from_math(number1, number2):
    return math.gcd(number1, number2)

if __name__ == "__main__":
    print(get_gcd(12, 18))
    print(get_gcd_from_math(27, 63))
