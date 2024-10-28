# Загрузить изображение Go для сборки бинарного файла
FROM golang:1.19 AS build

# Установить рабочую директорию
WORKDIR /app

# Скопировать все файлы в рабочую директорию
COPY . .

# Собрать бинарный файл
RUN go build -o main .

# Создать минимальное изображение from scratch
FROM scratch

# Скопировать бинарный файл из стадии сборки
COPY --from=build /app/main /main

# Указать команду для запуска
ENTRYPOINT ["/main"]
