## Онлайн Библиотека Песен
# Описание:
Данный проект реализован на языке Go (Golang) и предоставляет онлайн библиотеку песен.
Приложение подключается к API Last.fm для получения дополнительной информации о песнях и исполнителях.

-**Запуск проекта**:

Для запуска проекта необходимо выполнить следующие шаги:

Сборка и запуск Docker-контейнера:
        
    docker-compose up --build

Эта команда соберет проект и запустит приложение вместе с базой данных на порту 4444.

Запуск Swagger-интерфейса:

      go run cmd/swagger/main.go

Эта команда запустит Swagger-приложение на порту 8080. Вы также можете перейти к файлу main.go в директории cmd/swagger/ и запустить его из интерфейса вашей IDE.


## Функциональность:

* Добавление песни: 
При добавлении новой песни необходимо указать название группы и название песни на латинице. Допускается ввод только латинских символов.


* Обновление записи:
Имеется функция обновления полей в записи о песне. Вы можете изменить информацию о песне, предоставив необходимые данные.


* Удаление песни
Вы можете удалить запись о песне из библиотеки, используя функцию удаления по идентификатору песни.


* Просмотр текста песни с пагинацией
Имеется возможность вывести текст песни с поддержкой пагинации по страницам и указанием размера страницы.


* Вывод всех песен с сортировкой и фильтрацией
Вы можете получить список всех песен с возможностью сортировки по любому полю. Для этого необходимо указать название поля и установить его значение в true, а также параметр isAscending=true для сортировки по возрастанию. Аналогично можно установить false для сортировки по убыванию.


## Логирование
Весь проект реализует подробное логирование для удобства разработки и отладки. Логи включают информацию о выполнении операций и возможных ошибках.


