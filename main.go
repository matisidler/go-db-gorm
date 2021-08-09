package main

import (
	"github.com/matisidler/go-db-gorm.git/model"
	"github.com/matisidler/go-db-gorm.git/storage"
)

func main() {
	driver := storage.MySQL // driver := storage.MySQL
	storage.NewConnection(driver)
	/* storage.Pool().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{}) */
	/*
		product1 := model.Product{
			Name:  "Curso de JS",
			Price: 85,
		}
		storage.DB().Create(&product1)

			product2 := model.Product{
				Name:          "Curso de Testing",
				Price:         150,
				Observaciones: storage.StringToNull("Probando Nulos 2"),
			}
			storage.Pool().Create(&product2)

			product3 := model.Product{
				Name:          "Curso de Python",
				Price:         200,
				Observaciones: storage.StringToNull("Probando Nulos"),
			}
			storage.Pool().Create(&product3) */

	/* products := make([]model.Product, 0)
	storage.Pool().Find(&products)
	for _, product := range products {
		fmt.Println(product.ID, product.Name, product.CreatedAt, product.UpdatedAt, product.Price, product.Observaciones.String)
	} */

	/* myProduct := model.Product{}
	storage.Pool().First(&myProduct, 2)
	fmt.Println(myProduct.ID, myProduct.Name, myProduct.Observaciones.String) */

	/* 	products := make([]model.Product, 0)
	   	storage.DB().Find(&products)
	   	for _, product := range products {
	   		product.Name = "Probando"
	   		storage.DB().Save(&product)
	   	}
	*/

	/* 	myProduct := model.Product{}
	   	myProduct.ID = 2
	   	storage.DB().Unscoped().Delete(&myProduct) */

	invoice := model.InvoiceHeader{
		Client: "Matias Sidler",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 1},
			{ProductID: 4},
		},
	}
	storage.DB().Create(&invoice)

}
