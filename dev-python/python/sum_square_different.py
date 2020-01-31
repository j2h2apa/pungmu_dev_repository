# 1 ~ 100 제곱 후 sum 과 sum 이후 제곱의 차이
# case 1
total_sum_square = 0
total_square_sum = 0

for x in range(1, 11):
    total_square_sum += pow(x, 2)
    total_sum_square += x

print("total_square_sum : %d / total_sum_square : %d" \
    % (total_square_sum, total_sum_square ** 2))
print("sum square different : {0}".format(total_sum_square**2 - total_square_sum))

# case 2
import sympy
 
i, n = sympy.symbols('i n')
sum_of_squares = sympy.Sum(i**2, (i, 1, n))
square_of_sums = sympy.Sum(i, (i, 1, n))**2
difference = sympy.lambdify(n, square_of_sums - sum_of_squares)
 
print(difference(100))
