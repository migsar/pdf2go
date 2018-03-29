package main

import (
  "fmt"
  "log"
  "github.com/BurntSushi/toml"
  "github.com/jung-kurt/gofpdf"
)

type Field struct {
  Name string
  Type string
}

type Template struct {
 TemplateName string
 Fields []Field
}

func main() {
  // argsOnly := os.Args[1:]

  var template Template
  if _, err := toml.DecodeFile("./layouts/ModelRelease.toml", &template); err != nil {
    log.Fatal(err)
  }

  fmt.Println(template.Fields)

  doc := gofpdf.New("P", "mm", "A4", "")
  doc.SetTitle("Model Release", true)
  doc.SetDisplayMode("fullwidth", "continuous")
  doc.SetCreator("PDF2Go", true)
  doc.SetAuthor("PDF2Go", true)
  doc.SetKeywords("PDF2Go, PDF, Go, generator, templates", true)
  doc.AddPage()
  doc.SetFont("Arial","B", 16)
  doc.Cell(40, 10, "Hola, mundo")
  doc.SetDrawColor(255, 0, 0)
  doc.SetLineWidth(2.0)
  doc.Line(40, 40, 200, 40)
  doc.OutputFileAndClose("test.pdf")
}
