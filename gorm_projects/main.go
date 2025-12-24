package main

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})

	// Read
	product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx)       // find product with integer primary key
	products, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx) // find product with code D42
	_ = products

	// Update - update product's price to 200
	_, err = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
	// Update - update multiple fields
	_, err = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, Product{Code: "D42", Price: 100})

	// Delete - delete product
	_, err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
}
