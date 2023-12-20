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

func AddNameText(pdf *gofpdf.Fpdf, Text string, spacing, x, size float64) *gofpdf.Fpdf {

	pdf.SetFont("Times", "B", size)
	//pdf.Text(137, 138, Text)
	pdf.SetX(x)
	pdf.CellFormat(0, 10, Text, "0", 0, "C", false, 0, "")
	pdf.Ln(0.5 * size)

	currentY := pdf.GetY()

	pdf.SetY(currentY + spacing)

	return pdf
}

func SetMergedCell(pdf *gofpdf.Fpdf, text string, width float64) *gofpdf.Fpdf {
	pdf.SetFont("Times", "B", 10)
	pdf.SetFillColor(255, 165, 0)
	totalWidth := 0.0
	totalWidth += width

	// Calculate the X-coordinate to center the table on the page
	pageWidth, _ := pdf.GetPageSize()
	x := (pageWidth - totalWidth) / 2

	// Set the X-coordinate
	pdf.SetX(x)

	// Create 6 cells that make up the merged cell
	pdf.CellFormat(width, 7, text, "1", 0, "L", true, 0, "")

	// Move to the next line
	pdf.Ln(-1)
	return pdf
}

func SetHeaderTable(pdf *gofpdf.Fpdf, hdr []string, widths []float64) *gofpdf.Fpdf {
	pdf.SetFont("Times", "B", 8)
	pdf.SetFillColor(240, 240, 240)
	// Calculate the total width of the table
	totalWidth := 0.0
	for _, width := range widths {
		totalWidth += width
	}

	// Calculate the X-coordinate to center the table on the page
	pageWidth, _ := pdf.GetPageSize()
	x := (pageWidth - totalWidth) / 2

	// Set the X-coordinate
	pdf.SetX(x)
	for i, str := range hdr {
		// The `CellFormat()` method takes a couple of parameters to format
		// the cell. We make use of this to create a visible border around
		// the cell, and to enable the background fill.
		pdf.CellFormat(widths[i], 7, str, "1", 0, "C", true, 0, "")
	}

	// Passing `-1` to `Ln()` uses the height of the last printed cell as
	// the line height.
	pdf.Ln(-1)
	return pdf
}

func SetTableContent(pdf *gofpdf.Fpdf, tbl [][]string, widths []float64, align []string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "", 8)
	pdf.SetFillColor(255, 255, 255)

	for _, line := range tbl {
		// Calculate the total width of the table
		totalWidth := 0.0
		for _, width := range widths {
			totalWidth += width
		}

		// Calculate the X-coordinate to center the table on the page
		pageWidth, _ := pdf.GetPageSize()
		x := (pageWidth - totalWidth) / 2

		// Set the X-coordinate
		pdf.SetX(x)
		for i, str := range line {
			pdf.CellFormat(widths[i], 7, str, "1", 0, align[i], true, 0, "")
		}
		pdf.Ln(-1)
	}
	return pdf
}
