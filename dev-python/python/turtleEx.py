from turtle import Turtle, Screen
import os

screen = Screen()
turtle = Turtle()

turtle.forward(100)
turtle.left(90)
turtle.shape('turtle')
turtle.forward(50)
turtle.left(90)      #왼쪽으로 회전 90도
turtle.fd(100)       #forward 와 같은 명령어
turtle.speed('fast')  #속도 조절 '빠르게'
turtle.fd(100)    
turtle.speed(1)  # 속도를 숫자로 표시. 0~10까지 숫자.
turtle.fd(50)

# os.system('Pause')
screen.mainloop()

