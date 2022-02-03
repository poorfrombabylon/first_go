### Выбор хранилища
Чтобы запустить программу с in-memory, нужно в docker-compose.yaml поменять под строкой command "out" на "in", и чтобы
запустить с postgres "in" на "out" :)

### Запуск
1. docker compose build
2. docker compose up

### Left to do
1. Добавить unit tests (handler_test.go, inmemory_test.go, postgres_test.go)
2. Добавление параметров out/in с клавиатуры при запуске

### Использование
Заходите на localhost:8080