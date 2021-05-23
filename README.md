## Реализованная функциональность (подробно в презентации)
* Профили (регистрация и авторизация)
* Карты (создание счета, пополнение и оплата услуг, просмотр баланса)

## Особенность проекта в следующем (подробно в презентации)
* Кроссплатформенное приложение
* Удобное моментальное создание виртуального счета

## Основной стек технологий бэкенда

* Golang
* MongoDB
* JWT
* redis

## Демо

* [Демо веб приложения](http://citi-card.dchudinov.ru)

* Демо БЭКЕНДА доступно [по адресу](http://citi-card.dchudinov.ru:8081/api/v1)

Реквизиты тестового пользователя не нужны, есть специальные кнопки

## Установка зависимостей проекта

`go mod download`

## Среда запуска

Linux \
` sudo apt install golang` \
`source .env && go run cmd/main.go`

.env файл
```
export MONGO_URI="mongodb+srv://admin:****"
export HTTP_PORT=8081
export REDIS_URI=""
```

## Запуск в докере

```
docker build -t city-card-api .
docker run -p 8081:8081 -e MONGO_URI="..." -e HTTP_PORT=8081 -e REDIS_URI=""
```

# Разработчики

* [Дмитрий Чудинов](https://t.me/dchudik) - backend
* [Новиков Семен](https://t.me/semyon_dev) - backend
* [Эльдар Курманалиев](https://t.me/elik_sir) - web app









