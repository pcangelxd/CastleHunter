import shutil

import image
import os


def clear_cache():
    try:
        os.remove("result.txt")
    except:
        pass

    files = os.listdir(f'cache/')
    for file in files:
        shutil.rmtree(f'cache/{file}', ignore_errors=True, onerror=None)

    files = os.listdir(f'cache/')
    for file in files:
        os.remove(f'cache/{file}')


def read_screen(purpose, count):
    os.mkdir(f"cache/{count}")
    image.convert_to_black_white(count)

    # Исходное изображение
    source_img = image.Cropyble(f"cache/black{count}.png")

    # Список всех найденных слов
    words = source_img.get_words()

    for ind, word in enumerate(words):
        if len(word) > 1:
            source_img.crop(word, f'cache/{count}/', f"output_{ind}.png")

    coordinate_image = source_img.coordinate_image
    coordinate = None

    files = os.listdir(f'cache/{count}/')
    find = False

    for file in files:
        if not find:
            s1 = image.image_to_string(f'cache/{count}/{file}')
            if len(s1) >= 5:
                chance = image.search_partial_text(s1, purpose)

                if chance >= 70:
                    find = True
                    coordinate = coordinate_image[file]

    if find:
        fl = open(f"result.txt", "w", encoding="UTF-8")
        fl.write(f"{count}\n{coordinate}")
        fl.close()