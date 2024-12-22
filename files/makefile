# Переменные
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html

# Цель по умолчанию
all: test coverage open-coverage benchmark

# Тестирование и сбор данных о покрытии
test:
	go test -coverprofile=$(COVERAGE_FILE)

# Генерация HTML отчета о покрытии
coverage:
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)

# Открытие отчета в браузере
open-coverage:
	xdg-open $(COVERAGE_HTML) || open $(COVERAGE_HTML)

# Цель для запуска бенчмарков
benchmark:
	go test -bench .

# Цель для тестирования, генерации отчета и открытия HTML отчета
run: test coverage open-coverage benchmark

# Чистка сгенерированных файлов
clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)
