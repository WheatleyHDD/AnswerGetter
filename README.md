# AnswerGetter
## О программе
Небольшая утилита для разработчиков чат-ботов (преимущественно ВКонтакте), позволяющая быстро получить ответ на вопрос.
**Используются базы iHa Bot**
## Как пользоваться?
1. Установить Go >= 1.12
2. Собрать программу: `go build`
В папке должен появиться исполняемый файл AnswerGetter (на Windows: AnswerGetter.exe). Его нужно запустить. 
Запуститься веб-сервер на порте 8080.
Чтобы начать работать достаточно отправить запрос такого вида: `http://<ВАШ_ДОМЕН>/getMessage/<СООБЩЕНИЕ>`
Например: `http://localhost:8080/getMessage/привет`
Сервер должен вернуть:
`{"answer": "<ОТВЕТ>", "attachment": "<ВЛОЖЕНИЯ>"}`
