package services

import (
	"bytes"
	"fmt"
	"time"

	"github.com/LuneLoops/MIAFIN/internal/database"
	"github.com/LuneLoops/MIAFIN/pkg/models"
	"github.com/jung-kurt/gofpdf"
)

type ReporteService struct{}

func NewReporteService() *ReporteService {
	return &ReporteService{}
}

// GenerarReporteGeneral genera un PDF con el reporte financiero general
func (s *ReporteService) GenerarReporteGeneral() ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Courier", "B", 16)

	// Título
	pdf.Cell(190, 10, "MIAFIN - Reporte Financiero General")
	pdf.Ln(15)

	// Fecha
	pdf.SetFont("Courier", "", 10)
	pdf.Cell(190, 6, fmt.Sprintf("Fecha: %s", time.Now().Format("02/01/2006 15:04")))
	pdf.Ln(10)

	// Obtener métricas
	var totalCartera float64
	database.DB.Model(&models.Prestamo{}).
		Where("estado IN ?", []string{"aprobado", "desembolsado"}).
		Select("COALESCE(SUM(monto), 0)").Scan(&totalCartera)

	var totalDesembolsado float64
	database.DB.Model(&models.Prestamo{}).
		Where("estado IN ?", []string{"desembolsado", "pagado"}).
		Select("COALESCE(SUM(monto), 0)").Scan(&totalDesembolsado)

	var totalRecuperado float64
	database.DB.Model(&models.Pago{}).
		Select("COALESCE(SUM(monto), 0)").Scan(&totalRecuperado)

	var totalSaldoPendiente float64
	database.DB.Model(&models.Prestamo{}).
		Where("estado != ?", "pagado").
		Select("COALESCE(SUM(saldo_actual), 0)").Scan(&totalSaldoPendiente)

	liquidez := totalRecuperado - totalSaldoPendiente

	// Datos
	pdf.SetFont("Courier", "B", 12)
	pdf.Cell(190, 8, "Metricas Financieras")
	pdf.Ln(10)

	pdf.SetFont("Courier", "", 10)
	pdf.Cell(100, 6, "Total cartera activa:")
	pdf.Cell(90, 6, fmt.Sprintf("Bs. %.2f", totalCartera))
	pdf.Ln(8)

	pdf.Cell(100, 6, "Total desembolsado:")
	pdf.Cell(90, 6, fmt.Sprintf("Bs. %.2f", totalDesembolsado))
	pdf.Ln(8)

	pdf.Cell(100, 6, "Total recuperado:")
	pdf.Cell(90, 6, fmt.Sprintf("Bs. %.2f", totalRecuperado))
	pdf.Ln(8)

	pdf.Cell(100, 6, "Liquidez actual:")
	pdf.Cell(90, 6, fmt.Sprintf("Bs. %.2f", liquidez))
	pdf.Ln(15)

	// Footer
	pdf.SetFont("Courier", "I", 8)
	pdf.Cell(190, 6, "Generado por MIAFIN")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GenerarReportePrestamo genera un PDF con los detalles de un préstamo
