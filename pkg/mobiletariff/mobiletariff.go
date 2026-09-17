package mobiletariff

import (
	"fmt"
)

// CallCost рассчитывает стоимость звонков по количеству минут и тарифу.
func CallCost(minutes float64, pricePerMin float64) (float64, error) {
	if minutes < 0 {
		return 0, fmt.Errorf("минуты не могут быть отрицательными: %.2f", minutes)
	}
	if pricePerMin < 0 {
		return 0, fmt.Errorf("цена за минуту не может быть отрицательной: %.2f", pricePerMin)
	}
	return minutes * pricePerMin, nil
}

// InternetCost рассчитывает стоимость интернет-трафика в ГБ.
func InternetCost(gb float64, pricePerGb float64) (float64, error) {
	if gb < 0 {
		return 0, fmt.Errorf("объём трафика не может быть отрицательным: %.2f GB", gb)
	}
	if pricePerGb < 0 {
		return 0, fmt.Errorf("цена за ГБ не может быть отрицательной: %.2f", pricePerGb)
	}
	return gb * pricePerGb, nil
}

// ApplyPackageDiscount применяет процентную скидку к итоговой сумме через указатель.
func ApplyPackageDiscount(total *float64, percent float64) error {
	if total == nil {
		return fmt.Errorf("указатель на итоговую сумму не может быть nil")
	}
	if *total < 0 {
		return fmt.Errorf("итоговая сумма не может быть отрицательной: %.2f", *total)
	}
	if percent < 0 || percent > 100 {
		return fmt.Errorf("некорректный процент скидки: %.2f%% (допустимо от 0 до 100)", percent)
	}

	discountAmount := (*total) * (percent / 100.0)
	*total = *total - discountAmount
	return nil
}

// FormatTariffReport формирует форматированную строку отчёта.
func FormatTariffReport(user string, calls, internet, total float64) (string, error) {
	if user == "" {
		return "", fmt.Errorf("имя пользователя не может быть пустым")
	}
	if calls < 0 || internet < 0 || total < 0 {
		return "", fmt.Errorf("финансовые показатели не могут быть отрицательными")
	}

	report := fmt.Sprintf(
		"Отчёт по тарифу для: %-15s | Звонки: %6.2f ₸ | Интернет: %6.2f ₸ | Итого к оплате: %7.2f ₸",
		user, calls, internet, total,
	)
	return report, nil
}