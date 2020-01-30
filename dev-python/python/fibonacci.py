# Even Fibonacci numbers
index = 1; previousTerm = 1; nextTerm = 0
totalEvenTerm = 0

while index <= 4000000:
    totalEvenTerm += index if index % 2 == 0 else 0

    nextTerm = previousTerm + index
    previousTerm = index
    index = nextTerm

print("Fibonacci sum even_valued : ", totalEvenTerm)
    
