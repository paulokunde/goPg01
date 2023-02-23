package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"app/modelo"
)

// master form html
func (server *Server) formModelo(c *gin.Context) {
	c.HTML(
		http.StatusOK, "modelo.html", gin.H{
			"title":  "Mestre Detalhe",
			"limit":  "10",
			"offset": "0",
		},
	)
}

// create form html
func (server *Server) addModelo(c *gin.Context) {
	c.HTML(
		http.StatusOK, "addModelo.html", gin.H{
			"title": "Registro de Modelo",
		},
	)
}

func (server *Server) createModelo(c *gin.Context) {
	m_id := c.PostForm("marca_id")
	nome := c.PostForm("nome")
	log.Print("MarcaID: ")
	log.Println(m_id)
	log.Print("Nome: ")
	log.Println(nome)
	marca_id, err := strconv.Atoi(m_id)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", errorResponse(err))
		//marca_id = 9
	}
	log.Print("MarcaID: ")
	log.Println(marca_id)
	//create new reg
	modelo, er := server.store.CreateModelo(c, modelo.CreateModeloParams{
		Nome:    nome,
		MarcaID: int32(marca_id),
	})
	if er != nil {
		c.HTML(http.StatusInternalServerError, "error.html", errorResponse(err))
	} else {
		c.HTML(
			http.StatusOK, "addModelo.html", gin.H{
				"title": "Modelo Registrado:",
				"msg":   fmt.Sprintf("Registado:%v :ID %v", modelo.Nome, modelo.ID),
			},
		)
	}

}

type deleteModeloRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteModelo(ctx *gin.Context) {
	var req deleteModeloRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.HTML(http.StatusOK, "error.html", errorResponse(err))
	}

	er := server.store.DeleteModelo(ctx, req.ID)
	if er != nil {
		ctx.HTML(http.StatusOK, "error.html", errorResponse(er))
	}
	ctx.HTML(http.StatusOK, "modelo.html", gin.H{
		"msg": "Registro Deletado - OK",
	})
}

type updateModeloRequest struct {
	ID      int32  `form:"id"`
	Nome    string `form:"nome"`
	MarcaID int32  `form:"marca_id"`
}

func (server *Server) updateModelo(ctx *gin.Context) {
	var req updateModeloRequest
	//if ctx.Bind(&req) == nil {
	err := ctx.Bind(&req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	log.Println("====== Bind By Query String ======")
	log.Println(req.ID)
	log.Println(req.Nome)
	log.Println(req.MarcaID)
	log.Print("====================================")
	arg := modelo.UpdateModeloParams{
		ID:      req.ID,
		Nome:    req.Nome,
		MarcaID: req.MarcaID,
	}

	modelo, er := server.store.UpdateModelo(ctx, arg)
	if er != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", errorResponse(er))
	}
	ctx.HTML(http.StatusOK, "modelo.html", gin.H{
		"modelo": modelo,
		"msg":    "Atualizado...OK",
	})

}

type requestFindModelos struct {
	Nome   string `form:"nome" query:"nome"`
	Limit  int32  `form:"limit" query:"limit"`
	Offset int32  `form:"offset" query:"offset"`
}

func (server *Server) findModelo(c *gin.Context) {
	var req requestFindModelos
	err := c.ShouldBind(&req)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", errorResponse(err))
	}
	log.Println("-------- Log --------------")
	log.Println(req.Nome)
	log.Println(req.Limit)
	log.Println(req.Offset)
	log.Print("-----------log---------------")
	if req.Limit < 10 {
		req.Limit = 10
	}

	arg := modelo.FindModelosByNameParams{
		Nome:   req.Nome + "%",
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	modelos, err := server.store.FindModelosByName(c, arg)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", errorResponse(err))
	} else {
		tam := len(modelos)
		c.HTML(http.StatusOK, "modelo.html", gin.H{
			"modelos": modelos,
			"limit":   req.Limit,
			"offset":  req.Offset + req.Limit,
			"nome":    req.Nome,
			"msg":     "Retorno:" + strconv.Itoa(tam) + " Itens",
		})
	}
}

type jsonPesquisaMarcas struct {
	Nome string `json:"nome"`
}

func (server *Server) findMarcasForModeloJson(ctx *gin.Context) {
	var pesq jsonPesquisaMarcas
	log.Println("Json chamou.........")
	err := ctx.BindJSON(pesq)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", errorResponse(err))
	}
	nome := pesq.Nome
	marcas, err := server.store.FindMarcasByName(ctx, nome+"%")
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", errorResponse(err))
	}
	ctx.HTML(http.StatusOK, "addModelo.html", gin.H{
		"marcas": marcas,
		"msg":    "Retorno:" + strconv.Itoa(len(marcas)) + " Itens",
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, marcas)

}

func (server *Server) findMarcasForModelo(ctx *gin.Context) {
	nome := ctx.PostForm("nome")
	marcas, err := server.store.FindMarcasByName(ctx, nome+"%")
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", errorResponse(err))
	}
	ctx.HTML(http.StatusOK, "addModelo.html", gin.H{
		"marcas": marcas,
		"msg":    "Retorno:" + strconv.Itoa(len(marcas)) + " Itens",
	})

}
