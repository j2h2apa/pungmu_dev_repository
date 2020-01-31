from functools import reduce

with open("./lagest_product_in_a_series.dat", "rt", encoding="utf8") as f:
    lines = f.readlines()

# with end

src = ""
for line in lines:
    src += line.rstrip()

# 공백 및 개행없이 문자열 합치는 다른 방법
# src = ''.join([line.rstrip() for line in lines])

f.close()

# case 1
# for x in range(0, 1000):
#     product = 1
#     for y in src[x:x+13]:
#         product *= int(y)

# for end

# case 2
max_value = 0
for i in range(0, 1000):
    product = reduce(lambda x, y : int(x)*int(y), src[i:i+13])
    max_value = max_value if max_value > int(product) else int(product)

    print(src[i:i+13], product)

print("max_value : ", max_value)