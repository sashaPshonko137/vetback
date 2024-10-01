#!/bin/sh

# Небольшая задержка для обеспечения того, что база данных уже запущена
sleep 2

# Выполнение миграций через Goose
goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v
