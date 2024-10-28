# mqtt_lab
Лабораторная работа по IoT. Исследование работы протокола MQTT.

## Цель работы
Получить опыт работы с системой MQTT, научиться находить уязвимости в системе и испарвлять их. \
Рекомендуется выполнять работу на ОС Linux.

## Ход работы
Даны исходники клиента и сервера MQTT, написанные на Go. \
Для запуска системы использовать docker. \
Для успешного выполнения работы неоходимо:
1. Изучить технологию Docker (Dockerfile, Docker Compose)
2. Изучить протокол MQTT
3. Запустить тестовый стенд (см. _Полезные ссылки и команды_ )
4. Снять дампы трафика работы протокола MQTT
5. Исследовать дампы, найти уязвимости работы клиент-сервера
6. Внести измененния в исходный код и устранить уязвимости
7. Перезапустить систему, снять дампы трафика, убедиться в безопасности работы протокола
8. Отразить ход работы в отчете


## Требования к отчету
1. Схема работы протокола MQTT
2. Назначение и принцип работы брокера сообщений, используемого в лабораторной
3. Список найденных уязвимостей клиент-сервера
4. Решения по устранению уязвимостей
5. Показать на дампе трафика, что уязвимости были исправлены

## Полезные ссылки и команды
- Установка Docker в Ubuntu: https://docs.docker.com/engine/install/ubuntu/
- Запуск проекта: `docker compose -f docker/docker-compose.yml up --build`
- Для работы с дампами трафика использовать `Wireshark`

---
Go 1.23.2
