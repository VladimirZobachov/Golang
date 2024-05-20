package repository

import (
	"hostess-service/internal/model"
)

//type HallService interface {
//	GetHallsMap() (model.HallResponse, error)
//	ImportHallsFromGoulash() model.HallResponse
//}
//
//type hallService struct {
//	db      *gorm.DB
//	apiURL  string
//	apiKEY  string
//	apiUUID string
//}
//
//func NewHallService(db *gorm.DB, api config.Goulash) HallService {
//	return &hallService{db: db, apiURL: api.APIUrl, apiKEY: api.APIKey, apiUUID: api.APIUuid}
//}

func (r *repo) GetHallsMap() (model.HallResponse, error) {
	var halls []model.Hall
	if err := r.db.Preload("Tables").Find(&halls).Error; err != nil {
		return model.HallResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	hallsWithTables := make([]model.HallApi, 0, len(halls))
	for _, hall := range halls {
		tables := make([]model.TableApi, 0, len(hall.Tables))
		for _, t := range hall.Tables {
			table := model.TableApi{
				ID:      t.ID,
				Name:    t.Name,
				Persons: t.Persons,
				Place: model.Place{
					X: t.PlaceX,
					Y: t.PlaceY,
				},
				Size: model.Size{
					Width:  t.Width,
					Height: t.Height,
				},
				Type: t.Type,
			}
			tables = append(tables, table)
		}
		hallWithTables := model.HallApi{
			ID:     hall.ID,
			Name:   hall.Name,
			Tables: tables,
		}

		hallsWithTables = append(hallsWithTables, hallWithTables)
	}

	return model.HallResponse{
		Success: true,
		Halls:   hallsWithTables,
	}, nil

}

// вот эту логику ты выносишь в отдельный пакет и передаёшь сюда уже готовые структуры
//func (s *repo) ImportHallsFromGoulash() model.HallResponse {
//	req, err := http.NewRequest("GET", s.apiURL, nil)
//	if err != nil {
//		log.Fatal("NewRequest: ", err)
//	}
//	req.Header.Set("UUID", s.apiUUID)
//	req.Header.Set("AUTHORIZATION", s.apiKEY)
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Fatal("Do: ", err)
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//
//		}
//	}(resp.Body)
//
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal("ReadAll: ", err)
//	}
//
//	var halls model.HallResponse
//	if err := json.Unmarshal(body, &halls); err != nil {
//		fmt.Println(string(body))
//		log.Fatal("Unmarshal: ", err)
//	}
//
//	return s.UpdateDatabase(halls)
//}

func (r *repo) UpdateDatabase(response model.HallResponse) model.HallResponse {
	if !response.Success {
		return model.HallResponse{
			Success:      false,
			ErrorMessage: response.ErrorMessage,
		}
	}

	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return model.HallResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}
	}

	for _, hallGoulash := range response.Halls {
		var hall model.Hall
		if err := tx.Where("id = ?", hallGoulash.ID).FirstOrCreate(&hall).Error; err != nil {
			tx.Rollback()
			return model.HallResponse{
				Success:      false,
				ErrorMessage: err.Error(),
			}
		}

		hall.Name = hallGoulash.Name
		if err := tx.Save(&hall).Error; err != nil {
			tx.Rollback()
			return model.HallResponse{
				Success:      false,
				ErrorMessage: err.Error(),
			}
		}

		for _, tableGoulash := range hallGoulash.Tables {
			var table model.Table
			table.HallID = hall.ID

			if err := tx.Where("id = ?", tableGoulash.ID).Assign(model.Table{
				Name: tableGoulash.Name,
			}).FirstOrCreate(&table).Error; err != nil {
				tx.Rollback()
				return model.HallResponse{
					Success:      false,
					ErrorMessage: err.Error(),
				}
			}

			if err := tx.Save(&table).Error; err != nil {
				tx.Rollback()
				return model.HallResponse{
					Success:      false,
					ErrorMessage: err.Error(),
				}
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return model.HallResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}
	}

	return model.HallResponse{
		Success:      true,
		ErrorMessage: "",
	}
}
