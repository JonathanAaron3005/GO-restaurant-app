package menu

import (
	"context"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"github.com/JonathanAaron3005/go-restaurant-app/internal/tracing"
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

func (m *menuRepo) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()

	var menuData []model.MenuItem

	if err := m.db.WithContext(ctx).Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return nil, err
	}

	return menuData, nil
}

func (m *menuRepo) GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenu")
	defer span.End()

	var menuData model.MenuItem

	if err := m.db.WithContext(ctx).Where(model.MenuItem{OrderCode: orderCode}).First(&menuData).Error; err != nil {
		return menuData, err
	}

	return menuData, nil
}

func (m *menuRepo) CreateMenu(ctx context.Context, menu model.MenuItem) (model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateMenu")
	defer span.End()

	if err := m.db.WithContext(ctx).Create(&menu).Error; err != nil {
		return menu, err
	}

	return menu, nil
}

/*
Dalam kasus ini, interface Repository memiliki metode GetMenu yang dideklarasikan dengan receiver pointer (*menuRepo).
Oleh karena itu, untuk mengembalikan sebuah nilai sebagai Repository, nilai tersebut harus memiliki tipe *menuRepo (pointer ke menuRepo)
untuk mematuhi kontrak interface tersebut.
*/
