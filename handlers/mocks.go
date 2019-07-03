package handlers

import "github.com/benja-M-1/thegoodcorner/models"

type mockDB struct{}

func (mDB *mockDB) AllStatistics() ([]*models.Statistic, error) {
	return nil, nil
}

func (mDB *mockDB) CreateRequest(r *models.Request) (*models.Request, error) {
	r.Id = 1
	return r, nil
}


