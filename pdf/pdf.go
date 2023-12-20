package pdf

import "github.com/jung-kurt/gofpdf"

func AddHeadText(pdf *gofpdf.Fpdf, spacing, x float64, align, text string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "B", 9)

	// Get the current Y position
	currentY := pdf.GetY()

	// Set the X position
	pdf.SetX(x)
	// Add the text
	pdf.CellFormat(0, 10, text, "0", 1, align, false, 0, "")
	//pdf.Text(147, 140, "Juru Bayar")

	// Adjust the Y position to create spacing
	pdf.SetY(currentY + spacing)

	return pdf
}
