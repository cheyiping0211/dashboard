package service

import (
	"ping/global"
	"ping/model"
	"fmt"
)

func GetModules(module string) (err error, modules model.Modules) {
	err = global.GVA_DB.Where("module =  ? ", module).Preload("RelationChart").Find(&modules).Error
	return
}

func UpdateChart(chart model.Chart) (err error) {
	upDateChart := make(map[string]interface{})
	upDateChart["Id"] = chart.Id
	upDateChart["Title"] = chart.Title
	upDateChart["LastX"] = chart.LastX
	upDateChart["LastY"] = chart.LastY

	err = global.GVA_DB.Model(&chart).Where("id =  ? ", chart.Id).Omit("DomId","Type","ModuleType").Updates(upDateChart).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
