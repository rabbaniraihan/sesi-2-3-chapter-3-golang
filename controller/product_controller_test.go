package controller

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestProductController_GetAllProduct(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		pc   *ProductController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pc.GetAllProduct(tt.args.ctx)
		})
	}
}

func TestProductController_GetProductById(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		pc   *ProductController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pc.GetProductById(tt.args.ctx)
		})
	}
}
