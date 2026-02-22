package graph

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SaiTejaBandamidi/go-graphql-sqlc-api/db"
	"github.com/SaiTejaBandamidi/go-graphql-sqlc-api/graph/model"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB      *sql.DB
	Queries *db.Queries
}

// StudentHouse implements Student.House field resolver - lazy loads house only when requested
func (r *Resolver) StudentHouse(ctx context.Context, obj *model.Student) (*model.House, error) {
	if obj.GetHouseID() == "" {
		return nil, nil
	}

	house, err := r.Queries.GetHouseByID(ctx, obj.GetHouseID())
	if err != nil {
		return nil, fmt.Errorf("failed to get house for student %s: %v", obj.ID, err)
	}

	result := &model.House{
		ID:   house.ID,
		Name: house.FirstName + " " + house.LastName,
	}
	if house.HeadOfHouse.Valid {
		result.HeadOfHouse = &house.HeadOfHouse.String
	}
	return result, nil
}

// Field resolver for House.Students - lazy loads students only when requested
func (r *Resolver) HouseStudents(ctx context.Context, obj *model.House) ([]*model.Student, error) {
	students, err := r.Queries.GetStudentsByHouse(ctx, sql.NullString{
		String: obj.ID,
		Valid:  true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get students for house %s: %v", obj.ID, err)
	}

	var result []*model.Student
	for _, s := range students {
		student := &model.Student{
			ID:   s.ID,
			Name: s.FirstName + " " + s.LastName,
		}
		if s.Age.Valid {
			age := s.Age.Int32
			student.Age = &age
		}
		if s.HouseID.Valid {
			student.SetHouseID(s.HouseID.String)
		}
		result = append(result, student)
	}
	return result, nil
}
