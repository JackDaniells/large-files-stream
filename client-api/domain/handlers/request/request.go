package request

import proto "github.com/JackDaniells/port-service/proto"

type CreatePortRequest struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type UploadPortByFileRequest map[string]CreatePortRequest

func ParseUploadPortRequestToProtoPortArray(portsJson UploadPortByFileRequest) (ports []*proto.Port) {
	for key, portJson := range portsJson {
		port := &proto.Port{
			Id:          key,
			Name:        portJson.Name,
			City:        portJson.City,
			Province:    portJson.Province,
			Country:     portJson.Country,
			Alias:       portJson.Alias,
			Regions:     portJson.Regions,
			Coordinates: portJson.Coordinates,
			Timezone:    portJson.Timezone,
			Unlocs:      portJson.Unlocs,
			Code:        portJson.Code,
		}
		ports = append(ports, port)
	}
	return
}
