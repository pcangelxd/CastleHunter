import ssl
import datetime
from email import encoders
from email.mime.base import MIMEBase
from email.mime.multipart import MIMEMultipart  # Многокомпонентный объект
from email.mime.text import MIMEText  # Текст/HTML
from email.mime.image import MIMEImage  # Изображения
import smtplib
import time
import os
import requests

import bot
import screen
import config


def main():
    while True:
        response = requests.get(config.url + '/castle/found/')
        content = response.json()
        if content['status']:
            requests.post(f"http://{config.ip}:3333/bot/find/", params={'nick': content['purpose'],
                                                                      "kor": content["kor"],
                                                                      "email": content["email"]})
            start_find(content['purpose'], content["kor"], content["email"])
        time.sleep(20)


def start_find(purpose, kor, email):
    screen.clear_cache()
    count = 0

    bot1 = bot.Bot_1(config.win1, purpose, kor)
    bot2 = bot.Bot_2(config.win2, purpose, kor)
    bot3 = bot.Bot_3(config.win3, purpose, kor)
    bot4 = bot.Bot_4(config.win4, purpose, kor)

    bot1.start()
    bot2.start()
    bot3.start()
    bot4.start()
    time.sleep(5)

    while True:
        print("Поиск")
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

        files = os.listdir(config.path_hunter2)
        if "result.txt" in files:
            break

        bot1.screen(count)
        count += 1
        res1 = bot1.find()

        bot2.screen(count)
        count += 1
        res2 = bot2.find()

        bot3.screen(count)
        count += 1
        res3 = bot3.find()

        bot4.screen(count)
        count += 1
        res4 = bot4.find()

        if res1 and res2 and res3 and res4:
            subject = config.subject
            body = config.body_not_found
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
            text = message.as_string()

            # Подключение к серверу при помощи безопасного контекста и отправка письма
            context = ssl.create_default_context()
            with smtplib.SMTP_SSL(config.host, config.port, context=context) as server:
                server.login(sender_email, password)
                server.sendmail(sender_email, receiver_email, text)
            break


    screen.clear_cache()


if __name__ == '__main__':
    main()