package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AdiLambe/TestGoLangCodes/workspace/dto"
	"github.com/AdiLambe/TestGoLangCodes/workspace/errs"
	"github.com/AdiLambe/TestGoLangCodes/workspace/mocks/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

// PASS
func Test_Orders_status_code_200(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dummyOrders := []dto.OrderResponse{
		{Id: "1002", Name: "Iphone", Number: "P780", Description: "Mobile Phone", Status: "1"},
		{Id: "1004", Name: "lenovo", Number: "P423", Description: "laptop", Status: "0"},
	}
	mockService := service.NewMockOrderService(ctrl)
	mockService.EXPECT().GetOrdersList("").Return(dummyOrders, nil)
	ch := OrderHandlers{mockService}

	router := gin.Default()
	router.GET("/orders", ch.GetOrdersList)

	request, _ := http.NewRequest(http.MethodGet, "/orders", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

// PASS
func Test_Should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockOrderService(ctrl)
	mockService.EXPECT().GetOrdersList("").Return(nil, errs.NewUnexpectedError("some database error"))

	ch := OrderHandlers{mockService}
	// Act
	router := gin.Default()
	router.GET("/orders", ch.GetOrdersList)

	request, _ := http.NewRequest(http.MethodGet, "/orders", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}

}

// PASS
func Test_GetOrder_Success(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockOrderService(ctrl)
	ch := OrderHandlers{service: mockService}

	router := gin.Default()
	router.GET("/orders/:order_id", ch.GetOrder)

	expectedOrder := &dto.OrderResponse{
		Id:          "1009",
		Name:        "Boat",
		Number:      "P232",
		Description: "Headphones",
		Status:      "1",
	}
	// Print the expected call to help with debugging
	t.Logf("Expecting GetOrder(1009)")
	mockService.EXPECT().GetOrder(gomock.Any()).Return(expectedOrder, nil)

	// Update the request URL to include the order ID
	request, _ := http.NewRequest(http.MethodGet, "/orders/1009", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

// PASS
func Test_GetOrder_status_code_500(t *testing.T) {
	// Arrange

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockOrderService(ctrl)
	ch := OrderHandlers{mockService}
	mockService.EXPECT().GetOrder(gomock.Any()).Return(nil, errs.NewUnexpectedError("some database error"))

	// Act
	router := gin.Default()
	router.GET("/orders/:order_id", ch.GetOrder)

	request, _ := http.NewRequest(http.MethodGet, "/orders/1009", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}

}

func Test_PostOrder_Success(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockOrderService(ctrl)
	ch := OrderHandlers{service: mockService}

	router := gin.Default()
	router.POST("/postOrder", ch.CreateOrder)

	expectedOrder := &dto.OrderResponse{
		Id:          "1002",
		Name:        "Apple",
		Number:      "P232",
		Description: "Mobile phone",
		Status:      "1",
	}

	mockService.EXPECT().CreateOrder(gomock.Any()).Return(expectedOrder, nil)

	expectedRequestBody := `{"Id":"1002","Name":"Apple","Number":"P232","Description":"Mobile phone","Status":"1"}`
	request, _ := http.NewRequest(http.MethodPost, "/postOrder", strings.NewReader(expectedRequestBody))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusCreated {
		t.Error("Failed while testing the status code")
	}
}
