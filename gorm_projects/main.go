package main

import (
	"context"
	"fmt"
	"gorm_projects/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	ctx := context.Background()
	gorm.G[models.User](db, clause.OnConflict{DoNothing: true}).Create(ctx, &models.User{Name: "Alice"})
	gorm.G[models.User](db).CreateInBatches(ctx, &[]models.User{{Name: "joins"}, {Name: "denos"}, {Name: "joins1"}, {Name: "joins2"}, {Name: "joins3"}}, 10)
	//Query records
	user, err := gorm.G[models.User](db).Where("name = ?", "joins1").First(ctx)
	fmt.Println(user.Name)

	_, err = gorm.G[models.User](db).Where("id = ?", 1).Update(ctx, "name", "newalice")
	if err != nil {
		fmt.Println(err)
	}

	gorm.G[models.User](db).Where("id = ?", 2).Updates(ctx, models.User{Name: "wjuidnew", Age: 18})
	user1, err := gorm.G[models.User](db).Where("id <= ?", 10).Find(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(user1)

	for _, u := range user1 {
		gorm.G[models.User](db).Where("id = ?", u.ID).Delete(ctx)
	}
	fmt.Println(user1)
	var res []models.User
	if err := db.Find(&res).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
