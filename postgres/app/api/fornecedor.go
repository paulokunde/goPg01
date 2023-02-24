package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"app/modelo"
)

func (server *Server) formListaFornecedor(c *gin.Context) {
	c.HTML(
		http.StatusOK, "listaFornecedor.html", gin.H{
			"title": "Fornecedor",
		},
	)
}
func (server *Server) formFornecedor(c *gin.Context) {
	log.Println("Abrindo formulario fornecedor........")
	c.HTML(
		http.StatusOK, "formFornecedor.html", gin.H{
			"title": "Fornecedor",
		},
	)
}

type NomeParam struct {
	Nome string `json:"nome" form:"nome" query:"nome"`
	ID   int32  `json:"id" form:"id" query:"id" uri:"id"`
}

func (server *Server) searchFornecedor(c *gin.Context) {
	var param NomeParam
	err := c.Bind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Erro ao Pesquisar Fornecedores",
		})
	}
	lista, erro := server.store.FindFornecedorByName(c, param.Nome+"%")
	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Erro ao Listar Fornecedores da Base",
		})
	}

	/*
		jsonResp, er := json.Marshal(lista)
		if er != nil {
			panic(er)
		}

		c.Writer.WriteHeader(200)
		c.Writer.Write(jsonResp)
	*/
	tam := len(lista)
	c.HTML(http.StatusOK, "listaFornecedor.html", gin.H{
		"lista": lista,
		"msg":   "Retorno:" + strconv.Itoa(tam) + " Itens",
		"nome":  param.Nome,
	})
}

func (server *Server) createFornecedor(c *gin.Context) {
	var req NomeParam
	log.Println("Chamanda a createForecedor...")

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	log.Println("Nome:" + req.Nome)
	marca, erro := server.store.CreateFornecedor(c, req.Nome)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(erro))
	}
	resp, er := json.Marshal(marca)
	if er != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(erro))
	}
	c.Writer.WriteHeader(200)
	c.Writer.Write(resp)

}

func (server *Server) updateFornecedor(c *gin.Context) {
	var req NomeParam

	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	log.Println("Nome:" + req.Nome)

	param := modelo.UpdateFornecedorParams{
		ID:   req.ID,
		Nome: req.Nome,
	}
	obj, erro := server.store.UpdateFornecedor(c, param)
	if erro != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(erro))
	}
	/*
		resp, er := json.Marshal(obj)
		if er != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(erro))
		}
		c.Writer.WriteHeader(200)
		c.Writer.Write(resp)
	*/

	c.HTML(http.StatusOK, "listaFornecedor.html", gin.H{
		"fornecedor": obj,
		"msg":        "Objeto Atualizado...: id:" + strconv.Itoa(int(obj.ID)) + " Nome:" + obj.Nome,
		"nome":       obj.Nome,
	})

}
func (server *Server) deleteFornecedor(c *gin.Context) {
	var req NomeParam
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", errorResponse(err))
	}

	er := server.store.DeleteFornecedor(c, req.ID)
	if er != nil {
		c.HTML(http.StatusOK, "error.html", errorResponse(er))
	}
	c.HTML(http.StatusOK, "listaFornecedor.html", gin.H{
		"msg": "Registro Deletado - OK",
	})

}

func (server *Server) findFornecedor(c *gin.Context) {

}
