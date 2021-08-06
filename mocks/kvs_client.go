package mocks

import (
	"fmt"

	"github.com/mercadolibre/go-meli-toolkit/gokvsclient"
)

type kvsClient map[string]gokvsclient.Item

func NewKvsClient() gokvsclient.Client {
	return &kvsClient{}
}

func (kc kvsClient) DeleteContainer() error {
	return nil
}

func (kc kvsClient) Get(key string) (gokvsclient.Item, error) {
	return kc[key], nil
}

func (kc kvsClient) Save(item gokvsclient.Item) error {
	kc[item.GetKey()] = item
	return nil
}

func (kc kvsClient) Update(item gokvsclient.Item) error {
	kc[item.GetKey()] = item
	return nil
}

func (kc kvsClient) Delete(key string) error {
	_ = fmt.Sprint(key)
	return nil
}

func (kc kvsClient) BatchGet(keys []string) ([]gokvsclient.Item, error) {
	_ = fmt.Sprint(keys)
	return nil, nil
}

func (kc kvsClient) BatchSave(items []gokvsclient.Item) error {
	_ = fmt.Sprint(items)
	return nil
}

func (kc kvsClient) BatchUpdate(items []gokvsclient.Item) error {
	_ = fmt.Sprint(items)
	return nil
}

func (kc kvsClient) BatchDelete(keys []string) error {
	_ = fmt.Sprint(keys)
	return nil
}

func (kc kvsClient) BulkGet(keys []string) ([]gokvsclient.BulkItem, error) {
	_ = fmt.Sprint(keys)
	return nil, nil
}

func (kc kvsClient) BulkSave(items []gokvsclient.Item) ([]gokvsclient.BulkItem, error) {
	_ = fmt.Sprint(items)
	return nil, nil
}

func (kc kvsClient) BulkUpdate(items []gokvsclient.Item) ([]gokvsclient.BulkItem, error) {
	_ = fmt.Sprint(items)
	return nil, nil
}

func (kc kvsClient) BulkDelete(keys []string) ([]gokvsclient.BulkItem, error) {
	_ = fmt.Sprint(keys)
	return nil, nil
}

func (kc kvsClient) GetCounter(key string) (*int64, error) {
	_ = fmt.Sprint(key)
	return nil, nil
}

func (kc kvsClient) IncrementCounter(key string, value *int64) error {
	_ = fmt.Sprint(key, value)
	return nil
}

func (kc kvsClient) DecrementCounter(key string, value *int64) error {
	_ = fmt.Sprint(key, value)
	return nil
}

func (kc kvsClient) ResetCounter(key string) error {
	_ = fmt.Sprint(key)
	return nil
}
