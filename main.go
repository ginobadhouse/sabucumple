package main

import (
  "fmt"
  "net/http"
  "strings"

  "github.com/labstack/echo/v5"

  "github.com/heizeisaburou/sabucumple/module"
  "github.com/heizeisaburou/sabucumple/people/ejemplo"
)

func main() {
  e := echo.New()

  modules := []module.Module{
    ejemplo.New(),
  }

  e.GET("/", func(c *echo.Context) error {
    var html strings.Builder
    html.WriteString(`
      <h1>Cumple 🎂</h1>
      <ul>
    `)

    for _, m := range modules {
      html.WriteString(`<li><a href="/` + m.Endpoint() + `">` + m.Endpoint() + `</a></li>`)
    }

    html.WriteString(`</ul>`)

    return c.HTML(http.StatusOK, html.String())
  })

  for _, m := range modules {
    g := e.Group("/" + m.Endpoint())

    // Placeholder para que /chavsi no pete aunque Register esté vacío.
    g.GET("", placeholder(m.Endpoint()))

    // Rutas propias de cada persona.
    m.Register(g)
  }

  e.Static("/static", "static")

  fmt.Println("Servidor escuchando en http://localhost:8080")
  e.Start(":8080")

}

func placeholder(name string) echo.HandlerFunc {
  return func(c *echo.Context) error {
    return c.HTML(http.StatusOK, `
      <h1>Zona de `+name+` 🚧</h1>
      <p>Todavía no hay nada aquí.</p>
      <a href="/">Volver</a>
    `)
  }
}
