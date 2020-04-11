package templates

type Field struct {
  Name string
  Type string
}

type Component struct {
  Name string
  Type string
}

type Template struct {
  TemplateName string
  Components []Component
  Fields []Field
}
