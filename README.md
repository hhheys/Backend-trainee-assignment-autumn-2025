# Backend Trainee Assignment Autumn 2025

Репозиторий содержит решение задания для стажёра по backend на осень 2025 года.
Проект написан на Go и использует Docker и Docker Compose для упрощённого развёртывания.
API специфицировано через OpenAPI, линтинг и форматирование кода обеспечивается `golangci-lint`.

## Структура проекта

```
├── cmd/server/            # Точка входа приложения
├── internal/              # Бизнес-логика и вспомогательные пакеты
├── .env.example           # Пример конфигурации окружения
├── Dockerfile             # Сборка Docker-образа приложения
├── docker-compose.yaml    # Сборка и запуск контейнеров
├── .golangci.yaml         # Конфигурация линтера golangci-lint
├── go.mod / go.sum        # Модули Go
├── openapi.yml            # Спецификация API
└── task.md                # Описание задания
```

## Технологии

* Язык: Go
* API: OpenAPI (`openapi.yml`)
* Контейнеризация: Docker, Docker Compose
* Линтер: GolangCI-Lint
* Миграции: Golang-Migrate
* Драйвер базы данных: database/sql

## Что было сделано
* Реализованы основные эндпоинты
* Валидация входных данных
* Кастомные ошибки
* Доступ к защищенным эндпоинтам через Authorization: Bearer и Access Token, который содержится в .env файле
* Настроен линтер
* Настроено логирование
* Написаны тесты для работы с базой данных
* Сгенерированы seed данные для упрощения тестирования

## Быстрый старт

1. Клонируйте репозиторий:

```bash
git clone https://github.com/hhheys/Backend-trainee-assignment-autumn-2025.git
cd Backend-trainee-assignment-autumn-2025
```

2. Создайте файл окружения:

```bash
cp .env.example .env
# при необходимости отредактируйте переменные среды
```

3. Сборка и запуск через Docker Compose:

```bash
docker-compose up --build
```

4. Сервер будет доступен по адресу, указанному в `.env` (по умолчанию `http://localhost:8080`).

5. Документация API доступна через файл `openapi.yml`.
