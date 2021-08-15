package service

import (
	"context"

	"github.com/go-zelus/api-layout/entity"
	"github.com/go-zelus/api-layout/repository"
)

type AreaService interface {
	Provinces(ctx context.Context) ([]entity.Area, error)
	Cities(ctx context.Context, code int64) ([]entity.Area, error)
	Districts(ctx context.Context, code int64) ([]entity.Area, error)
	Streets(ctx context.Context, code int64) ([]entity.Area, error)
	Committees(ctx context.Context, code int64) ([]entity.Area, error)
}

func NewAreaService() AreaService {
	return &areaService{
		areaRepo: repository.NewAreaRepo(),
	}
}

type areaService struct {
	areaRepo repository.AreaRepo
}

func (a areaService) Provinces(ctx context.Context) ([]entity.Area, error) {
	res, err := a.areaRepo.Provinces()
	if err != nil {
		return nil, err
	}
	out := make([]entity.Area, 0)
	for _, area := range res {
		out = append(out, entity.Area{
			AreaCode:   area.AreaCode,
			ParentCode: area.ParentCode,
			Level:      area.Level,
			Name:       area.Name,
			ShortName:  area.ShortName,
			MergerName: area.MergerName,
			Pinyin:     area.Pinyin,
			ZipCode:    area.ZipCode,
			CityCode:   area.CityCode,
			Lng:        area.Lng,
			Lat:        area.Lat,
		})
	}
	return out, nil
}

func (a areaService) Cities(ctx context.Context, code int64) ([]entity.Area, error) {
	res, err := a.areaRepo.Cities(code)
	if err != nil {
		return nil, err
	}
	out := make([]entity.Area, 0)
	for _, area := range res {
		out = append(out, entity.Area{
			AreaCode:   area.AreaCode,
			ParentCode: area.ParentCode,
			Level:      area.Level,
			Name:       area.Name,
			ShortName:  area.ShortName,
			MergerName: area.MergerName,
			Pinyin:     area.Pinyin,
			ZipCode:    area.ZipCode,
			CityCode:   area.CityCode,
			Lng:        area.Lng,
			Lat:        area.Lat,
		})
	}
	return out, nil
}

func (a areaService) Districts(ctx context.Context, code int64) ([]entity.Area, error) {
	res, err := a.areaRepo.District(code)
	if err != nil {
		return nil, err
	}
	out := make([]entity.Area, 0)
	for _, area := range res {
		out = append(out, entity.Area{
			AreaCode:   area.AreaCode,
			ParentCode: area.ParentCode,
			Level:      area.Level,
			Name:       area.Name,
			ShortName:  area.ShortName,
			MergerName: area.MergerName,
			Pinyin:     area.Pinyin,
			ZipCode:    area.ZipCode,
			CityCode:   area.CityCode,
			Lng:        area.Lng,
			Lat:        area.Lat,
		})
	}
	return out, nil
}

func (a areaService) Streets(ctx context.Context, code int64) ([]entity.Area, error) {
	res, err := a.areaRepo.Streets(code)
	if err != nil {
		return nil, err
	}
	out := make([]entity.Area, 0)
	for _, area := range res {
		out = append(out, entity.Area{
			AreaCode:   area.AreaCode,
			ParentCode: area.ParentCode,
			Level:      area.Level,
			Name:       area.Name,
			ShortName:  area.ShortName,
			MergerName: area.MergerName,
			Pinyin:     area.Pinyin,
			ZipCode:    area.ZipCode,
			CityCode:   area.CityCode,
			Lng:        area.Lng,
			Lat:        area.Lat,
		})
	}
	return out, nil
}

func (a areaService) Committees(ctx context.Context, code int64) ([]entity.Area, error) {
	res, err := a.areaRepo.Committee(code)
	if err != nil {
		return nil, err
	}
	out := make([]entity.Area, 0)
	for _, area := range res {
		out = append(out, entity.Area{
			AreaCode:   area.AreaCode,
			ParentCode: area.ParentCode,
			Level:      area.Level,
			Name:       area.Name,
			ShortName:  area.ShortName,
			MergerName: area.MergerName,
			Pinyin:     area.Pinyin,
			ZipCode:    area.ZipCode,
			CityCode:   area.CityCode,
			Lng:        area.Lng,
			Lat:        area.Lat,
		})
	}
	return out, nil
}
