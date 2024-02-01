#!/bin/bash

# Проверяем, передан ли аргумент скрипту
if [ "$#" -ne 1 ]; then
    echo "Использование: $0 <имя_файла.go>"
    exit 1
fi

# Имя файла Go, переданное в качестве аргумента
go_file="$1"

# Извлекаем цифру из имени файла с помощью регулярного выражения
task_number=$(echo "$go_file" | grep -o -E '[0-9]+')

# Проверяем, удалось ли извлечь номер
if [ -z "$task_number" ]; then
    echo "Не удалось извлечь номер задачи из имени файла."
    exit 1
fi

# Формируем путь к папке с тестами
test_path="tests/$task_number/*.a"

# Считаем количество файлов .a
total_tests=$(ls $test_path | wc -l)
current_test=0

# Переменная для отслеживания статуса тестов
all_tests_passed=true

echo "Начало тестирования..."

# Перебираем все файлы .a в папке с тестами
for file in $test_path; do
    if [[ -f "$file" ]]; then
        ((current_test++))
        # Убираем расширение .a и добавляем входной файл для теста
        input_file="${file%.a}"
        if ! diff -qZ <(go run "$go_file" < "$input_file") "$file" > /dev/null; then
            echo -e "\nТест не пройден: файлы $input_file и $file не совпадают"
            all_tests_passed=false
        fi
        # Обновляем прогресс-бар
        printf "\rПрогресс: [%-50s] %d%%" $(printf "%-${current_test}s" | tr ' ' '#') $((current_test * 100 / total_tests))
    fi
done

echo # Добавляем новую строку после прогресс-бара

# Выводим сообщение о статусе прохождения тестов
if $all_tests_passed; then
    echo "Все тесты пройдены успешно."
else
    echo "Обнаружены ошибки в тестах."
fi
