# CRUD Project on Golang

Этот проект представляет собой простое CRUD-приложение на языке программирования Go, с взаимодействием с базой данных PostgreSQL. Проект разработан с целью обучения работе с базами данных и реализации базовых операций CRUD (создание, чтение, обновление и удаление).

## Описание проекта

Основная модель приложения — это сущность Task (задача). Задача содержит следующие поля:
- Название: краткое название задачи.
- Описание: подробное описание задачи.
- Статус: текущий статус выполнения задачи (например, "в процессе", "завершено").
- Дата и время создания.
- Дата и время последнего обновления.

## Архитектура проекта

Проект разделен на несколько основных модулей:

### 1. Модели

Модель задачи Task описана в [internal/models/task.go](https://github.com/notblinkyet/Crud/blob/master/internal/models/task.go). Это Go-структура, представляющая сущность задачи и ее свойства.

### 2. Интерфейс работы с базой данных

Интерфейс для взаимодействия с базой данных описан в [internal/storage/storage.go](https://github.com/notblinkyet/Crud/blob/master/internal/storage/storage.go). Он включает следующие методы:
- Create — создание новой задачи.
- ReadId — получение задачи по её ID.
- ReadAll — получение списка всех задач.
- Update — обновление полей задачи (название, описание, статус).
- Delete — удаление задачи по ID.

В будущем планируется поддержка других баз данных (MySQL, MongoDB), для этого необходимо реализовать указанный интерфейс.

### 3. Работа с PostgreSQL

Реализация методов работы с PostgreSQL находится в [internal/storage/posgresql](https://github.com/notblinkyet/Crud/tree/master/internal/storage/posgresql). Для взаимодействия с базой данных используется пакет pgx.

### 4. Хэндлеры HTTP-запросов

Обработка HTTP-запросов осуществляется в [internal/transport/handlers](https://github.com/notblinkyet/Crud/tree/master/internal/transport/handlers). Хэндлеры вызывают методы, описанные в интерфейсе работы с базой данных. Для реализации используется стандартный пакет net/http.

### 5. Сервис и запуск сервера

Структура сервера, его конструктор и методы для настройки маршрутов и запуска сервера описаны в [internal/services](https://github.com/notblinkyet/Crud/tree/master/internal/services). Для работы используется стандартный пакет net/http.

### 6. Конфигурация

Конфигурация приложения задается в файле config.yaml и описана в [internal/config/config.go](https://github.com/notblinkyet/Crud/blob/master/internal/config/config.go). Для работы с конфигурацией используется пакет gopkg.in/yaml.v3.

### 7. Точка входа в приложение

Главная функция приложения находится в [cmd/main.go](https://github.com/notblinkyet/Crud/blob/master/cmd/main.go). Она выполняет следующие действия:
1. Считывает конфигурацию из config.yaml.
2. Подключается к базе данных, указанной в конфигурации.
3. Настраивает параметры сервера и маршруты.
4. Запускает HTTP-сервер.

## Зависимости
- Go (https://golang.org/) 1.16+
- pgx (https://github.com/jackc/pgx) для взаимодействия с PostgreSQL
- gopkg.in/yaml.v3 для работы с YAML

## Планы на будущее
- Реализация поддержки других баз данных, таких как MySQL и MongoDB.
- Улучшение обработки ошибок и логирования.
- Добавление аутентификации и авторизации.