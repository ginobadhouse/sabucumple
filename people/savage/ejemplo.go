package savage 

import (
  "io"
  "net/http"
  "text/template"

  "github.com/labstack/echo/v5"
)

// Puedes poner tu nombre en lugar de Module
type Module struct{}

// New y Name, Register deben respetar exactamente esa interfaz y manera de uso
func New() Module {
  return Module{}
}

const NAME = "savage";


type TemplateRenderer struct {
  templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(c *echo.Context, w io.Writer, name string, data any) error {

  // Add global methods if data is a map
  if viewContext, isMap := data.(map[string]any); isMap {
    viewContext["reverse"] = c.RouteInfo().Reverse
  }

  return t.templates.ExecuteTemplate(w, name, data)
}


func (Module) Endpoint() string {
  return NAME 
}

func (Module) Register(g *echo.Group) {
  g.GET("/", home)
}

func home(c *echo.Context) error {
  return c.HTML(http.StatusOK, ` ok`);
}
