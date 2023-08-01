package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simpletextbr/fullcycle-ports-and-adapters/application"
	mock_application "github.com/simpletextbr/fullcycle-ports-and-adapters/tests/mocks"
	"github.com/stretchr/testify/require"
)

func TestApplicationProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockIProduct(ctrl)
	persistence := mock_application.NewMockIProductPersistence(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.GetByID("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}
