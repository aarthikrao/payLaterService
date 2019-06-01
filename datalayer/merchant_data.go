package datalayer

import (
	conn "github.com/aarthikrao/payLaterService/connections"
	"github.com/aarthikrao/payLaterService/models"
)

// MerchantData will be used to fetch data
type MerchantData struct{}

// NewMerchantData is used to create a new instance of Merchnatdata
func NewMerchantData() *MerchantData {
	return &MerchantData{}
}

// AddNewMerchant is used to add new users to the system
func (md *MerchantData) AddNewMerchant(merchant models.Merchant) (err error) {
	err = conn.PGDB.Insert(&merchant)
	return
}

// GetMerchantByName is used to fetch merchant by name
func (md *MerchantData) GetMerchantByName(name string) (merchant models.Merchant, err error) {
	err = conn.PGDB.Model(&merchant).Where("name = ?", name).Select()
	return
}

// UpdateMerchantByName is used to update the merchant data by name
func (md *MerchantData) UpdateMerchantByName(name string, merchant models.Merchant) (err error) {
	_, err = conn.PGDB.Model(&merchant).Where("name = ? ", name).UpdateNotNull(merchant)
	return
}

// GetAllMerchantData is used to get all merchant related data
func (md *MerchantData) GetAllMerchantData() (merchants []models.Merchant, err error) {
	err = conn.PGDB.Model(&merchants).Select()
	return
}
