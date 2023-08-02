package cli

import (
	"fmt"

	"github.com/simpletextbr/fullcycle-ports-and-adapters/application"
)

func Run(service application.IProductService, action string, productId string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

		return result, nil
	case "enable":
		product, err := service.GetByID(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Enable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been enabled", product.GetName())
		return result, nil

	case "disable":
		product, err := service.GetByID(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Disable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been disabled", product.GetName())
		return result, nil

	default:
		product, err := service.GetByID(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s\n with the name %s\n with the price %f\n and status %s\n has been found",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

		return result, nil
	}
}
