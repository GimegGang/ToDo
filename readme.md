# ToDo Application

Это простое веб-приложение ToDo, написанное на Go. Приложение использует библиотеку chi для маршрутизации и SQLite для хранения данных.

## Требования

- Go 1.16 или выше
- SQLite

## Установка

1. Склонируйте репозиторий:

   ```sh
   git clone https://github.com/GimegGang/ToDo.git
    ```
   
2. Перейдите в директорию проекта:

    ```sh
   cd ToDo
    ```
   
3. Установите необходимые зависимости:

    ```sh
   go mod tidy
    ```

4. Создайте файл конфигурации:

    ```yaml
    env: "local"
    address: ":8080"
    timeout: 5s
    idle_timeout: 60s
    storage_path: "storage/storage.db"
    ```

## Запуск

1. Запустите приложение:

    ```sh
      go run src/ToDo/main.go
    ```

## Структура проекта

* `src/ToDo/main.go`: основной файл приложения.
* `internal/config`: пакет для загрузки конфигурации.
* `internal/handlers`: обработчики HTTP-запросов.
* `internal/logger`: пакет для логирования.
* `internal/storage/sqlite`: пакет для работы с базой данных SQLite.
* `templates/`: HTML-шаблоны.
* `static/`: статические файлы (CSS, JavaScript и т.д.).
