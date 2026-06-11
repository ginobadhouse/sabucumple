package ejemplo

import (
  "net/http"

  "github.com/labstack/echo/v5"
)

// Puedes poner tu nombre en lugar de Module
type Module struct{}

// New y Name, Register deben respetar exactamente esa interfaz y manera de uso
func New() Module {
  return Module{}
}

func (Module) Endpoint() string {
  return "ejemplo"
}

func (Module) Register(g *echo.Group) {
  g.GET("/", home)
  g.GET("/otro", otro)
}

func home(c *echo.Context) error {
  return c.HTML(http.StatusOK, `
    <h1>Zona de MY_NAME</h1>
    <p>Aquí MY_NAME puede hacer lo que quiera.</p>
    <a href="/ejemplo/otro">Ver ejemplo</a>
  `)
}

func otro(c *echo.Context) error {
  return c.HTML(http.StatusOK, `
    <h1>SubEndpoint de ejemplo</h1>
    <p>blablabla.</p>
  `)
}
