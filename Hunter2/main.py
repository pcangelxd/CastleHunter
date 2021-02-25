import datetime
import ssl
import threading
from email import encoders
from email.mime.base import MIMEBase
from email.mime.multipart import MIMEMultipart  # Многокомпонентный объект
from email.mime.text import MIMEText  # Текст/HTML
from email.mime.image import MIMEImage  # Изображения
import smtplib
import time
import os
from fastapi import FastAPI
import uvicorn

import bot
import screen
import config

app = FastAPI()


@app.post("/bot/find/")
def bot_find(nick: str, kor: str, email: str):
    my_thread = threading.Thread(target=start_find, args=(nick, kor, email, ))
    my_thread.start()


def main():
    uvicorn.run(app, host=config.ip, port=3333)


def start_find(purpose, kor, email):
    screen.clear_cache()
    count = 0

    bot1 = bot.Bot_1(config.win1, purpose, kor)
    bot2 = bot.Bot_2(config.win2, purpose, kor)

    bot1.start()
    bot2.start()
    time.sleep(5)

    while True:
        files = os.listdir(os.getcwd())

        if "result.txt" in files:
            with open("result.txt", "r") as file:
                number = file.readline()

            subject = config.subject
            body = config.body
            sender_email = config.sender_email
            receiver_email = email
            password = config.password

            # Создание составного сообщения и установка заголовка
            message = MIMEMultipart()
            message["From"] = sender_email
            message["To"] = receiver_email
            message["Subject"] = subject
            message["Bcc"] = receiver_email  # Если у вас несколько получателей

            # Внесение тела письма
            message.attach(MIMEText(body, "plain"))

            filename = f"cache/win{number[:-1]}.png"  # В той же папке что и код

            # Открытие PDF файла в бинарном режиме
            with open(filename, "rb") as attachment:
                # Заголовок письма application/octet-stream
                # Почтовый клиент обычно может загрузить это автоматически в виде вложения
                part = MIMEBase("application", "octet-stream")
                part.set_payload(attachment.read())

            # Шифровка файла под ASCII символы для отправки по почте
            encoders.encode_base64(part)

            # Внесение заголовка в виде пара/ключ к части вложения
            part.add_header(
                "Content-Disposition",
                f"attachment; filename= {filename}",
            )

            # Внесение вложения в сообщение и конвертация сообщения в строку
            message.attach(part)
            text = message.as_string()

            # Подключение к серверу при помощи безопасного контекста и отправка письма
            context = ssl.create_default_context()
            with smtplib.SMTP_SSL(config.host, config.port, context=context) as server:
                server.login(sender_email, password)
                server.sendmail(sender_email, receiver_email, text)

            break

        files = os.listdir(config.path_hunter)
        if "result.txt" in files:
            break

        bot1.screen(count)
        count += 1
        res1 = bot1.find()

        bot2.screen(count)
        count += 1
        res2 = bot2.find()

        if res1 and res2:
            break

    screen.clear_cache()


if __name__ == '__main__':
    main()