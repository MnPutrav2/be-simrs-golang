package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/MnPutrav2/be-simrs-golang/internal/models"
	"github.com/joho/godotenv"
)

type ObservationSatuSehatService interface {
	CreateObservationHeartRate(patient models.ObservatioClientRequest, token string) ([]byte, error)
}

type observationHeartRate struct {
	sql *sql.DB
}

func NewSatuSehatObservation(sql *sql.DB) ObservationSatuSehatService {
	return &observationHeartRate{sql}
}

type resObservation struct {
	Name     string `json:"name"`
	Resource string `json:"resource"`
}

func (q *observationHeartRate) CreateObservationHeartRate(patient models.ObservatioClientRequest, token string) ([]byte, error) {
	var data []resObservation

	heartBeat, _ := observationCommon(patient, patient.Pulse, "heart-beat", token)
	respiratori, _ := observationCommon(patient, patient.Respiratory, "respiratory", token)
	temp, _ := observationtemperature(patient, token)
	sistol, _ := observationTension(patient, token, patient.Sistol, "sistol")
	diastol, _ := observationTension(patient, token, patient.Diastol, "diastol")

	data = append(data, resObservation{Name: "heart beat", Resource: heartBeat})
	data = append(data, resObservation{Name: "respiratory", Resource: respiratori})
	data = append(data, resObservation{Name: "temperature", Resource: temp})
	data = append(data, resObservation{Name: "sistolik", Resource: sistol})
	data = append(data, resObservation{Name: "diastolik", Resource: diastol})

	res, _ := json.Marshal(data)

	return res, nil
}

func observationCommon(patient models.ObservatioClientRequest, d int, ty string, token string) (string, error) {

	_ = godotenv.Load()

	url := os.Getenv("SATU_SEHAT_END_POINT") + "/Observation"

	var lonic []string

	if ty == "heart-beat" {
		lonic = append(lonic, "Nadi ", "8867-4", "Heart rate", "beats/minute")
	} else if ty == "respiratory" {
		lonic = append(lonic, "Pernafasan ", "9279-1", "Respiratory rate", "breaths/minute")
	}

	c, _ := json.Marshal(models.ObservationCommon{
		ResourceType: "Observation",
		Status:       "final",
		Category: []models.Code{
			{
				Coding: []models.Coding{
					{
						System:  "http://terminology.hl7.org/CodeSystem/observation-category",
						Code:    "vital-signs",
						Display: "Vital Signs",
					},
				},
			},
		},
		Code: models.Code{
			Coding: []models.Coding{
				{
					System:  "http://loinc.org",
					Code:    lonic[1],
					Display: lonic[2],
				},
			},
		},
		Subject: models.ServiceProvider{
			Reference: "Patient/" + patient.PatientID,
		},
		Performer: []models.ServiceProvider{
			{
				Reference: "Practitioner/" + patient.PractitionerID,
			},
		},
		Encounter: models.Encounter2{
			Reference: "Encounter/" + patient.Encounter,
			Display:   "Pemeriksaan Fisik " + lonic[0] + patient.PatientName + " di tanggal " + patient.DateDisplay,
		},
		EffectiveDateTime: patient.Date,
		Issued:            patient.Date,
		ValueQuantity: models.ValueQuantity{
			Value:  int64(d),
			Unit:   lonic[3],
			System: "http://unitsofmeasure.org",
			Code:   "/min",
		},
	})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(c))
	if err != nil {
		return "", err
	}

	// Tambahkan header Content-Type dan Authorization Bearer
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Kirim request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == 201 {
		var data models.EncounterBodyResponse
		_ = json.Unmarshal(body, &data)

		return data.ID, nil
	}

	return "", nil
}

