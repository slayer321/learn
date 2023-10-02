package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	api "github.com/slayer321/learn/projects/activitytracker1/pkg/server"
)

type Actvities struct {
	URL       string
	Actvities []api.Activity
}

func (a *Actvities) Insert(act api.Activity) error {
	logrus.Infof("Inserting activity: %v", act)

	actReq := api.ActivityDocument{Activity: act}

	jsonBytes, err := json.Marshal(actReq)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, a.URL, bytes.NewReader(jsonBytes))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating request: %v\n", err)
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending request: %v\n", err)
		return err
	}

	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		return err
	}
	var reqID api.Activity
	json.Unmarshal(resByte, &reqID)
	reqID.ID = uint64(len(a.Actvities) + 1)
	a.Actvities = append(a.Actvities, reqID)
	fmt.Printf("ID: %d\n", reqID.ID)
	return nil
}

func (a *Actvities) Retrieve(id uint64) (api.Activity, error) {
	reqID := api.IDDocument{ID: id}
	jsonBytes, err := json.Marshal(reqID)
	if err != nil {
		logrus.Errorf("Error marshalling request: %v", err)
		return api.Activity{}, err
	}

	req, err := http.NewRequest(http.MethodGet, a.URL, bytes.NewReader(jsonBytes))
	if err != nil {
		logrus.Errorf("Error creating request: %v", err)
		return api.Activity{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Errorf("Error sending request: %v", err)
		return api.Activity{}, err
	}

	readBody, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Errorf("Error reading response: %v", err)
		return api.Activity{}, err
	}

	var act api.ActivityDocument

	json.Unmarshal(readBody, &act)
	return act.Activity, nil
}
