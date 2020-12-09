package model


type Modules struct {
    ID              uint       `gorm:"primary_key" json:"id"` 
    Module          string     `gorm:"not null" json:"module"` 
    ModuleName      string     `gorm:"not null" json:"moduleName"` 
    RelationChart   []Chart    `gorm:"ForeignKey:ModuleType;AssociationForeignKey:Module" json:"chartList" binding:"required"`
}

type Chart struct{
    Id         int       `gorm:"primary_key" json:"id"`   
    DomId      string    `json:"domId"` 
	Type       string    `json:"type"`
    Title      string    `json:"title"`
    LastX      int       `json:"lastX"`
    LastY      int       `json:"lastY"`
    ModuleType string    `json:"moduleType"`
}