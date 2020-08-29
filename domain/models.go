package domain

type Migrator interface {
	Migrate()
}

func GetModels() []interface{} {
	return []interface{}{
		&Member{},
		&Product{},
	}
}
