package service

import (
	"encoding/json"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

// @Title listProduct
// @Description list all User
// @Router /users [get]
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
