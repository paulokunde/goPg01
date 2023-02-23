package api

import (
	"github.com/gin-gonic/gin"

	"app/modelo"
)

type Server struct {
	store  *modelo.ExecuteStore
	router *gin.Engine
}

func InstanceServer(store *modelo.ExecuteStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("static/**/*.html")

	//login
	router.GET("/", server.Login)
	//autenticação
	router.GET("/autentica", server.Autentica)

	g1 := router.Group("/modelo")
	{
		g1.GET("/master", server.formModelo)
		g1.POST("/listMarcas", server.findMarcasForModelo)
		g1.GET("/add", server.addModelo)
		g1.POST("/create", server.createModelo)
		g1.POST("/update/", server.updateModelo)
		g1.POST("/find", server.findModelo)
		g1.GET("/find", server.findModelo)
		g1.GET("/delete/:id", server.deleteModelo)
	}

	g2 := router.Group("/marca")
	{
		g2.GET("/master", server.formMarca)
		g2.GET("/add", server.addMarca)
		g2.POST("/create", server.createMarca)
		g2.POST("/update/", server.updateMarca)
		g2.POST("/find", server.findMarcaByName)
		g2.GET("/find", server.findMarcaByName)
		g2.GET("/delete/:id", server.deleteMarca)
	}

	g3 := router.Group("/fornecedor")
	{
		//forms
		g3.GET("/lista", server.formListaFornecedor)
		g3.GET("/form", server.formFornecedor)
		//funçoes
		g3.POST("/search", server.searchFornecedor)
		g3.POST("/create", server.createFornecedor)
		g3.POST("/update", server.updateFornecedor)
		g3.POST("/find", server.findFornecedor)
		g3.POST("/delete", server.deleteFornecedor)
	}

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"erro": err.Error()}
}
func ErrorResponse(err error) gin.H {
	return gin.H{"erro": err.Error()}
}
