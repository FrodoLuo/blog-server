package services

import "blog-server/models"

func CreateConfig(configToSave *models.Config) *models.Config {
	db := GetDB()

	db.Save(&configToSave)
	return configToSave
}

func UpdateConfig(configToSave *models.Config) *models.Config {
	db := GetDB()

	db.Save(&configToSave)
	db.Commit()

	return configToSave
}

func GetConfigByTitle(title string) *models.Config {
	db := GetDB()
	config := &models.Config{}
	db.
		Where("title = ?", title).
		First(&config)
	return config
}

func GetAllConfig() *[]models.Config {
	db := GetDB()
	configs := make([]models.Config, 0)
	db.Find(&configs)
	return &configs
}