func observationtemperature(patient models.ObservatioClientRequest, token string) (string, error) {

	_ = godotenv.Load()

	url := os.Getenv("SATU_SEHAT_END_POINT") + "/Observation"

	var exam []string
	if patient.Temperature < 36 {
		exam = append(exam, "L", "Low", "dibawah nilai referensi")
	} else if patient.Temperature >= 36 && patient.Temperature < 37 {
		exam = append(exam, "N", "Normal", "normal")
	} else {
		exam = append(exam, "H", "High", "diatas nilai referensi")
	}

	c, _ := json.Marshal(models.ObservationTemperature{
		ResourceType: "Observation",
		Status:       "final",
		Category: []models.Code{
			{
				Coding: []models.Coding{
					{
						System:  "http://terminology.hl7.org/CodeSystem/observation-category",
						Code:    "vital-signs",
						Display: "Vital Signs",
					},
				},
			},
		},
		Code: models.Code{
			Coding: []models.Coding{
				{
					System:  "http://loinc.org",
					Code:    "8310-5",
					Display: "Body temperature",
				},
			},
		},
		Subject: models.ServiceProvider{
			Reference: "Patient/" + patient.PatientID,
		},
		Performer: []models.ServiceProvider{
			{
				Reference: "Practitioner/" + patient.PractitionerID,
			},
		},
		Encounter: models.Encounter2{
			Reference: "Encounter/" + patient.Encounter,
			Display:   "Pemeriksaan Fisik Suhu " + patient.PatientName + " di tanggal " + patient.DateDisplay,
		},
		EffectiveDateTime: patient.Date,
		Issued:            patient.Date,
		ValueQuantity: models.ValueQuantityTemp{
			Value:  patient.Temperature,
			Unit:   "C",
			System: "http://unitsofmeasure.org",
			Code:   "Cel",
		},
		Interpretation: []models.Interpretation{
			{
				Coding: []models.Coding{
					{
						System:  "http://terminology.hl7.org/CodeSystem/v3-ObservationInterpretation",
						Code:    exam[0],
						Display: exam[1],
					},
				},
				Text: exam[2],
			},
		},
	})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(c))
	if err != nil {
		return "", err
	}

	// Tambahkan header Content-Type dan Authorization Bearer
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Kirim request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == 201 {
		var data models.EncounterBodyResponse
		_ = json.Unmarshal(body, &data)

		return data.ID, nil
	}

	return "", nil
}

func observationTension(patient models.ObservatioClientRequest, token string, q int, ty string) (string, error) {

	_ = godotenv.Load()

	url := os.Getenv("SATU_SEHAT_END_POINT") + "/Observation"

	checkType := checkTension(q, ty)

	c, _ := json.Marshal(models.ObservationTension{
		ResourceType: "Observation",
		Status:       "final",
		Category: []models.Code{
			{
				Coding: []models.Coding{
					{
						System:  "http://terminology.hl7.org/CodeSystem/observation-category",
						Code:    "vital-signs",
						Display: "Vital Signs",
					},
				},
			},
		},
		Code: models.Code{
			Coding: []models.Coding{
				{
					System:  "http://loinc.org",
					Code:    checkType[3],
					Display: checkType[4],
				},
			},
		},
		Subject: models.Subject{
			Reference: "Patient/" + patient.PatientID,
			Display:   patient.PatientName,
		},
		Performer: []models.ServiceProvider{
			{
				Reference: "Practitioner/" + patient.PractitionerID,
			},
		},
		Encounter: models.Encounter2{
			Reference: "Encounter/" + patient.Encounter,
			Display:   "Pemeriksaan Fisik " + checkType[5] + patient.PatientName + " di tanggal " + patient.DateDisplay,
		},
		EffectiveDateTime: patient.Date,
		Issued:            patient.Date,
		BodySite: models.BodySite{
			Coding: []models.Coding{
				{
					System:  "http://snomed.info/sct",
					Code:    "368209003",
					Display: "Right arm",
				},
			},
		},
		ValueQuantity: models.ValueQuantity{
			Value:  int64(q),
			Unit:   "mm[Hg]",
			System: "http://unitsofmeasure.org",
			Code:   "mm[Hg]",
		},
		Interpretation: []models.Interpretation{
			{
				Coding: []models.Coding{
					{
						System:  "http://terminology.hl7.org/CodeSystem/v3-ObservationInterpretation",
						Code:    checkType[0],
						Display: checkType[1],
					},
				},
				Text: checkType[2],
			},
		},
	})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(c))
	if err != nil {
		return "", err
	}

	// Tambahkan header Content-Type dan Authorization Bearer
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Kirim request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == 201 {
		var data models.EncounterBodyResponse
		_ = json.Unmarshal(body, &data)

		return data.ID, nil
	}

	return "", nil
}

func checkTension(q int, ty string) []string {
	var exam []string

	if ty == "sistol" {
		if q < 90 {
			exam = append(exam, "L", "Low", "di bawah nilai referensi", "8480-6", "Systolic blood pressure", "Sistolik ")
		} else if q > 90 && q < 130 {
			exam = append(exam, "N", "Normal", "normal", "8480-6", "Systolic blood pressure", "Sistolik ")
		} else {
			exam = append(exam, "H", "High", "di atas nilai referensi", "8480-6", "Systolic blood pressure", "Sistolik ")
		}
	} else if ty == "diastol" {
		if q < 70 {
			exam = append(exam, "L", "Low", "di bawah nilai referensi", "8462-4", "Diastolic blood pressure", "Diastolik ")
		} else if q > 70 && q <= 90 {
			exam = append(exam, "N", "Normal", "normal", "8462-4", "Diastolic blood pressure", "Diastolik ")
		} else {
			exam = append(exam, "H", "High", "di atas nilai referensi", "8462-4", "Diastolic blood pressure", "Diastolik ")
		}
	}

	return exam
}
