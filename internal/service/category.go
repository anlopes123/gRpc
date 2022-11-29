package service

import {
	"github.com/anlopes/gRpc/internal/pb"
	"github.com/anlopes/gRpc/internal/database"
}


type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	categoryDB database.Category

}

func NewCategoryService(CategoryDB database.Category) *CategoryService {
	return &CategoryService {
		CategoryDB: categoryDB,
	}	
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.categoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}
	return &pb.CategoryResponse {
		Category: categoryResponse,
	}, nil
}