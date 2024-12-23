# Анализатор логов

## Проект III модуля Академии Бэкенда 2024 Т-Образования

Программа-анализатор логов.

На вход программе через аргументы командной строки задаётся:
* путь к одному или нескольким NGINX лог-файлам в виде локального шаблона или URL
* необязательные временные параметры from и to в формате ISO8601
* необязательный параметр формата вывода результата: markdown или adoc
* необязательные параметры filter-field и filter-value для фильтрации логов по значению поля
* необязательный параметр highest, определяющий количество строк в таблицах метрик отчёта  
* необязательный параметр read, указывающий на количество строк, которое нужно прочитать из каждого файла

Программа, анализируя логи:
* Подсчитывает общее количество запросов
* Определяет наиболее часто запрашиваемые ресурсы
* Определяет наиболее часто встречающиеся коды ответа
* Определяет наиболее часто встречающиеся IP-адреса клиентов
* Определяет наиболее часто встречающиеся HTTP-заголовки User-Agent
* Рассчитывает средний размер ответа сервера
* Рассчитывает 95% перцентиль размера ответа сервера

В результате работы программы получается файл с отчётом в соответствующем формате.
