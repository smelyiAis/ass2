package main

import (
	"fmt"
)

// Report represents the product interface.
type Report interface {
	Generate() string
}

// PDFReport is a concrete product that generates PDF reports.
type PDFReport struct{}

func (p PDFReport) Generate() string {
	return "Generating PDF report..."
}

// ExcelReport is a concrete product that generates Excel reports.
type ExcelReport struct{}

func (e ExcelReport) Generate() string {
	return "Generating Excel report..."
}

// ReportGenerator is the creator interface.
type ReportGenerator interface {
	GenerateReport() Report
}

// PDFReportGenerator is a concrete creator that generates PDF reports.
type PDFReportGenerator struct{}

func (p PDFReportGenerator) GenerateReport() Report {
	return PDFReport{}
}

// ExcelReportGenerator is a concrete creator that generates Excel reports.
type ExcelReportGenerator struct{}

func (e ExcelReportGenerator) GenerateReport() Report {
	return ExcelReport{}
}


type Application struct{}

func (app Application) CreateReport(generator ReportGenerator) string {
	report := generator.GenerateReport()
	return report.Generate()
}

func main() {
	app := Application{}

	pdfGenerator := PDFReportGenerator{}
	excelGenerator := ExcelReportGenerator{}

	pdfReport := app.CreateReport(pdfGenerator)
	excelReport := app.CreateReport(excelGenerator)

	fmt.Println(pdfReport)
	fmt.Println(excelReport)
}
