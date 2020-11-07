package service

import (
	"encoding/json"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

// @Title listProduct
// @Description list single Product
// @Router /product [get]
// @Accept  json
// @Success 200 {object}
// @Failure 400 {object}
func listProductHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		product := deps.Store.ListProduct(req.Context())

		respBytes, err := json.Marshal(product)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling product data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.Write(respBytes)
	})
}

// @Title createProduct
// @Description create single Product
// @Router /product [post]
// @Accept  json
// @Success 200 {object}
// @Failure 400 {object}
func createProductHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		deps.Store.CreateProduct(req.Context())

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
	})
}
