package usecase

import (
	"context"
	"marine/internal/models"
)

type MaterialService interface {
	UploadMaterial(ctx context.Context, file *models.Material) (int, error)
	GetAllMaterials(ctx context.Context) ([]*models.Material, error)
}

type materialService struct {
	materials map[int]*models.Material // Хранение материалов в памяти (для примера)
}

func NewMaterialService() MaterialService {
	return &materialService{
		materials: make(map[int]*models.Material),
	}
}

func (s *materialService) UploadMaterial(ctx context.Context, material *models.Material) (int, error) {
	material.ID = len(s.materials) + 1
	s.materials[material.ID] = material
	return material.ID, nil
}

func (s *materialService) GetAllMaterials(ctx context.Context) ([]*models.Material, error) {
	var materialList []*models.Material
	for _, material := range s.materials {
		materialList = append(materialList, material)
	}
	return materialList, nil
}
