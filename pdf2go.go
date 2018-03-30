package main

import (
  "fmt"
  "log"
  "flag"
  "github.com/BurntSushi/toml"
  "github.com/jung-kurt/gofpdf"
)

type PDF struct {
  File *gofpdf.Fpdf
  Filename string
}

type Field struct {
  Name string
  Type string
}

type Template struct {
 TemplateName string
 Fields []Field
}

func (pdf *PDF) init() {
  var doc = pdf.File
  doc.SetDisplayMode("fullwidth", "continuous")
  doc.SetCreator("PDF2Go", true)
  doc.SetAuthor("PDF2Go", true)
  // doc.SetKeywords("PDF2Go, PDF, Go, generator, templates", true)
  doc.SetFont("Arial","", 16)
}

func (pdf *PDF) save() {
  var doc = pdf.File
  doc.OutputFileAndClose(pdf.Filename)
}

func (pdf *PDF) buildFromTemplate() {
  var doc = pdf.File
  // doc.SetTitle("Model Release", true)
  doc.AddPage()
  doc.MultiCell(0, 10, "Hola, mundo Hola, mundo Holala, mundo Hola, mundo Hola, mundo Hola, mundo Hola, mundo Hola, mundo Hola, mundo ", "", "L", false)
  // doc.SetDrawColor(255, 0, 0)
  // doc.SetLineWidth(2.0)
  // doc.Line(40, 40, 200, 40)
}

func main() {
  inputFilePtr := flag.String("in", "", "Input file name. Required.")
  outputFilePtr := flag.String("out", "", "Output file name. Required.")
  templatePtr := flag.String("template", "CV", "Template file name.")

  flag.Parse()

  if *inputFilePtr == "" {
    log.Fatal("Input file is required. run 'pdf2go -h' for help.")
  }

  if *outputFilePtr == "" {
    log.Fatal("Output file is required. run 'pdf2go -h' for help.")
  }
  _ = templatePtr

  var template Template
  if _, err := toml.DecodeFile("./layouts/ModelRelease.toml", &template); err != nil {
    log.Fatal(err)
  }

  fmt.Println(template.Fields)

  doc := &PDF{ gofpdf.New("P", "mm", "A4", ""), *outputFilePtr }

  doc.init()
  doc.buildFromTemplate()
  doc.save()
}
