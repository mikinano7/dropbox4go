package dropbox4go

import (
	"time"
	"net/http"
	"fmt"
	"encoding/json"
	"io"
	"io/ioutil"
)

const api = "https://content.dropboxapi.com/2/files/upload"

type Request struct {
	File       io.Reader
	Parameters Parameters
}

type Parameters struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	AutoRename bool `json:"autorename"`
	ClientModified string `json:"client_modified,omitempty"`
	Mute bool `json:"mute"`
}

type Response struct {
	Name string `json:"name"`
	PathLower string `json:"path_lower"`
	ClientModified time.Time `json:"client_modified"`
	ServerModified time.Time `json:"server_modified"`
	Rev string `json:"rev"`
	Size int `json:"size"`
	Id string `json:"id,omitempty"`
	MediaInfo string `json:"media_info,omitempty"`
	SharingInfo string `json:"sharing_info,omitempty"`
}

func (s *Service) Upload(request Request) (*Response, error) {
	parameters, err := json.Marshal(request.Parameters)
	req, err := http.NewRequest("POST", api, request.File)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))
	req.Header.Add("Dropbox-API-Arg", string(parameters))
	req.Header.Add("Content-Type", "application/octet-stream")

	apiResponse, err := s.c.Do(req)
	defer apiResponse.Body.Close()

	body, err := ioutil.ReadAll(apiResponse.Body)
	response := new(Response)
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	} else {
		return response, nil
	}
}
