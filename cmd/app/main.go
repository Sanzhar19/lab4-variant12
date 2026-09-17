package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/uuid"

	"github.com/Sanzhar19/lab4-variant12/pkg/mobiletariff"
)

func main() {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)

	cyan.Println("=== Система расчёта мобильного тарифа ===")

	// Вызов внешней библиотеки №1 (uuid)
	transactionID := uuid.New()
	fmt.Printf("ID транзакции: %s\n\n", transactionID.String())

	userName := "Sanzhar"
	minutes := 150.5
	pricePerMin := 12.0
	gb := 25.0
	pricePerGb := 150.0
	discountPercent := 15.0

	// 1. Звонки
	callCost, err := mobiletariff.CallCost(minutes, pricePerMin)
	if err != nil {
		red.Printf("Ошибка: %v\n", err)
		return
	}

	// 2. Интернет
	netCost, err := mobiletariff.InternetCost(gb, pricePerGb)
	if err != nil {
		red.Printf("Ошибка: %v\n", err)
		return
	}

	total := callCost + netCost

	// 3. Скидка (через указатель &total)
	err = mobiletariff.ApplyPackageDiscount(&total, discountPercent)
	if err != nil {
		red.Printf("Ошибка: %v\n", err)
		return
	}

	// 4. Формирование отчёта
	report, err := mobiletariff.FormatTariffReport(userName, callCost, netCost, total)
	if err != nil {
		red.Printf("Ошибка: %v\n", err)
		return
	}

	// Вызов внешней библиотеки №2 (fatih/color)
	green.Println(report)
}