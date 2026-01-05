package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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

	//steps, err := strconv.Atoi(strings.TrimSpace(arData[0]))
	steps, err := strconv.Atoi(arData[0])

	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка в количестве шагов: %v", err)
	}

	if steps < 1 {
		return 0, 0, fmt.Errorf("Количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(arData[1])

	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка в продолжительности прогулки: %v", err)
	}

	if !(duration > 0) {
		return 0, 0, fmt.Errorf("Продолжительность прогулки отрицательная: %v", duration)
	}

	return steps, duration, nil
}

// Разбирает входящую строку с шагами и длительностью прогулки, плюс дополнительно вес и рост пользователя
// Возвращает информацию о проделанной прогулке:
// количество шагов,
// дистанция в километрах,
// количество сожжённых калорий
func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	steps, duration, err := parsePackage(data)

	if err != nil {
		log.Println(err)

		/*file, errOF := os.OpenFile("../errors.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

		if errOF != nil {
			return ""
		}

		log.SetOutput(file)

		log.Println("daysteps: ", err)

		defer file.Close()
		/**/

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
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %v.\nДистанция составила %0.2f км.\nВы сожгли %0.2f ккал.\n", steps, distance, calories)
}
