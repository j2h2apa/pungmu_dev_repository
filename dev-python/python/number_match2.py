import random as rd
import sys

# 숫자를 문자로 바꾸고 0으로 채움
answerlist = str(rd.randint(1,999)).zfill(3)

answerlist = "198"
print('{0}'.format(answerlist))

answer = list(answerlist)

while True:
    answer = list(answerlist)
    try:
        guess = input("숫자를 입력하세요.")
        if len(guess) != 3:
            raise RuntimeError('0~999 사이의 숫자만 입력해주세요.')
        if not guess.isdigit():
            print("숫자만 입력해주세요.")
            continue
    except RuntimeError as e:
        print(e)    
    except KeyboardInterrupt as e:
        print("\nexit process!")
        sys.exit()
    else:
        if guess == answer:
            print("exact answer!")
            break
        else:
            strike = 0
            ball = 0

            for idx in range(3):
                if answer[idx] == guess[idx]:
                    strike += 1
                    answer[idx] = 'a'
                elif guess[idx] in answer:
                    ball += 1
                    answer[answer.index(guess[idx])] = 'a'
                    
                print(answer)

            print("strike : %s ball : %s" % (strike, ball))