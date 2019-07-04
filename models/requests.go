package models

type FizzbuzzRequest struct {
	Id   int64  `json:"id,omitempty"`
	Int1 int    `json:"int1"`
	Int2 int    `json:"int2"`
	Str1 string `json:"str1"`
	Str2 string `json:"str2"`
}

func (db *DB) CreateRequest(r *FizzbuzzRequest) (*FizzbuzzRequest, error) {
	query := "INSERT INTO fizzbuzz_requests(int1, int2, str1, str2) VALUES ($1, $2, $3, $4) RETURNING id"
	err := db.QueryRow(query, r.Int1, r.Int2, r.Str1, r.Str2).Scan(&r.Id)
	if err != nil {
		return nil, err
	}

	return r, nil
}
