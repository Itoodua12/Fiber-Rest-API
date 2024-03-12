package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/itoodua12/Fiber-Rest-API/database"
	"github.com/itoodua12/Fiber-Rest-API/model"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel model.Product) Product {
	return Product{
		ID:           productModel.ID,
		Name:         productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product model.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	products := []model.Product{}

	database.Database.DB.Find(&products)
	responseProducts := []Product{}
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}

func findProduct(id int, product *model.Product) error {
	database.Database.DB.Find(&product, "id=?", id)
	if product.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}


func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product model.Product

	if err != nil {
		return c.Status(400).JSON("Bad Request")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}


func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product model.Product

	if err != nil {
		return c.Status(400).JSON("Bad Request")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name string `json:"name"`
		SerialNumber  string `json:"serial_number"`
	}
	var updateData UpdateProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.SerialNumber = updateData.SerialNumber
	product.Name = updateData.Name
	database.Database.DB.Save(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}


func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product model.Product

	if err != nil {
		return c.Status(400).JSON("Bad Request")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.Database.DB.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)

}