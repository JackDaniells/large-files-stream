package handlers

import (
	"encoding/json"
	"github.com/JackDaniells/port-service/client-api/domain/contracts"
	"github.com/JackDaniells/port-service/client-api/domain/handlers/request"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type portHandler struct {
	portService contracts.PortService
}

func NewPortHandler(portService contracts.PortService) contracts.PortHandler {
	return &portHandler{
		portService: portService,
	}
}

func (s *portHandler) GetFileByID(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	portID := mux.Vars(request)["id"]

	port, err := s.portService.FindByID(ctx, portID)
	if err != nil {
		log.Println("error when get port by id: ", err)
		http.Error(response, "error when get port by id", http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(port)
	if err != nil {
		log.Println("error when marshal response: ", err)
		http.Error(response, "error when marshal response", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	_, err = response.Write(bytes)
	if err != nil {
		log.Println("error when write response: ", err)
		http.Error(response, "error when write response", http.StatusInternalServerError)
		return
	}
}

func (s *portHandler) UploadPortFileHandler(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	ports, err := s.decodeFile(request.Body)
	if err != nil {
		log.Println("error to decode file: ", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	createResponse, err := s.portService.StreamCreate(ctx, ports)
	if err != nil {
		log.Println("error when streaming ports: ", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(createResponse)
	if err != nil {
		log.Println("error when marshal response: ", err)
		http.Error(response, "error when marshal response", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	_, err = response.Write(bytes)
	if err != nil {
		log.Println("error when write response: ", err)
		http.Error(response, "error when write response", http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusCreated)
}

func (s *portHandler) decodeFile(fileStream io.ReadCloser) (request.UploadPortByFileRequest, error) {
	dec := json.NewDecoder(fileStream)
	var ports request.UploadPortByFileRequest
	for {
		if err := dec.Decode(&ports); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	return ports, nil
}
