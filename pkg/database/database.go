package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Driver string

const (
	MySQLDriver  Driver = "mysql"
	SQLiteDriver Driver = "sqlite"
)

func resolveDialector(driver Driver, conn string) gorm.Dialector {
	switch driver {
	case MySQLDriver:
		return mysql.Open(conn)
	case SQLiteDriver:
		return sqlite.Open(conn)
	}
	return nil
}

type Manager struct {
	*gorm.DB
	models []interface{}
}

func NewManager(driver Driver, conn string) (m *Manager, err error) {
	m = &Manager{}

	m.DB, err = gorm.Open(resolveDialector(driver, conn))
	return
}

func (m *Manager) RegisterModels(models ...interface{}) {
	m.models = append(m.models, models...)
}

func (m *Manager) Migrate() error {
	return m.AutoMigrate(m.models...)
}
