package main

import "fmt"

var storeItems = make(map[string]products)

var soldItems = make(map[string]products)

type store struct {
	supplierName string
	products
}

type products struct {
	car
	// motocycle
}

type car struct {
	make            string
	model           string
	transmission    string
	fuelType        string
	price           float32
	quantityInStock int
}

type motocycle struct {
	make            string
	model           string
	price           float32
	quantityInStock int
}

func (p *products) displayProduct() {
	fmt.Println(p)
}

func (p *products) productStatus() {

	//display car products
	if p.car.quantityInStock > 0 {
		fmt.Printf("We have %v %v %v's available\n", p.car.quantityInStock, p.car.make, p.car.model)
	} else if p.car.quantityInStock == 0 {
		fmt.Printf("There are no %v %vs available\n", p.car.make, p.car.model)
	}

	// display motocycle products
	// if p.motocycle.quantityInStock > 0 {
	// 	fmt.Printf("We have %v %v %v's available\n", p.motocycle.quantityInStock, p.motocycle.make, p.motocycle.model)
	// } else if p.motocycle.quantityInStock == 0 {
	// 	fmt.Printf("There are no %v %vs available\n", p.motocycle.make, p.motocycle.model)
	// }

}

func (s *store) up4Sale() {
	sum := 0
	carQuantity := 0
	motocycleQuantity := 0

	for _, v := range storeItems {
		quantity := v.car.quantityInStock //+ v.motocycle.inStock
		carQuantity += v.car.quantityInStock
		// motocycleQuantity += v.motocycle.inStock
		sum += quantity
	}
	fmt.Printf("There are %v products available for sale. %v cars and %v motocycles\n", sum, carQuantity, motocycleQuantity)
}

func (s *store) addItem(productMakeAndModel string) {

	storeItems[productMakeAndModel] = s.products
}

func (s *store) listProducts() {
	fmt.Println("Here's the list of all products in the store:")
	for _, v := range storeItems {
		fmt.Println(v)
	}
}

func (s *store) sellItem(productMakeAndModel string, quantity int) {

	if quantity <= s.quantityInStock {
		newQuantity := s.quantityInStock - quantity
		storeItems[productMakeAndModel] = products{
			car{
				s.make,
				s.model,
				s.transmission,
				s.fuelType,
				s.price,
				newQuantity,
			},
			// motocycle{},
		}

		soldItems[productMakeAndModel] = products{
			car{
				s.make,
				s.model,
				s.transmission,
				s.fuelType,
				s.price,
				quantity,
			},
			// motocycle{},
		}
		// }else if quantity <= s.motocycle.quantityInStock {
		// newQuantity := s.motocycle.quantityInStock - quantity
		// storeItems[productMakeAndModel] = products{
		// 	car{},
		// 	motocycle{
		// 		s.make,
		// 		s.model,
		// 		s.transmission,
		// 		s.fuelType,
		// 		s.price,
		// 		newQuantity,
		// 	},
		// }

		// soldItems[productMakeAndModel] = products{
		// 	car{},
		// 	motocycle{
		// 		s.make,
		// 		s.model,
		// 		s.transmission,
		// 		s.fuelType,
		// 		s.price,
		// 		quantity,
		// 	},
		// }
	} else {
		fmt.Printf("Not enough products. Only %v %v %v's available\n", s.quantityInStock, s.make, s.model)
	}
}

func (s *store) listSoldItems() {
	var totalPrice float32
	fmt.Println("Here's the list of all sold items:")

	for _, v := range soldItems {
		fmt.Println(v)
		price := v.price * float32(v.quantityInStock)
		totalPrice += price
	}
	fmt.Printf("The total price of these items is $%v\n", totalPrice)
}

func main() {

	firstBatch := &store{
		supplierName: "Joseph Eziwanne",
		products: products{
			car: car{
				"Toyota",
				"Corolla",
				"automatic",
				"petrol",
				8227.85,
				7,
			},
			// motocycle: motocycle{
			// 	"Bajaj",
			// 	"Pulsar NS160",
			// 	437.5,
			// 	10,
			// },
		},
	}

	secondBatch := &store{
		"Dele Bishop",
		products{
			car{
				"Ford",
				"Explorer",
				"automatic",
				"petrol",
				15063.3,
				3,
			},
			// motocycle{
			// 	"Qlink",
			// 	"Ranger",
			// 	512.5,
			// 	5,
			// },
		},
	}

	thirdBatch := &store{
		"Akanbi & JOhnsons ltd",
		products{
			car{
				"Honda",
				"CRV",
				"Automatic",
				"Petrol",
				6329.1,
				12,
			},
			// motocycle{
			// 	"Honda",
			// 	"CBR 600cc Power Bike",
			// 	2625,
			// 	2,
			// },
		},
	}

	//*************** DEMO ***************
	// METHODS ON PRODUCT CLASS
	firstBatch.displayProduct()
	// secondBatch.products.displayProduct()
	thirdBatch.displayProduct()

	// firstBatch.productStatus()
	secondBatch.productStatus()
	// thirdBatch.productStatus()

	// METHODS ON STORE CLASS
	//adding items to store
	firstBatch.addItem("Toyota Corolla")
	secondBatch.addItem("Ford Explorer")
	thirdBatch.addItem("Honda CRV")

	//items available for sale
	firstBatch.up4Sale()

	//list all product items
	firstBatch.listProducts()

	//sell an item
	firstBatch.sellItem("Toyota Corolla", 4)
	// secondBatch.sellItem("Ford Explorer", 4)
	thirdBatch.sellItem("Honda CRV", 5)

	//list all sold items and total price
	firstBatch.listSoldItems()
}
