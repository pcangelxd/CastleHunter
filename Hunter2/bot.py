import window
import image
import threading
import win32gui
import win32api
import pyautogui
import time

import screen


class Bot_1:
    def __init__(self, player: str, purpose, kor: int):
        self.win = window.Window(player)
        win32gui.MoveWindow(self.win.hwnd, 0, 0, int(win32api.GetSystemMetrics(0) / 2) + int(win32api.GetSystemMetrics(0) / 2 / 2),
                            int(win32api.GetSystemMetrics(1) / 2), True)

        self.purpose = purpose
        self.kor = kor
        self.x_start = 4
        self.y_start = 4
        self.x = self.x_start
        self.y = self.y_start
        self.x_max = 511 - 15
        self.y_max = 256

        pyautogui.FAILSAFE = True

    def start(self):
        self.open_find_win()
        self.change_kor()
        self.change_x(self.x_start)
        self.change_y(self.y_start)
        self.click_go()

    def find(self):
        if self.x >= self.x_max and self.y >= self.y_max:
            return True

        if self.x < self.x_max:
            pyautogui.moveTo(
                int(win32api.GetSystemMetrics(0) / 1.5),
                int(win32api.GetSystemMetrics(1) / 2 / 2))
            pyautogui.dragTo(
                int(win32api.GetSystemMetrics(0) / 2 / 2 - win32api.GetSystemMetrics(0) / 2 / 2 / 100 * 85),
                int(win32api.GetSystemMetrics(1) / 2 / 2), duration=0.4)
            self.x += 8
        else:
            pyautogui.moveTo(
                int(win32api.GetSystemMetrics(0) / 2 / 2 - win32api.GetSystemMetrics(0) / 2 / 2 / 100 * 85),
                int(win32api.GetSystemMetrics(1) / 2 / 2))
            pyautogui.dragTo(
                int(win32api.GetSystemMetrics(0) / 2 / 2 + win32api.GetSystemMetrics(0) / 2 / 2 / 100 * 90),
                int(win32api.GetSystemMetrics(1) / 2 / 2), duration=0.4)

            self.y += 10
            self.x = self.x_start

            self.open_find_win()
            self.change_x(self.x)
            self.change_y(self.y)
            self.click_go()

    def open_find_win(self):
        pyautogui.moveTo(int(win32api.GetSystemMetrics(0) / 2 / 1.25),
                         int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 82)
        pyautogui.click()

    def change_x(self, x):
        # x
        pyautogui.moveTo(int(win32api.GetSystemMetrics(0) / 2 / 1.4),
                         int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 60)
        pyautogui.click()

        for ch in str(x):
            if ch == "0":
                self.click_0()
            elif ch == "1":
                self.click_1()
            elif ch == "2":
                self.click_2()
            elif ch == "3":
                self.click_3()
            elif ch == "4":
                self.click_4()
            elif ch == "5":
                self.click_5()
            elif ch == "6":
                self.click_6()
            elif ch == "7":
                self.click_7()
            elif ch == "8":
                self.click_8()
            elif ch == "9":
                self.click_9()

        self.click_ok()

    def change_y(self, y):
        # y
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 2 / 1.15),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 60)
        pyautogui.click()

        for ch in str(y):
            if ch == "0":
                self.click_0()
            elif ch == "1":
                self.click_1()
            elif ch == "2":
                self.click_2()
            elif ch == "3":
                self.click_3()
            elif ch == "4":
                self.click_4()
            elif ch == "5":
                self.click_5()
            elif ch == "6":
                self.click_6()
            elif ch == "7":
                self.click_7()
            elif ch == "8":
                self.click_8()
            elif ch == "9":
                self.click_9()

        self.click_ok()

    def change_kor(self):
        # k
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 2 / 1.7),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 60)
        pyautogui.click()

        for ch in str(self.kor):
            if ch == "0":
                self.click_0()
            elif ch == "1":
                self.click_1()
            elif ch == "2":
                self.click_2()
            elif ch == "3":
                self.click_3()
            elif ch == "4":
                self.click_4()
            elif ch == "5":
                self.click_5()
            elif ch == "6":
                self.click_6()
            elif ch == "7":
                self.click_7()
            elif ch == "8":
                self.click_8()
            elif ch == "9":
                self.click_9()

        self.click_ok()

    def click_1(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 2 * 0.95),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 57)
        pyautogui.click()

    def click_2(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.95),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 57)
        pyautogui.click()

    def click_3(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.75),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 57)
        pyautogui.click()

    def click_4(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 2 * 0.95),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 49)
        pyautogui.click()

    def click_5(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.95),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 49)
        pyautogui.click()

    def click_6(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.75),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 49)
        pyautogui.click()

    def click_7(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 2 * 0.95),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 40)
        pyautogui.click()

    def click_8(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.95),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 40)
        pyautogui.click()

    def click_9(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.75),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 40)
        pyautogui.click()

    def click_0(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 2.1),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 32)
        pyautogui.click()

    def click_ok(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.8),
            int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 32)
        pyautogui.click()

    def click_go(self):
        pyautogui.moveTo(int(win32api.GetSystemMetrics(0) / 2 / 1.4),
                         int(win32api.GetSystemMetrics(1) / 2) - int(win32api.GetSystemMetrics(1) / 2) / 100 * 46)
        pyautogui.click()

    def screen(self, count):
        img = image.screen_window(self.win)
        img.save(f"cache/win{count}.png")

        my_thread = threading.Thread(target=screen.read_screen, args=(self.purpose, count,))
        my_thread.start()

