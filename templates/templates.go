package templates

type Field struct {
  Name string
  Type string
}

type Template struct {
 TemplateName string
 Fields []Field
}
