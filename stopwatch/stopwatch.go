package stopwatch

import (
	"fmt"
	"time"
)

// Função para calcular e formatar a duração
func FormatDuration(d time.Duration) string {
	return fmt.Sprintf("%.3f seconds", d.Seconds())
}

// Função para medir o tempo de execução
func MeasureExecutionTime(startTime time.Time) string {
	duration := time.Since(startTime)
	return FormatDuration(duration)
}
