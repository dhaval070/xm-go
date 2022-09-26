package seeder

import (
	"log"
	"xmgo/pkg/models"
	"xmgo/pkg/repo"
)

var companies = []models.Company{
	{
		Name:    "Rahul Patel",
		Country: "India",
		Phone:   "2234432345",
		Website: "http://google.com",
	},
	{
		Name:    "Pankaj Pandya",
		Country: "USA",
		Phone:   "4112312222",
		Website: "http://pankaj.com",
	},
	{
		Name:    "John Bond",
		Country: "Canada",
		Phone:   "223522333",
		Website: "http://john.com",
	},
}

func Seed(repository repo.IRepository) {
	for _, rec := range companies {
		if err := repository.CreateCompany(&rec); err != nil {
			log.Println(err)
		}
	}
}
