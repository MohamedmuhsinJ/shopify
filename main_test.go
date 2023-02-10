package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MohamedmuhsinJ/shopify/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEditProduct(t *testing.T) {

	r := SetUpRouter()
	r.PUT("/editproduct/:id", controllers.EditProduct)
	product := controllers.Editproduct{
		Price: 13000,
	}
	json, _ := json.Marshal(product)

	// if err != nil {
	// 	t.Fatal(err)
	// }

	id := "2"
	req, _ := http.NewRequest("PUT", "/editproduct/"+id, bytes.NewBuffer(json))
	// if err != nil {
	// 	t.Fatal(err)
	// }
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusAccepted, w.Code)
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
