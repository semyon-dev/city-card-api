## Реализованная функциональность
* Профили (регистрация и авторизация)
* Карты (создание счета, пополнение и оплата услуг, просмотр баланса)

## Особенность проекта в следующем
* Кроссплатформенное приложение
* Удобное моментальное создание виртуального счета

## Основной стек технологий бэкенда

* Golang
* MongoDB
* JWT
* redis

## Демо

Демо доступно [по адресу](https://vk.com/)

Реквизиты тестового пользователя: email: admin@test.ru, пароль: testuser

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
```

# Разработчики

* [Дмитрий Чудинов](https://t.me/dchudik) - backend
* [Новиков Семен](https://t.me/semyon_dev) - backend
* [Эльдар Курманалиев](https://t.me/elik_sir) - web app









