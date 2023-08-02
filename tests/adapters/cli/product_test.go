package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simpletextbr/fullcycle-ports-and-adapters/adapters/cli"
	mock_application "github.com/simpletextbr/fullcycle-ports-and-adapters/tests/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 10.30
	productStatus := "disabled"
	productId := "abc"

	productMock := mock_application.NewMockIProduct(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockIProductService(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().GetByID(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	result, err := cli.Run(service, "create", "", productName, productPrice)
	if err != nil {
		t.Errorf("Error during the test: %s", err)
	}

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", productId, productName, productPrice, productStatus)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(service, "enable", productId, "", 0)
	if err != nil {
		t.Errorf("Error during the test: %s", err)
	}

	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(service, "disable", productId, "", 0)
	if err != nil {
		t.Errorf("Error during the test: %s", err)
	}

	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(service, "get", productId, "", 0)
	if err != nil {
		t.Errorf("Error during the test: %s", err)
	}

	resultExpected = fmt.Sprintf("Product ID %s\n with the name %s\n with the price %f\n and status %s\n has been found", productId, productName, productPrice, productStatus)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
