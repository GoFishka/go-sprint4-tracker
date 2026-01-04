package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	//"github.com/golang/protobuf/ptypes/duration"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

// Разбирает входящую строку с данными о тренировке (разделитель запятая) и возвращает:
// количество шагов,
// тип тренировки,
// продолжительность тренировки
// и ошибку, если она возникла
func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию

	arData := strings.Split(data, ",")

	if len(arData) != 3 {
		return 0, "", 0, fmt.Errorf("Ошибка в формате данных")
	}

	for k, v := range arData {
		arData[k] = strings.TrimSpace(v)
	}

	steps, err := strconv.Atoi(arData[0])

	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка в количестве шагов: %v", err)
	}

	if steps < 1 {
		return 0, "", 0, fmt.Errorf("Количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(arData[2])

	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка в продолжительности тренировки: %v", err)
	}

	if !(duration > 0) {
		return 0, "", 0, fmt.Errorf("Ошибка в продолжительности тренировки: отрицательная - %v", duration)
	}

	activityType := arData[1]

	if activityType != "Бег" && activityType != "Ходьба" {
		return 0, "", 0, fmt.Errorf("неизвестный тип тренировки")
	}

	return steps, activityType, duration, nil
}

// возвращает дистанцию в километрах
func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	return float64(steps) * (height * stepLengthCoefficient) / mInKm
}

// фукнция принимает шаги, рост и продолжительность занятия
// возвращает среднюю скорость в км/ч
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	if !(duration.Hours() > 0) {
		return 0
	}

	if !(steps > 0) {
		return 0
	}

	if !(height > 0) {
		return 0
	}

	fullDistance := distance(steps, height)

	return fullDistance / (duration.Hours())
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию

	steps, activityType, duration, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		/*
			file, errOF := os.OpenFile("../errors.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

			if errOF != nil {
				return "", err
			}

			log.SetOutput(file)

			log.Println(err)

			defer file.Close()
		*/
		return "", err
	}

	switch activityType {
	case "Бег":
		calories, err := RunningSpentCalories(steps, weight, height, duration)

		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %v\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityType, duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), calories), nil

	case "Ходьба":
		calories, err := WalkingSpentCalories(steps, weight, height, duration)

		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %v\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityType, duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), calories), nil
	}

	return "", nil
}

// принимает шаги, вес в кг, рост в м, продолжительность бега
// возвращает количество калорий потраченных при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if !(steps > 0) {
		return 0, fmt.Errorf("Ошибка в количестве шагов: %v", steps)
	}

	if !(weight > 0) {
		return 0, fmt.Errorf("Ошибка в весе: %v", weight)
	}

	if !(height > 0) {
		return 0, fmt.Errorf("Ошибка в росте: %v", height)
	}

	if !(duration > 0) {
		return 0, fmt.Errorf("Ошибка в продолжительности: %v", duration)
	}

	return (weight * duration.Minutes() * meanSpeed(steps, height, duration)) / minInH, nil
}

// принимает шаги, вес в кг, рост в м, продолжительность ходьбы
// возвращает количество калорий потраченных при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	if !(steps > 0) {
		return 0, fmt.Errorf("Ошибка в количестве шагов: %v", steps)
	}

	if !(weight > 0) {
		return 0, fmt.Errorf("Ошибка в весе: %v", weight)
	}

	if !(height > 0) {
		return 0, fmt.Errorf("Ошибка в росте: %v", height)
	}

	if !(duration > 0) {
		return 0, fmt.Errorf("Ошибка в продолжительности: %v", duration)
	}

	//средняя скорость в км/ч
	meanSpeedValue := meanSpeed(steps, height, duration)

	//рассчитываем количество калорий
	calories := weight * meanSpeedValue * duration.Minutes() / 60

	return calories * walkingCaloriesCoefficient, nil
}
