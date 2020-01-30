# 소인수 분해 (나누어 1이 남는 수)
import time

def prime_factor(number):
    factor = 2

    while number != 1:
        if number % factor == 0:
            print("number : %d / factor : %d" % (number, factor))
            number /= factor
        else:
            factor += 1
    
    return factor

if __name__ == "__main__":
    start_job = time.time()

    number = 100
    print(prime_factor(number))

    print("작업 시간 : ", time.time() - start_job)
