# Product Tracker

#### Запуск
В корне проекта создать файл app.env со следующим содержимым

    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=[имя пользователя]
    DB_PASSWORD=[пароль к БД]
    DB_NAME=[имя БД]
    HTTP_PORT=[порт сервера]
    
Далее запустить файл cmd/productTracker.go

Готово, сервер запущен и сервис начал свою работу!

#### Описание
Cервис, через который продавцы смогут передавать нам свои товары пачками в формате excel (xlsx).


Сервис принимает на вход ссылку на файл и id продавца, к чьему аккаунту будут привязаны загружаемые товары. Сервис читает файл и сохраняет, либо обновляет товары в БД. Обновление будет происходить, если пара (id продавца, offer_id) уже есть у нас в базе. В ответ на запрос выдаёт краткую статистику: количество созданных товаров, обновлённых, удалённых

Можно производить поиск по базе по seller_id, offer_id и имени позиции (по тексту "теле" находятся и "телефоны", и "телевизоры"). Ни один параметр не обязателен.

Также реализован сервер.
- /import - загрузка таблицы в базу
- /get - поиск позиций

#### Структура БД
- offer_id уникальный идентификатор товара в системе продавца
- name название товара
- price цена в рублях
- quantity количество товара на складе продавца
- available true/false, в случае false продавец хочет удалить товар из нашей базы
