package structLegajo

import (
	"github.com/jinzhu/gorm"
)

type Conyuge struct {
	gorm.Model
	Nombre       string     `json:"nombre"`
	Codigo       string     `json:"codigo"`
	Descripcion  string     `json:"descripcion"`
	Activo       int        `json:"activo"`
	Cuil         string     `json:"cuil"`
	Obrasocial   Obrasocial `json:"obrasocial" gorm:"ForeignKey:Obrasocialid;association_foreignkey:ID"`
	Obrasocialid uint       `json:"obrasocialid" sql:"type:int REFERENCES Obrasocial(ID)"`
	Legajoid     uint       `json:"legajoid"`
}
