package controller

import (
	"net/http"
	"sesi-2/model"
	"sesi-2/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var request model.ProductCreateRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}

	product, err := pc.ProductService.Create(request, userId.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetAllProduct(ctx *gin.Context) {
	products, err := pc.ProductService.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	var getProduct model.Product
	getProduct.Id = id

	err := pc.ProductService.GetById(&getProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get product by id",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, getProduct)
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {

}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {

}
