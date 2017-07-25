package Index

import "gopkg.in/kataras/iris.v6"

//Get Renderiza Pagina principal
func Get(ctx *iris.Context) {
	ctx.Render("index.html", nil)
}
