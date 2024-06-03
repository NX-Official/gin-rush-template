package main

import (
	"gin-rush-template/internal/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/global/query",
		//Mode:    gen.WithoutContext, // generate mode
	})

	//g.ApplyBasic(database.Models...)
	g.ApplyBasic(
		model.User{},
	)

	g.Execute()
}
