#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Использование: $0 <имя_файла.go>"
    exit 1
fi

go_file="$1"

task_number=$(echo "$go_file" | grep -o -E '[0-9]+')

if [ -z "$task_number" ]; then
    echo "Не удалось извлечь номер задачи из имени файла."
    exit 1
fi

# папка с тестами
test_path="tests/$task_number/*.a"

# считаем файлы
total_tests=$(ls $test_path | wc -l)
current_test=0

all_tests_passed=true
failed_tests=()

echo "Начало тестирования..."

# перебор тестов
for file in $test_path; do
    if [[ -f "$file" ]]; then
        ((current_test++))
        input_file="${file%.a}"
        if ! diff -qZ <(go run "$go_file" < "$input_file") "$file" > /dev/null; then
            failed_tests+=("go run $go_file < $input_file")
            all_tests_passed=false
        fi
        # progress bar
        printf "\rПрогресс: [%-50s] %d%%" $(printf "%-${current_test}s" | tr ' ' '#') $((current_test * 100 / total_tests))
    fi
done

echo 

if $all_tests_passed; then
    echo "Все тесты пройдены успешно."
else
    echo "Тесты, которые не пройдены:"
    for test in "${failed_tests[@]}"; do
        echo "$test"
    done
fi
