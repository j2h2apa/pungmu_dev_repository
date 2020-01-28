
import random as rd

answer = rd.randint(1, 40)

print("{0}".format(answer))

while True:
    try:
        guess = int(input("숫자를 입력하세요."))
        print("answer type is %s / guess type is %s" % (type(answer), type(guess)))

        if answer == guess:
            print("exact number!")
            break
        elif guess > answer:
            print("더 작은 수 입니다.")
        else:
            print("더 큰 수 입니다.")
    except:
        print("1-40 사이의 숫자를 입력하세요")