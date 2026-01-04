package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"go-sprint4-tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

// Разбирает входящую строку данных (разделитель запятая) и возвращает:
// количество шагов,
// продолжительность прогулки
// и ошибку, если она возникла
func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	arData := strings.Split(data, ",")

	if len(arData) != 2 {
		return 0, 0, fmt.Errorf("Ошибка в формате входных данных")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(arData[0]))

	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка в количестве шагов: %v", err)
	}

	if steps < 1 {
		return 0, 0, fmt.Errorf("Количество шагов должно быть больше нуля")
	}

	durationMinutes, err := time.ParseDuration(arData[1])

	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка в продолжительности прогулки: %v", err)
	}

	if !(durationMinutes > 0) {
		return 0, 0, fmt.Errorf("Продолжительность прогулки отрицательная: %v", durationMinutes)
	}

	return steps, durationMinutes, nil
}

// Разбирает входящую строку с шагами и длительностью прогулки, плюс дополнительно вес и рост пользователя
// Возвращает информацию о проделанной прогулке:
// количество шагов,
// дистанция в километрах,
// количество сожжённых калорий
func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	steps, durationMinutes, err := parsePackage(data)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	if !(steps > 0) {
		return ""
	}

	//дистанция в метрах колчество шагов * длина шага
	var distance float64 = float64(steps) * stepLength

	//переводим дистанцию в километры
	distance /= mInKm

	//рассчитываем количество калорий, потраченных на прогулку
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, durationMinutes)

	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %v.\nДистанция составила %v км.\nВы сожгли %v ккал.", steps, distance, calories)
}
