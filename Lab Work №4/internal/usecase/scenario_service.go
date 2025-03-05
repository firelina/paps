package usecase

import (
	"context"
	"errors"
	"marine/internal/models"
)

var ErrScenarioNotFound = errors.New("scenario not found")

type ScenarioService interface {
	CreateScenario(ctx context.Context, scenario *models.Scenario) (int, error)
	GetAllScenarios(ctx context.Context) ([]*models.Scenario, error)
	DeleteScenario(ctx context.Context, id int) error
	UpdateScenario(ctx context.Context, id int, scenario *models.Scenario) (*models.Scenario, error)
}

type scenarioService struct {
	scenarios map[int]*models.Scenario // Хранение сценариев в памяти (для примера)
}

func NewScenarioService() ScenarioService {
	return &scenarioService{
		scenarios: make(map[int]*models.Scenario),
	}
}

func (s *scenarioService) CreateScenario(ctx context.Context, scenario *models.Scenario) (int, error) {
	scenario.ID = len(s.scenarios) + 1
	s.scenarios[scenario.ID] = scenario
	return scenario.ID, nil
}

func (s *scenarioService) GetAllScenarios(ctx context.Context) ([]*models.Scenario, error) {
	var scenarioList []*models.Scenario
	for _, scenario := range s.scenarios {
		scenarioList = append(scenarioList, scenario)
	}
	return scenarioList, nil
}

func (s *scenarioService) DeleteScenario(ctx context.Context, id int) error {
	if _, exists := s.scenarios[id]; !exists {
		return ErrScenarioNotFound
	}
	delete(s.scenarios, id)
	return nil
}

func (s *scenarioService) UpdateScenario(ctx context.Context, id int, scenario *models.Scenario) (*models.Scenario, error) {
	if _, exists := s.scenarios[id]; !exists {
		return nil, ErrScenarioNotFound
	}
	s.scenarios[id].Title = scenario.Title
	s.scenarios[id].Description = scenario.Description
	return s.scenarios[id], nil
}
