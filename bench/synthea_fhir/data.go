// Code generated from github.com/go-json-experiment/jsonbench testdata; the
// synthea_fhir root type, with the root renamed to Benchmark for the lightning harness.
package bench

import (
	"encoding/json"
	"time"
)

type (
	Benchmark struct {
		Entry []struct {
			FullURL string `json:"fullUrl,nocopy"`
			Request *struct {
				Method string `json:"method,nocopy"`
				URL    string `json:"url,nocopy"`
			} `json:"request"`
			Resource *struct {
				AbatementDateTime time.Time   `json:"abatementDateTime"`
				AchievementStatus syntheaCode `json:"achievementStatus"`
				Active            bool        `json:"active"`
				Activity          []struct {
					Detail *struct {
						Code     syntheaCode      `json:"code"`
						Location syntheaReference `json:"location"`
						Status   string           `json:"status,nocopy"`
					} `json:"detail"`
				} `json:"activity"`
				Address        []syntheaAddress   `json:"address"`
				Addresses      []syntheaReference `json:"addresses"`
				AuthoredOn     time.Time          `json:"authoredOn"`
				BillablePeriod syntheaRange       `json:"billablePeriod"`
				BirthDate      string             `json:"birthDate,nocopy"`
				CareTeam       []struct {
					Provider  syntheaReference `json:"provider"`
					Reference string           `json:"reference,nocopy"`
					Role      syntheaCode      `json:"role"`
					Sequence  int64            `json:"sequence"`
				} `json:"careTeam"`
				Category       []syntheaCode    `json:"category"`
				Claim          syntheaReference `json:"claim"`
				Class          syntheaCoding    `json:"class"`
				ClinicalStatus syntheaCode      `json:"clinicalStatus"`
				Code           syntheaCode      `json:"code"`
				Communication  []struct {
					Language syntheaCode `json:"language"`
				} `json:"communication"`
				Component []struct {
					Code          syntheaCode   `json:"code"`
					ValueQuantity syntheaCoding `json:"valueQuantity"`
				} `json:"component"`
				Contained []struct {
					Beneficiary  syntheaReference   `json:"beneficiary"`
					ID           string             `json:"id,nocopy"`
					Intent       string             `json:"intent,nocopy"`
					Payor        []syntheaReference `json:"payor"`
					Performer    []syntheaReference `json:"performer"`
					Requester    syntheaReference   `json:"requester"`
					ResourceType string             `json:"resourceType,nocopy"`
					Status       string             `json:"status,nocopy"`
					Subject      syntheaReference   `json:"subject"`
					Type         syntheaCode        `json:"type"`
				} `json:"contained"`
				Created          time.Time   `json:"created"`
				DeceasedDateTime time.Time   `json:"deceasedDateTime"`
				Description      syntheaCode `json:"description"`
				Diagnosis        []struct {
					DiagnosisReference syntheaReference `json:"diagnosisReference"`
					Sequence           int64            `json:"sequence"`
					Type               []syntheaCode    `json:"type"`
				} `json:"diagnosis"`
				DosageInstruction []struct {
					AsNeededBoolean bool `json:"asNeededBoolean"`
					DoseAndRate     []struct {
						DoseQuantity *struct {
							Value float64 `json:"value"`
						} `json:"doseQuantity"`
						Type syntheaCode `json:"type"`
					} `json:"doseAndRate"`
					Sequence int64 `json:"sequence"`
					Timing   *struct {
						Repeat *struct {
							Frequency  int64   `json:"frequency"`
							Period     float64 `json:"period"`
							PeriodUnit string  `json:"periodUnit,nocopy"`
						} `json:"repeat"`
					} `json:"timing"`
				} `json:"dosageInstruction"`
				EffectiveDateTime time.Time          `json:"effectiveDateTime"`
				Encounter         syntheaReference   `json:"encounter"`
				Extension         []syntheaExtension `json:"extension"`
				Gender            string             `json:"gender,nocopy"`
				Goal              []syntheaReference `json:"goal"`
				ID                string             `json:"id,nocopy"`
				Identifier        []struct {
					System string      `json:"system,nocopy"`
					Type   syntheaCode `json:"type"`
					Use    string      `json:"use,nocopy"`
					Value  string      `json:"value,nocopy"`
				} `json:"identifier"`
				Insurance []struct {
					Coverage syntheaReference `json:"coverage"`
					Focal    bool             `json:"focal"`
					Sequence int64            `json:"sequence"`
				} `json:"insurance"`
				Insurer syntheaReference `json:"insurer"`
				Intent  string           `json:"intent,nocopy"`
				Issued  time.Time        `json:"issued"`
				Item    []struct {
					Adjudication []struct {
						Amount   syntheaCurrency `json:"amount"`
						Category syntheaCode     `json:"category"`
					} `json:"adjudication"`
					Category                syntheaCode        `json:"category"`
					DiagnosisSequence       []int64            `json:"diagnosisSequence"`
					Encounter               []syntheaReference `json:"encounter"`
					InformationSequence     []int64            `json:"informationSequence"`
					LocationCodeableConcept syntheaCode        `json:"locationCodeableConcept"`
					Net                     syntheaCurrency    `json:"net"`
					ProcedureSequence       []int64            `json:"procedureSequence"`
					ProductOrService        syntheaCode        `json:"productOrService"`
					Sequence                int64              `json:"sequence"`
					ServicedPeriod          syntheaRange       `json:"servicedPeriod"`
				} `json:"item"`
				LifecycleStatus           string             `json:"lifecycleStatus,nocopy"`
				ManagingOrganization      []syntheaReference `json:"managingOrganization"`
				MaritalStatus             syntheaCode        `json:"maritalStatus"`
				MedicationCodeableConcept syntheaCode        `json:"medicationCodeableConcept"`
				MultipleBirthBoolean      bool               `json:"multipleBirthBoolean"`
				Name                      json.RawMessage    `json:"name"`
				NumberOfInstances         int64              `json:"numberOfInstances"`
				NumberOfSeries            int64              `json:"numberOfSeries"`
				OccurrenceDateTime        time.Time          `json:"occurrenceDateTime"`
				OnsetDateTime             time.Time          `json:"onsetDateTime"`
				Outcome                   string             `json:"outcome,nocopy"`
				Participant               []struct {
					Individual syntheaReference `json:"individual"`
					Member     syntheaReference `json:"member"`
					Role       []syntheaCode    `json:"role"`
				} `json:"participant"`
				Patient syntheaReference `json:"patient"`
				Payment *struct {
					Amount syntheaCurrency `json:"amount"`
				} `json:"payment"`
				PerformedPeriod syntheaRange     `json:"performedPeriod"`
				Period          syntheaRange     `json:"period"`
				Prescription    syntheaReference `json:"prescription"`
				PrimarySource   bool             `json:"primarySource"`
				Priority        syntheaCode      `json:"priority"`
				Procedure       []struct {
					ProcedureReference syntheaReference `json:"procedureReference"`
					Sequence           int64            `json:"sequence"`
				} `json:"procedure"`
				Provider        syntheaReference   `json:"provider"`
				ReasonCode      []syntheaCode      `json:"reasonCode"`
				ReasonReference []syntheaReference `json:"reasonReference"`
				RecordedDate    time.Time          `json:"recordedDate"`
				Referral        syntheaReference   `json:"referral"`
				Requester       syntheaReference   `json:"requester"`
				ResourceType    string             `json:"resourceType,nocopy"`
				Result          []syntheaReference `json:"result"`
				Series          []struct {
					BodySite syntheaCoding `json:"bodySite"`
					Instance []struct {
						Number   int64         `json:"number"`
						SopClass syntheaCoding `json:"sopClass"`
						Title    string        `json:"title,nocopy"`
						UID      string        `json:"uid,nocopy"`
					} `json:"instance"`
					Modality          syntheaCoding `json:"modality"`
					Number            int64         `json:"number"`
					NumberOfInstances int64         `json:"numberOfInstances"`
					Started           string        `json:"started,nocopy"`
					UID               string        `json:"uid,nocopy"`
				} `json:"series"`
				ServiceProvider syntheaReference `json:"serviceProvider"`
				Started         time.Time        `json:"started"`
				Status          string           `json:"status,nocopy"`
				Subject         syntheaReference `json:"subject"`
				SupportingInfo  []struct {
					Category       syntheaCode      `json:"category"`
					Sequence       int64            `json:"sequence"`
					ValueReference syntheaReference `json:"valueReference"`
				} `json:"supportingInfo"`
				Telecom              []map[string]string `json:"telecom,nocopy"`
				Text                 map[string]string   `json:"text,nocopy"`
				Total                json.RawMessage     `json:"total"`
				Type                 json.RawMessage     `json:"type"`
				Use                  string              `json:"use,nocopy"`
				VaccineCode          syntheaCode         `json:"vaccineCode"`
				ValueCodeableConcept syntheaCode         `json:"valueCodeableConcept"`
				ValueQuantity        syntheaCoding       `json:"valueQuantity"`
				VerificationStatus   syntheaCode         `json:"verificationStatus"`
			} `json:"resource"`
		} `json:"entry"`
		ResourceType string `json:"resourceType,nocopy"`
		Type         string `json:"type,nocopy"`
	}
	syntheaCode struct {
		Coding []syntheaCoding `json:"coding"`
		Text   string          `json:"text,nocopy"`
	}
	syntheaCoding struct {
		Code    string  `json:"code,nocopy"`
		Display string  `json:"display,nocopy"`
		System  string  `json:"system,nocopy"`
		Unit    string  `json:"unit,nocopy"`
		Value   float64 `json:"value"`
	}
	syntheaReference struct {
		Display   string `json:"display,nocopy"`
		Reference string `json:"reference,nocopy"`
	}
	syntheaAddress struct {
		City       string             `json:"city,nocopy"`
		Country    string             `json:"country,nocopy"`
		Extension  []syntheaExtension `json:"extension"`
		Line       []string           `json:"line,nocopy"`
		PostalCode string             `json:"postalCode,nocopy"`
		State      string             `json:"state,nocopy"`
	}
	syntheaExtension struct {
		URL          string             `json:"url,nocopy"`
		ValueAddress syntheaAddress     `json:"valueAddress"`
		ValueCode    string             `json:"valueCode,nocopy"`
		ValueDecimal float64            `json:"valueDecimal"`
		ValueString  string             `json:"valueString,nocopy"`
		Extension    []syntheaExtension `json:"extension"`
	}
	syntheaRange struct {
		End   time.Time `json:"end"`
		Start time.Time `json:"start"`
	}
	syntheaCurrency struct {
		Currency string  `json:"currency,nocopy"`
		Value    float64 `json:"value"`
	}
)
