package handlers

import "github.com/benja-M-1/thegoodcorner/models"

type mockDB struct{}

func (mDB *mockDB) AllStatistics() {
	panic("implement me")
}

func (mDB *mockDB) CreateRequest(r *models.Request) (*models.Request, error) {
	r.Id = 1
	return r, nil
}