func (s *ReporteService) GenerarReportePrestamo(prestamoID uint) ([]byte, error) {
	var prestamo models.Prestamo
	result := database.DB.Preload("Cliente").Preload("Grupo").First(&prestamo, prestamoID)
	if result.Error != nil {
		return nil, fmt.Errorf("préstamo no encontrado")
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Courier", "B", 16)

	// Título
	pdf.Cell(190, 10, "MIAFIN - Reporte de Prestamo")
	pdf.Ln(15)

	// Fecha
	pdf.SetFont("Courier", "", 10)
	pdf.Cell(190, 6, fmt.Sprintf("Fecha: %s", time.Now().Format("02/01/2006 15:04")))
	pdf.Ln(10)

	// Datos del préstamo
	pdf.SetFont("Courier", "B", 12)
	pdf.Cell(190, 8, "Datos del Prestamo")
	pdf.Ln(10)

	pdf.SetFont("Courier", "", 10)
	pdf.Cell(50, 6, "Tipo:")
	pdf.Cell(140, 6, prestamo.Tipo)
	pdf.Ln(6)

	if prestamo.Cliente != nil {
		pdf.Cell(50, 6, "Cliente:")
		pdf.Cell(140, 6, prestamo.Cliente.Nombre)
		pdf.Ln(6)
		pdf.Cell(50, 6, "CI:")
		pdf.Cell(140, 6, prestamo.Cliente.CI)
		pdf.Ln(6)
	}

	if prestamo.Grupo != nil {
		pdf.Cell(50, 6, "Grupo:")
		pdf.Cell(140, 6, prestamo.Grupo.Nombre)
		pdf.Ln(6)
	}

	pdf.Cell(50, 6, "Monto:")
	pdf.Cell(140, 6, fmt.Sprintf("Bs. %.2f", prestamo.Monto))
	pdf.Ln(6)

	pdf.Cell(50, 6, "Tasa de interes:")
	pdf.Cell(140, 6, fmt.Sprintf("%.2f%%", prestamo.TasaInteres))
	pdf.Ln(6)

	pdf.Cell(50, 6, "Plazo:")
	pdf.Cell(140, 6, fmt.Sprintf("%d meses", prestamo.PlazoMeses))
	pdf.Ln(6)

	pdf.Cell(50, 6, "Saldo actual:")
	pdf.Cell(140, 6, fmt.Sprintf("Bs. %.2f", prestamo.SaldoActual))
	pdf.Ln(6)

	pdf.Cell(50, 6, "Estado:")
	pdf.Cell(140, 6, prestamo.Estado)
	pdf.Ln(15)

	// Tabla de amortización simple
	pdf.SetFont("Courier", "B", 12)
	pdf.Cell(190, 8, "Plan de Amortizacion")
	pdf.Ln(10)

	pdf.SetFont("Courier", "B", 9)
	pdf.Cell(20, 6, "Mes")
	pdf.Cell(50, 6, "Cuota")
	pdf.Ln(8)

	// Calcular cuota mensual
	interes := prestamo.Monto * (prestamo.TasaInteres / 100)
	totalAPagar := prestamo.Monto + interes
	cuotaMensual := totalAPagar / float64(prestamo.PlazoMeses)

	pdf.SetFont("Courier", "", 9)
	for i := 1; i <= prestamo.PlazoMeses; i++ {
		pdf.Cell(20, 6, fmt.Sprintf("%d", i))
		pdf.Cell(50, 6, fmt.Sprintf("Bs. %.2f", cuotaMensual))
		pdf.Ln(6)
	}

	pdf.Ln(15)

	// Firmas
	pdf.SetFont("Courier", "B", 10)
	pdf.Cell(190, 6, "Firma del Cliente: ______________________")
	pdf.Ln(10)
	pdf.Cell(190, 6, "Firma del Asesor: ______________________")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GenerarComprobantePago genera un PDF de comprobante de pago
func (s *ReporteService) GenerarComprobantePago(pagoID uint) ([]byte, error) {
	var pago models.Pago
	result := database.DB.Preload("Prestamo.Cliente").Preload("Prestamo.Grupo").First(&pago, pagoID)
	if result.Error != nil {
		return nil, fmt.Errorf("pago no encontrado")
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Courier", "B", 16)

	// Título
	pdf.Cell(190, 10, "MIAFIN - Comprobante de Pago")
	pdf.Ln(15)

	// Fecha
	pdf.SetFont("Courier", "", 10)
	pdf.Cell(190, 6, fmt.Sprintf("Fecha: %s", pago.FechaPago.Format("02/01/2006 15:04")))
	pdf.Ln(10)

	// Datos del pago
	pdf.SetFont("Courier", "B", 12)
	pdf.Cell(190, 8, "Datos del Pago")
	pdf.Ln(10)

	pdf.SetFont("Courier", "", 10)
	pdf.Cell(50, 6, "ID Pago:")
	pdf.Cell(140, 6, fmt.Sprintf("%d", pago.ID))
	pdf.Ln(6)

	pdf.Cell(50, 6, "ID Prestamo:")
	pdf.Cell(140, 6, fmt.Sprintf("%d", pago.PrestamoID))
	pdf.Ln(6)

	if pago.Prestamo != nil {
		if pago.Prestamo.Cliente != nil {
			pdf.Cell(50, 6, "Cliente:")
			pdf.Cell(140, 6, pago.Prestamo.Cliente.Nombre)
			pdf.Ln(6)
		}
		if pago.Prestamo.Grupo != nil {
			pdf.Cell(50, 6, "Grupo:")
			pdf.Cell(140, 6, pago.Prestamo.Grupo.Nombre)
			pdf.Ln(6)
		}
	}

	pdf.Cell(50, 6, "Monto pagado:")
	pdf.Cell(140, 6, fmt.Sprintf("Bs. %.2f", pago.Monto))
	pdf.Ln(15)

	// Firma
	pdf.SetFont("Courier", "B", 10)
	pdf.Cell(190, 6, "Firma del Cliente: ______________________")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GetLiquidez calcula la liquidez actual
func (s *ReporteService) GetLiquidez() (map[string]interface{}, error) {
	var totalRecuperado float64
	database.DB.Model(&models.Pago{}).
		Select("COALESCE(SUM(monto), 0)").Scan(&totalRecuperado)

	var totalSaldoPendiente float64
	database.DB.Model(&models.Prestamo{}).
		Where("estado != ?", "pagado").
		Select("COALESCE(SUM(saldo_actual), 0)").Scan(&totalSaldoPendiente)

	liquidez := totalRecuperado - totalSaldoPendiente

	return map[string]interface{}{
		"total_recuperado":      totalRecuperado,
		"total_saldo_pendiente": totalSaldoPendiente,
		"liquidez":              liquidez,
	}, nil
}
