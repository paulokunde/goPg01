package api

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	"app/autentica"
)

func (server *Server) Login(c *gin.Context) {
	c.HTML(
		http.StatusOK, "index.html", gin.H{
			"title": "Home Page",
		},
	)
}

type autenticaRequest struct {
	Usuario string `json:"usuario" binding:"required"`
	Senha   string `json:"senha" binding:"required"`
}

func (server *Server) Autentica(ctx *gin.Context) {
	var req autenticaRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	usuario := req.Usuario
	senha := req.Senha

	if autentica.Autentica(usuario, senha) {
		http.ServeFile(ctx.Writer, ctx.Request, "./static/index.html")
	} else {
		http.ServeFile(ctx.Writer, ctx.Request, "./static/login.html")
	}

}

func AutenticaAd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usuario := r.Form.Get("usuario")
	senha := r.Form.Get("senha")
	campos := map[string]interface{}{
		"usuario": usuario,
	}
	if autentica.Autentica(usuario, senha) {
		outputHTML(w, "./static/altera.html", campos)
	} else {
		http.ServeFile(w, r, "./static/login.html")
	}
}

func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
