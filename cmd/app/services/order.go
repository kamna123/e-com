package services

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"e-commerce/cmd/app/models"
	"e-commerce/cmd/app/schema"

	"github.com/golang/glog"

	"e-commerce/cmd/app/repositories"
)

type IOrderSerivce interface {
	GetOrders(ctx context.Context, query *schema.OrderQueryParam) (*[]models.Order, error)
	GetOrderByID(ctx context.Context, uuid string) (*models.Order, error)
	CreateOrder(ctx context.Context, item *schema.OrderBodyParam) (*models.Order, error)
	RazorPayOrder(ctx context.Context, item *schema.RazorPayOrderParam) (*schema.RazorPayResp, error)
}

type order struct {
	repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) IOrderSerivce {
	return &order{repo: repo}
}

func (categ *order) RazorPayOrder(ctx context.Context, item *schema.RazorPayOrderParam) (*schema.RazorPayResp, error) {

	payloadBytes, err := json.Marshal(item)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.razorpay.com/v1/orders", body)
	if err != nil {
		// handle err
	}
	//req.CreateHeaderList("Accept", "application/json", "Accept-Language", "en_US", "Authorization", ID+":"+SECRET)
	// TestClientID := "rzp_test_GoCENyxwaOhUgI"
	// TestSecret := "qyZPkKiuCEgkeSwVjqP0sq9u"
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", fmt.Sprintf("Basic %s:%s", TestClientID, TestSecret))
	req.SetBasicAuth("rzp_test_GoCENyxwaOhUgI", "qyZPkKiuCEgkeSwVjqP0sq9u")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	// Read the response body using ioutil
	respBody, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	response := schema.RazorPayResp{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		glog.Error("err is ", err)
	}
	glog.Info("res  ", response)
	return &response, nil
}
func (categ *order) GetOrders(ctx context.Context, query *schema.OrderQueryParam) (*[]models.Order, error) {
	orders, err := categ.repo.GetOrders(query)
	if err != nil {
		return nil, err
	}

	return orders, err
}

func (categ *order) GetOrderByID(ctx context.Context, uuid string) (*models.Order, error) {
	order, err := categ.repo.GetOrderByID(uuid)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (categ *order) CreateOrder(ctx context.Context, item *schema.OrderBodyParam) (*models.Order, error) {
	order, err := categ.repo.CreateOrder(item)
	if err != nil {
		return nil, err
	}

	return order, nil
}
