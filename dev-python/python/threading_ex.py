import threading

class Messanger(threading.Thread):
    def __init__(self, name=""):
        """__init__ 메소드 안에서 threading.Thread를 init한다"""
        threading.Thread.__init__(self)
        pass

    def run(self):
        for _ in range(10):
            print(threading.currentThread().getName())


if __name__ == "__main__":
    x = Messanger(name = "메세지를 송신합니다.")
    y = Messanger(name = "메세지를 수신합니다.")

    x.start()
    y.start()
    pass   
