package templates

import (
  "github.com/jung-kurt/gofpdf"
)

type PDF struct {
  File *gofpdf.Fpdf
  Filename string
}

// Orientation:
// "P" or "Portrait" for portrait
// "L" or "Landscape" for landscape
// "" => "P"
// Unit:
// "pt" for point
// "mm" for millimeter
// "cm" for centimeter
// "in" for inch
// "" => "mm"
// Size: Page size
// "A3", "A4", "A5", "Letter", "Legal", "Tabloid"
// "" => "A4"
// Fontdir: The location of font resources
// "" => "."
// This argument only needs to reference an actual directory if a font other than
// one of the core fonts is used.
// Core fonts are "courier", "helvetica" (also called "arial"), "times"
// and "zapfdingbats" (also called "symbol").
type PDFConstructor struct {
  Orientation, Unit, Size, Fontdir string
}

func (pdf *PDF) Init(options *PDFConstructor) {
  pdf.File = gofpdf.New(
    options.Orientation,
    options.Unit,
    options.Size,
    options.Fontdir)

  var doc = pdf.File
  doc.SetDisplayMode("fullwidth", "continuous")
  doc.SetCreator("PDF2Go", true)
  doc.SetAuthor("PDF2Go", true)
  // doc.SetKeywords("PDF2Go, PDF, Go, generator, templates", true)
  doc.SetFont("Arial","", 16)
}

func (pdf *PDF) Save() {
  var doc = pdf.File
  doc.OutputFileAndClose(pdf.Filename)
}

func (pdf *PDF) BuildFromTemplate() {
  // var doc = pdf.File
  // doc.SetTitle("Model Release", true)
  // doc.SetDrawColor(255, 0, 0)
  // doc.SetLineWidth(2.0)
  // doc.Line(40, 40, 200, 40)
}

func (pdf *PDF) Page() {
  var doc = pdf.File
  doc.AddPage()
}

func (pdf *PDF) Paragraph() {
  var doc = pdf.File
  doc.MultiCell(0, 10, "Hola, mundo Hola, mundo Holala, mundo Hola, mundo Hola, mundo Hola, mundo Hola, mundo Hola, mundo Hola, mundo ", "", "L", false)
}

func (pdf *PDF) Group() {

}

func (pdf *PDF) Column() {

}

func (pdf *PDF) getState() {

}

func (pdf *PDF) restoreState() {

}
