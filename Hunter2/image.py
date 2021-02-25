import os
import cv2
import numpy as np
import pytesseract
from PIL import Image, ImageGrab, ImageDraw
import win32gui
import win32ui
from ctypes import windll

import window

# Отступ для кропа
CROP_INDENTS = 5
pytesseract.pytesseract.tesseract_cmd = r'D:\Program\Tesseract-OCR\tesseract.exe'


def screen_window(win: window.Window):
    hwnd = win.hwnd
    w = win.game["w"]
    h = win.game["h"]

    hwndDC = win32gui.GetWindowDC(hwnd)
    mfcDC = win32ui.CreateDCFromHandle(hwndDC)
    saveDC = mfcDC.CreateCompatibleDC()

    saveBitMap = win32ui.CreateBitmap()
    saveBitMap.CreateCompatibleBitmap(mfcDC, w, h)

    saveDC.SelectObject(saveBitMap)

    # Change the line below depending on whether you want the whole window
    # or just the client area.
    # result = windll.user32.PrintWindow(hwnd, saveDC.GetSafeHdc(), win1)
    result = windll.user32.PrintWindow(hwnd, saveDC.GetSafeHdc(), 0)

    bmpinfo = saveBitMap.GetInfo()
    bmpstr = saveBitMap.GetBitmapBits(True)

    img = Image.frombuffer(
        'RGB',
        (bmpinfo['bmWidth'], bmpinfo['bmHeight']),
        bmpstr, 'raw', 'BGRX', 0, 1)

    win32gui.DeleteObject(saveBitMap.GetHandle())
    saveDC.DeleteDC()
    mfcDC.DeleteDC()
    win32gui.ReleaseDC(hwnd, hwndDC)

    if result == 1:
        img = img.crop((0, 75, w-50, h))
        return img


def convert_to_black_white(count: int):
    # Считываю картинку
    img = cv2.imread(f"cache/win{count}.png")
    # Перевожу в HSV
    img_hsv = cv2.cvtColor(img, cv2.COLOR_BGR2HSV)

    # Параметры HSV (подбираются в файле main.py)
    lower = np.array([0, 0, 0])
    upper = np.array([179, 200, 255])

    # Применяю HSV и сохраняю файл
    mask = cv2.inRange(img_hsv, lower, upper)
    cv2.imwrite(f"cache/black{count}.png", mask)


def image_to_string(path: str) -> str:
    image = Image.open(path)
    text = pytesseract.image_to_string(image, lang='eng', config=r'--oem 3 --psm 13').split("\n")
    return text[0]


def search_partial_text(src: str, dst: str) -> int:
    dst_buf = dst
    result = 0
    for char in src:
        if char in dst_buf:
            dst_buf = dst_buf.replace(char, '', 1)
            result += 1
    r1 = int(result / len(src) * 100)
    r2 = int(result / len(dst) * 100)
    return r1 if r1 < r2 else r2


# Этот класс я скопипастил с либы cropyble. В нем был баг.
class Cropyble:
    """Container for OCR and cropping methods."""

    def __init__(self, input_image):
        """Initializes a Cropyble object."""
        self.coordinate_image = {}
        self.input_image_path = os.path.join(os.getcwd(), input_image)
        self.box_data = {}
        self.height = 0
        self.width = 0
        self._image_to_data()

    def __repr__(self):
        """Returns a representation of a Cropyble object."""
        return f'<Cropyble image={self.input_image_path}>'

    def __str__(self):
        """Returns a verbose string representation of a Cropyble object."""
        string_representation = f'Cropyble Object for image: {self.input_image_path}\n'
        for key, value in self.box_data.items():
            string_representation += f'Word: {key} - Location: {value}\n'
        return string_representation

    def _image_to_data(self):
        """
        Utilizes pytesseract OCR to generate bounding box data for the image.
        Returns the bounding box data.
        """
        found_image = False
        while not found_image:
            try:
                input_image = Image.open(self.input_image_path)
                found_image = True
            except FileNotFoundError:
                raise FileNotFoundError(f'\nThe file [{self.input_image_path}] was not found.')

        word_box_data = pytesseract.image_to_data(input_image, lang='eng')
        char_box_data = pytesseract.image_to_boxes(input_image, lang='eng')
        self.width, self.height = input_image.size
        self._normalize_word_boxes(word_box_data)
        self._normalize_char_boxes(char_box_data)

    def _normalize_char_boxes(self, char_box_data):
        """
        Takes in bounding box data for characters from pytesseract.
        Stores the character and X,Y coordinates for its bounding box in self.box_data
        """
        char_box_data = char_box_data.split('\n')

        lines = [line.split(' ') for line in char_box_data]
        for line in lines[:-1]:
            self.box_data[line[0]] = [int(line[1]), (self.height - int(line[4])), int(line[3]),
                                      (self.height - int(line[2]))]

    def _normalize_word_boxes(self, word_box_data):
        """
        Takes in bounding box data for words from pytesseract.
        Stores the word and X,Y coordinates for its bounding box in self.box_data
        """
        word_box_data = word_box_data.split('\n')
        word_box_data = word_box_data[1:]

        lines = [line.split('\t') for line in word_box_data]
        for line in lines[:-1]:
            self.box_data[line[11]] = [int(line[6]), int(line[7]),
                                       (int(line[6]) + int(line[8])), (int(line[7]) + int(line[9]))]

    def crop(self, text_query, output_path, name_file):
        """
        Takes in a text query string.
        Outputs an image of the query from the input image if present.
        """
        original_image = Image.open(self.input_image_path)
        box = self.box_data[text_query]

        x = int((box[0] - CROP_INDENTS + box[2] + CROP_INDENTS) / 2)
        y = int((box[1] - CROP_INDENTS + box[3] + CROP_INDENTS) / 2)
        self.coordinate_image[name_file] = (x, y)

        new_image = original_image.crop((box[0] - CROP_INDENTS, box[1] - CROP_INDENTS, box[2] + CROP_INDENTS, box[3] + CROP_INDENTS))

        output_path = os.path.join(os.getcwd(), output_path + name_file)
        new_image.save(output_path)

    def get_box(self, word):
        """
        Takes in a string representing a word that was recognized.
        Returns a tuple representing the bounding box for the word in (x1, y1, x2, y2) format.
        Remember, for images the origin (0, 0) is located in the top-left corner of the image.
        """
        return tuple(self.box_data[word])

    def get_words(self):
        """Returns a list of recognized words."""
        return [word for word in self.box_data]
