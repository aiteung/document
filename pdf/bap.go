package pdf

import (
	"github.com/aiteung/document/csv"
	"github.com/jung-kurt/gofpdf"
	"os"
	"strings"
)

const InfoImageURL = "https://home.ulbi.ac.id/ulbi.png"

func CreateHeaderBAP(Text []string, x float64) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Times", "B", 12)
	pdf.SetX(x)
	pdf.CellFormat(70, 10, Text[0], "0", 0, "C", false, 0, "")
	pdf.Ln(5)
	pdf.SetX(x)
	pdf.CellFormat(70, 10, Text[1], "0", 0, "C", false, 0, "")
	pdf.Ln(5)

	pdf.SetY(20)

	return pdf
}

func CreatePDFBAP(Text []string, filename, filename2, filename3, filename4, filenamee5 string) string {
	data := csv.CSVLoader(filename)
	data2 := csv.CSVLoader(filename2)
	data3 := csv.CSVLoader(filename3)
	data4 := csv.CSVLoader(filename3)
	data5 := csv.CSVLoader(filenamee5)
	pdf := CreateHeaderBAP(Text, 90)
	pdf = ImageCustomize(pdf, "./ulbi.png", InfoImageURL, 28, 11, 35, 12, 100, 100, 0.3)
	width := []float64{60, 5, 70}
	widthPertemuan := []float64{20, 20, 50, 30, 30}
	widthPertemuan1 := []float64{25, 60, 6, 6, 6, 6, 6, 6, 6, 6, 17}
	widthPertemuan2 := []float64{25, 20, 11, 11, 11, 11, 11, 11, 11, 11, 17}
	color := []int{255, 255, 153}
	align := []string{"J", "C", "J"}
	alignPertemuan := []string{"C", "C", "C", "C", "C"}
	alignPertemuan1 := []string{"C", "C", "C", "C", "C", "C", "C", "C", "C", "C", "C"}
	alignPertemuan2 := []string{"C", "C", "C", "C", "C", "C", "C", "C", "C", "C", "C"}
	yCoordinates := []float64{40, 45, 50}
	pdf = SetTableContentCustomY(pdf, data[1:], width, align, yCoordinates)
	pdf.Ln(5)
	pdf = SetMergedCell(pdf, "Tabel Log Aktivitas", "J", 150, color)
	pdf = SetHeaderTable(pdf, data2[0], widthPertemuan, color)
	pdf = SetKambingContent(pdf, data2[1:], widthPertemuan, alignPertemuan)

	pdf.Ln(10)
	pdf = SetMergedCell(pdf, "Tabel Presensi Paruh Pertama", "J", 150, color)
	pdf = SetHeaderTable(pdf, data3[0], widthPertemuan1, color)
	pdf = SetTableContent(pdf, data3[1:], widthPertemuan1, alignPertemuan1)

	pdf.Ln(10)
	pdf = SetMergedCell(pdf, "Tabel Presensi Paruh Kedua", "J", 150, color)
	pdf = SetHeaderTable(pdf, data4[0], widthPertemuan1, color)
	pdf = SetTableContent(pdf, data4[1:], widthPertemuan1, alignPertemuan1)

	pdf.Ln(10)
	pdf = SetMergedCell(pdf, "Tabel Nilai Akhir", "J", 150, color)
	pdf = SetHeaderTable(pdf, data5[0], widthPertemuan2, color)
	pdf = SetTableContent(pdf, data5[1:], widthPertemuan2, alignPertemuan2)

	csv := filename
	csv = strings.ReplaceAll(csv, "csv", "pdf")
	err := SavePDF(pdf, csv)
	if err != nil {
		return err.Error()
	}
	os.Remove(filename)
	os.Remove(filename2)
	os.Remove(filename3)
	os.Remove(filename4)
	os.Remove(filenamee5)
	return csv
}
