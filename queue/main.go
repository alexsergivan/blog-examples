package main

import (
	"fmt"
	"github.com/alexsergivan/blog-examples/queue/queue"
)

var products = []string{
	"books",
	"computers",
}

func main()  {
	newProducts := []string{
		"apples",
		"oranges",
		"wine",
		"bread",
		"orange juice",
	}

	productsQueue := queue.NewQueue("NewProducts")
	var jobs []queue.Job

	for _, newProduct := range newProducts {
		product := newProduct
		action := func() error {
			products = append(products, product)
			return nil
		}
		jobs = append(jobs, queue.Job{
			Name:   fmt.Sprintf("Importing new product: %s", newProduct),
			Action: action,
		})
	}

	productsQueue.AddJobs(jobs)

	worker := queue.NewWorker(productsQueue)
	worker.DoWork()
	defer fmt.Print(products)
}

