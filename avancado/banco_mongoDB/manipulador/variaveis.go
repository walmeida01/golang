package manipulador

import "text/template"

//ModeloOla armazena os modelos da pagina ola
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal armazena os modelos da pagina local
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
