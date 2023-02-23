package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"app/modelo"
)

func (server *Server) teste(c *gin.Context) {
	c.HTML(http.StatusOK, "teste.html", gin.H{})
}

func (server *Server) formMarca(c *gin.Context) {
	c.HTML(
		http.StatusOK, "marca.html", gin.H{
			"title":  "Mestre Detalhe",
			"limit":  "10",
			"offset": "0",
		},
	)
}
func (server *Server) editMarca(c *gin.Context) {
	c.HTML(
		http.StatusOK, "editMarca.html", gin.H{
			"title": "Alterar Marca",
		},
	)
}
func (server *Server) addMarca(c *gin.Context) {
	c.HTML(
		http.StatusOK, "addMarca.html", gin.H{
			"title": "Registro de Marca",
		},
	)
}
func (server *Server) viewMarca(c *gin.Context) {
	c.HTML(
		http.StatusOK, "viewMarca.html", gin.H{
			"title": "Marca",
		},
	)
}
func (server *Server) createMarca(c *gin.Context) {
	nome := c.PostForm("nome")

	//create new reg
	marca, er := server.store.CreateMarca(c, nome)
	if er != nil {
		c.HTML(http.StatusInternalServerError, "error.html", errorResponse(er))
	} else {
		c.HTML(
			http.StatusOK, "addMarca.html", gin.H{
				"title": "Marca Registrada:",
				"msg":   "Registrado:" + marca.Nome,
			},
		)
	}

}

type getMarcaRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getMarca(c *gin.Context) {
	var req getMarcaRequest
	er1 := c.ShouldBindUri(&req)
	if er1 != nil {
		c.HTML(http.StatusBadRequest, "error.html", errorResponse(er1))
	}
	marca, er2 := server.store.GetMarcaById(c, req.ID)
	if er2 != nil {
		c.HTML(http.StatusInternalServerError, "error.html", errorResponse(er2))
	} else {
		c.HTML(http.StatusOK, "editMarca.html", gin.H{
			"marca": marca,
			"title": "Atualizar Marca",
		})
	}
}

type getMarcaRequestByName struct {
	Nome string `json:"nome" binding:"required"`
}

func (server *Server) getMarcaByName(ctx *gin.Context) {
	var req getMarcaRequestByName
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	marca, err := server.store.GetMarcaByName(ctx, req.Nome)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, marca)
}

type deleteMarcaRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteMarca(ctx *gin.Context) {
	var req deleteMarcaRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", errorResponse(err))
	}

	er := server.store.DeleteMarca(ctx, req.ID)
	if er != nil {
		ctx.HTML(http.StatusOK, "error.html", errorResponse(er))
	}
	ctx.HTML(http.StatusOK, "marca.html", gin.H{
		"msg": "Registro Deletado - OK",
	})
}

type updateRequest struct {
	ID   int32  `form:"id" json:"id"`
	Nome string `form:"nome" json:"nome"`
}

func (server *Server) updateMarca(ctx *gin.Context) {
	var req updateRequest
	err := ctx.Bind(&req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	arg := modelo.UpdateMarcaParams{
		ID:   req.ID,
		Nome: req.Nome,
	}
	marca, er := server.store.UpdateMarca(ctx, arg)
	if er != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", errorResponse(er))

	}
	ctx.HTML(http.StatusOK, "marca.html", gin.H{
		"marca": marca,
		"msg":   "Atualizado...OK",
	})

}

func (server *Server) findMarcaByName(c *gin.Context) {
	var nome string = ""
	nomeP := c.PostForm("nome")
	nomeG, b1 := c.GetQuery("nome")
	//ctrl, _ := c.GetQuery("ctrl")
	lim, _ := c.GetQuery("limit")
	off, _ := c.GetQuery("offset")

	if b1 == true {
		nome = nomeG
	} else {
		nome = nomeP
	}
	l, er := strconv.Atoi(lim)
	o, err := strconv.Atoi(off)
	if err != nil || er != nil {
		l = 10
		o = 0
	}
	arg := modelo.FindMarcaByNameParams{
		Nome:   nome + "%",
		Limit:  int32(l),
		Offset: int32(o),
	}
	marcas, err := server.store.FindMarcaByName(c, arg)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", errorResponse(err))
	} else {
		tam := len(marcas)
		/*if ctrl == "next" {
			o = o + l
		} else if ctrl == "prev" && o > l {
			o = o - l
		}
		*/
		o = o + l
		c.HTML(http.StatusOK, "marca.html", gin.H{
			"marcas": marcas,
			"limit":  l,
			"offset": o,
			"nome":   nome,
			"msg":    "Retorno:" + strconv.Itoa(tam) + " Itens",
		})
	}
}

func (server *Server) listMarcas(ctx *gin.Context) {

	marcas, err := server.store.ListMarcas(ctx)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", errorResponse(err))
	}

	msg := fmt.Sprint("Listados:", len(marcas))
	ctx.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"listMarca.html",
		// Pass the data that the page uses
		gin.H{
			"title":  "Lista Marcas",
			"marcas": marcas,
			"msg":    msg,
		},
	)

}
