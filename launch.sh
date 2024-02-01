#!/bin/bash

for file in tests/1/*; do
    if [[ -f "${file%}.a" ]]; then
        if diff -q <(go run task_1.go < "$file") "${file%}.a"; then
            echo "Файлы $file и ${file%.*}a совпадают"
        else
            echo "Файлы $file и ${file%.*}a не совпадают"
        fi
    fi
done
