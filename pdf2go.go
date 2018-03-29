package main

import (
  "fmt"
  "os"
  "log"
  "encoding/json"
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

  input, err := os.Open("./ModelRelease.json")
  if err != nil {
    log.Fatal(err)
  }

  dec := json.NewDecoder(input)

  var template Template
  err = dec.Decode(&template);
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(template)

  doc := gofpdf.New("P", "mm", "A4", "")
  doc.AddPage()
  doc.SetFont("Arial","B", 16)
  doc.Cell(40, 10, "Hola, mundo")
  doc.OutputFileAndClose("test.pdf")
}
