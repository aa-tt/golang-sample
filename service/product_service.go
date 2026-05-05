package service

import (
	"fmt"
	"microservice/models"
	"sync"
)

var (
	MessageBus = make(chan models.Product, 10)
	wg         sync.WaitGroup
)

func StartProducer(product models.Product) {
	fmt.Printf("Producer: Sending product %s to bus\n", product.Name)
	MessageBus <- product
}

func StartConsumer() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for product := range MessageBus {
			fmt.Printf("Consumer: Received product %s, saving to DB\n", product.Name)
			err := SaveProduct(product)
			if err != nil {
				fmt.Printf("Consumer Error: %v\n", err)
			}
		}
	}()
}

func SaveProduct(p models.Product) error {
	_, err := models.DB.Exec("INSERT INTO products (id, name, price) VALUES (?, ?, ?)", p.ID, p.Name, p.Price)
	return err
}

func GetProducts() ([]models.Product, error) {
	rows, err := models.DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
