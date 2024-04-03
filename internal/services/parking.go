package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"parking-app-web-api-gateway/internal/utils"
)

type ParkingService struct {
	URL string
}

func NewParkingService() *ParkingService {
	url, _ := utils.GetEnvVar("PARKING_SERVICE_URL")

	return &ParkingService{
		URL: url,
	}
}

type ParkingSpot struct {
	Name       string `json:"name"`
	Location   string `json:"location"`
	Occupied   bool   `json:"occupied"`
	Car        string `json:"car"`
	OccupiedAt string `json:"occupiedAt"`
}

type CreateParkingSpotRequest struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (p *ParkingService) GetParkingSpots() ([]string, error) {

	req, err := http.NewRequest("GET", p.URL+"/park", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var parkingSpots []string

	err = json.NewDecoder(resp.Body).Decode(&parkingSpots)
	if err != nil {
		return nil, err
	}

	return parkingSpots, nil
}

func (p *ParkingService) GetParkingSpot(name string) (*ParkingSpot, error) {

	req, err := http.NewRequest("GET", p.URL+"/park/"+name, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var parkingSpot ParkingSpot

	err = json.NewDecoder(resp.Body).Decode(&parkingSpot)
	if err != nil {
		return nil, err
	}

	return &parkingSpot, nil
}

func (p *ParkingService) AddParkingSpot(parkingSpot *CreateParkingSpotRequest) error {

	body, err := json.Marshal(parkingSpot)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", p.URL+"/park", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
