package menu

import (
	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{
		db: db,
	}
}

func (m *menuRepo) GetMenu(menuType string) ([]model.MenuItem, error) {
	var menuData []model.MenuItem

	if err := m.db.Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return nil, err
	}

	return menuData, nil
}

/*
Dalam kasus ini, interface Repository memiliki metode GetMenu yang dideklarasikan dengan receiver pointer (*menuRepo).
Oleh karena itu, untuk mengembalikan sebuah nilai sebagai Repository, nilai tersebut harus memiliki tipe *menuRepo (pointer ke menuRepo)
untuk mematuhi kontrak interface tersebut.
*/
