package Adminsistrador

import "gopkg.in/kataras/iris.v6"

//Get Renderiza Pagina principal
func Get(ctx *iris.Context) {
	ctx.Render("Administrar/index.html", nil)
}
