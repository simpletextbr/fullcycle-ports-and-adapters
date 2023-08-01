package application

type ProductService struct {
	Persistence IProductPersistence
}

// func (s *ProductService) GetAll() ([]IProduct, error) {
// 	return s.Persistence.GetAll()
// }

func (s *ProductService) GetByID(id string) (IProduct, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
