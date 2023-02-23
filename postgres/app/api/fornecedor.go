package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	jsonResp, er := json.Marshal(lista)
	if er != nil {
		panic(er)
	}

	c.Writer.WriteHeader(200)
	c.Writer.Write(jsonResp)

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

}
func (server *Server) deleteFornecedor(c *gin.Context) {

}

func (server *Server) findFornecedor(c *gin.Context) {

}