class Bot_2:
    def __init__(self, player: str, purpose, kor: int):
        self.win = window.Window(player)
        win32gui.MoveWindow(self.win.hwnd, 0, int(win32api.GetSystemMetrics(1) / 2), int(win32api.GetSystemMetrics(0) / 2) + int(win32api.GetSystemMetrics(0) / 2 / 2),
                            int(win32api.GetSystemMetrics(1) / 2), True)

        self.purpose = purpose
        self.kor = kor
        self.x_start = 4
        self.y_start = 1018
        self.x = self.x_start
        self.y = self.y_start
        self.x_max = 511 - 15
        self.y_max = 765

        pyautogui.FAILSAFE = True

    def start(self):
        self.open_find_win()
        self.change_x(self.x_start)
        self.change_y(self.y_start)
        self.change_kor()
        self.click_go()

    def find(self):
        if self.x >= self.x_max and self.y <= self.y_max:
            return True

        if self.x < self.x_max:
            pyautogui.moveTo(
                int(win32api.GetSystemMetrics(0) / 1.5),
                int(win32api.GetSystemMetrics(1) / 2 * 1.5))
            pyautogui.dragTo(
                int(win32api.GetSystemMetrics(0) / 2 / 2 - win32api.GetSystemMetrics(0) / 2 / 2 / 100 * 85),
                int(win32api.GetSystemMetrics(1) / 2 * 1.5), duration=0.4)

            self.x += 8
        else:
            pyautogui.moveTo(
                int(win32api.GetSystemMetrics(0) / 2 / 2 - win32api.GetSystemMetrics(0) / 2 / 2 / 100 * 85),
                int(win32api.GetSystemMetrics(1) / 2 * 1.5))
            pyautogui.dragTo(
                int(win32api.GetSystemMetrics(0) / 2 / 2 + win32api.GetSystemMetrics(0) / 2 / 2 / 100 * 90),
                int(win32api.GetSystemMetrics(1) / 2 * 1.5), duration=0.4)

            self.y -= 10
            self.x = self.x_start

            self.open_find_win()
            self.change_x(self.x)
            self.change_y(self.y)
            self.click_go()

    def open_find_win(self):
        pyautogui.moveTo(int(win32api.GetSystemMetrics(0) / 2 / 1.25),
                         int(win32api.GetSystemMetrics(1) / 2 * 1.18))
        pyautogui.click()

    def change_x(self, x):
        # x
        pyautogui.moveTo(int(win32api.GetSystemMetrics(0) / 2 / 1.4),
                         int(win32api.GetSystemMetrics(1) / 2 * 1.4))
        pyautogui.click()

        for ch in str(x):
            if ch == "0":
                self.click_0()
            elif ch == "1":
                self.click_1()
            elif ch == "2":
                self.click_2()
            elif ch == "3":
                self.click_3()
            elif ch == "4":
                self.click_4()
            elif ch == "5":
                self.click_5()
            elif ch == "6":
                self.click_6()
            elif ch == "7":
                self.click_7()
            elif ch == "8":
                self.click_8()
            elif ch == "9":
                self.click_9()

        self.click_ok()

    def change_y(self, y):
        # y
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 2 / 1.15),
                         int(win32api.GetSystemMetrics(1) / 2 * 1.4))
        pyautogui.click()

        for ch in str(y):
            if ch == "0":
                self.click_0()
            elif ch == "1":
                self.click_1()
            elif ch == "2":
                self.click_2()
            elif ch == "3":
                self.click_3()
            elif ch == "4":
                self.click_4()
            elif ch == "5":
                self.click_5()
            elif ch == "6":
                self.click_6()
            elif ch == "7":
                self.click_7()
            elif ch == "8":
                self.click_8()
            elif ch == "9":
                self.click_9()

        self.click_ok()

    def change_kor(self):
        # k
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 2 / 1.7),
                         int(win32api.GetSystemMetrics(1) / 2 * 1.4))
        pyautogui.click()

        for ch in str(self.kor):
            if ch == "0":
                self.click_0()
            elif ch == "1":
                self.click_1()
            elif ch == "2":
                self.click_2()
            elif ch == "3":
                self.click_3()
            elif ch == "4":
                self.click_4()
            elif ch == "5":
                self.click_5()
            elif ch == "6":
                self.click_6()
            elif ch == "7":
                self.click_7()
            elif ch == "8":
                self.click_8()
            elif ch == "9":
                self.click_9()

        self.click_ok()

    def click_1(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 2 * 0.95),
            int(win32api.GetSystemMetrics(1) / 2 * 1.44))
        pyautogui.click()

    def click_2(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.95),
            int(win32api.GetSystemMetrics(1) / 2 * 1.44))
        pyautogui.click()

    def click_3(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.75),
            int(win32api.GetSystemMetrics(1) / 2 * 1.44))
        pyautogui.click()

    def click_4(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 2 * 0.95),
            int(win32api.GetSystemMetrics(1) / 2 * 1.52))
        pyautogui.click()

    def click_5(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.95),
            int(win32api.GetSystemMetrics(1) / 2 * 1.52))
        pyautogui.click()

    def click_6(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.75),
            int(win32api.GetSystemMetrics(1) / 2 * 1.52))
        pyautogui.click()

    def click_7(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 2 * 0.95),
            int(win32api.GetSystemMetrics(1) / 2 * 1.6))
        pyautogui.click()

    def click_8(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.95),
            int(win32api.GetSystemMetrics(1) / 2 * 1.6))
        pyautogui.click()

    def click_9(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.75),
            int(win32api.GetSystemMetrics(1) / 2 * 1.6))
        pyautogui.click()

    def click_0(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 2.1),
            int(win32api.GetSystemMetrics(1) / 2 * 1.68))
        pyautogui.click()

    def click_ok(self):
        pyautogui.moveTo(int(
            win32api.GetSystemMetrics(0) / 1.8),
            int(win32api.GetSystemMetrics(1) / 2 * 1.68))
        pyautogui.click()

    def click_go(self):
        pyautogui.moveTo(int(win32api.GetSystemMetrics(0) / 2 / 1.4),
                         int(win32api.GetSystemMetrics(1) / 2 * 1.54))
        pyautogui.click()

    def screen(self, count):
        img = image.screen_window(self.win)
        img.save(f"cache/win{count}.png")

        my_thread = threading.Thread(target=screen.read_screen, args=(self.purpose, count,))
        my_thread.start()