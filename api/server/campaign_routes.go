package server

import (
	"encoding/json"
	"fmt"
	"github.com/420Nat20/Nat20/nat-20/common"
	"github.com/420Nat20/Nat20/nat-20/data/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// RegisterCampaignRoutes attaches routes to the given router.
func (s *Server) RegisterCampaignRoutes(baseUrl string) {
	s.Router.HandleFunc(baseUrl+"/", s.getAllCampaigns).Methods("GET")
	s.Router.HandleFunc(baseUrl+"/{id:[0-9]+}", s.getCampaign).Methods("GET")
	s.Router.HandleFunc(baseUrl+"/", s.postCampaign).Methods("POST")
	s.Router.HandleFunc(baseUrl+"/{id:[0-9]+}", s.updateCampaign).Methods("PUT")
	s.Router.HandleFunc(baseUrl+"/{id:[0-9]+}", s.deleteCampaign).Methods("DELETE")
}

func (s *Server) getAllCampaigns(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	campaigns, err := s.CampaignService.GetAllCampaigns()
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
	}

	err = json.NewEncoder(w).Encode(campaigns)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}
}

func (s *Server) getCampaign(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	campaignId, err := strconv.Atoi(id)
	if err != nil {
		httpError := common.GetHttpError(common.BadRequest.New("Invalid id"))
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	campaign, err := s.CampaignService.GetCampaign(campaignId)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	err = json.NewEncoder(w).Encode(campaign)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}
}

func (s *Server) postCampaign(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var campaign models.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	err = s.CampaignService.CreateCampaign(&campaign)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	err = json.NewEncoder(w).Encode(campaign)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}
}

func (s *Server) updateCampaign(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	campaignId, err := strconv.Atoi(id)
	if err != nil {
		httpError := common.GetHttpError(common.BadRequest.New("Invalid id"))
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	var body map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	campaign, err := s.UserService.UpdateUser(campaignId, body)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	err = json.NewEncoder(w).Encode(campaign)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}
}

func (s *Server) deleteCampaign(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	campaignId, err := strconv.Atoi(id)
	if err != nil {
		httpError := common.GetHttpError(common.BadRequest.New("Invalid id"))
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	err = s.UserService.DeleteUser(campaignId)
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}

	_, err = fmt.Fprint(w, "Delete Successful")
	if err != nil {
		httpError := common.GetHttpError(err)
		http.Error(w, httpError.Message, httpError.StatusCode)
		return
	}
}
