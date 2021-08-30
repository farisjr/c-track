package database

import (
	"app/config"
	"app/models"
)

func CreateTrack(createTrack models.Track) (models.Track, error) {
	if err := config.DB.Save(&createTrack).Error; err != nil {
		return createTrack, err
	}
	return createTrack, nil
}

func GetAllTracks() ([]models.Track, error) {
	var track []models.Track
	if err := config.DB.Find(&track).Error; err != nil {
		return nil, err
	}
	return track, nil
}

func GetOneTrackbyChecker(checkerId int) (models.Track, error) {
	var track models.Track
	if err := config.DB.Where("checker_id=?", checkerId).First(&track).Error; err != nil {
		return track, err
	}
	return track, nil
}

func GetOneTrackbyPatient(patientId int) (models.Track, error) {
	var track models.Track
	if err := config.DB.Where("patient_id=?", patientId).First(&track).Error; err != nil {
		return track, err
	}
	return track, nil
}

//update track info from database
func UpdateTrack(track models.Track) (interface{}, error) {
	if err := config.DB.Save(&track).Error; err != nil {
		return nil, err
	}
	return track, nil
}

func GetUpdateTrack(trackId int) models.Track {
	var track models.Track
	config.DB.Find(&track, "track_id=?", trackId)
	return track
}
