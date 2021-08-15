package repository

import (
	"fmt"

	"github.com/go-zelus/api-layout/model"
	"github.com/go-zelus/zelus/core/stores/mysql"
	"github.com/jmoiron/sqlx"
)

type AreaRepo interface {
	Provinces() ([]model.Area, error)
	Cities(code int64) ([]model.Area, error)
	District(code int64) ([]model.Area, error)
	Streets(code int64) ([]model.Area, error)
	Committee(code int64) ([]model.Area, error)
}

func NewAreaRepo() AreaRepo {
	return &areaRepo{
		db:    mysql.DB(),
		table: "area",
	}
}

type areaRepo struct {
	db    *sqlx.DB
	table string
}

func (a areaRepo) Provinces() ([]model.Area, error) {
	query := fmt.Sprintf("select %s from %s where level=0",
		mysql.FieldRows(model.AreaFields, ","),
		a.table,
	)
	areas := make([]model.Area, 0)
	err := a.db.Select(&areas, query)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

func (a areaRepo) Cities(code int64) ([]model.Area, error) {
	query := fmt.Sprintf("select %s from %s where level=1 and parent_code=?",
		mysql.FieldRows(model.AreaFields, ","),
		a.table,
	)
	areas := make([]model.Area, 0)
	err := a.db.Select(&areas, query, code)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

func (a areaRepo) District(code int64) ([]model.Area, error) {
	query := fmt.Sprintf("select %s from %s where level=2 and parent_code=?",
		mysql.FieldRows(model.AreaFields, ","),
		a.table,
	)
	areas := make([]model.Area, 0)
	err := a.db.Select(&areas, query, code)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

func (a areaRepo) Streets(code int64) ([]model.Area, error) {
	query := fmt.Sprintf("select %s from %s where level=3 and parent_code=?",
		mysql.FieldRows(model.AreaFields, ","),
		a.table,
	)
	areas := make([]model.Area, 0)
	err := a.db.Select(&areas, query, code)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

func (a areaRepo) Committee(code int64) ([]model.Area, error) {
	query := fmt.Sprintf("select %s from %s where level=4 and parent_code=?",
		mysql.FieldRows(model.AreaFields, ","),
		a.table,
	)
	areas := make([]model.Area, 0)
	err := a.db.Select(&areas, query, code)
	if err != nil {
		return nil, err
	}
	return areas, nil
}
