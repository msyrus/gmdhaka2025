package model

type BaseModel struct {
	ID string `bson:"_id"`
}

func (m *BaseModel) SetID(id string) {
	m.ID = id
}

type User struct {
	BaseModel `bson:",inline"`
	Email     string `bson:"email"`
}

type Product struct {
	BaseModel `bson:",inline"`
	Name      string  `bson:"name"`
	Price     float64 `bson:"price"`
}
