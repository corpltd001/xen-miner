package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

func GetLatestMemoryDifficulty() (uint32, error) {
	resp, err := http.Get("http://xenminer.mooo.com/difficulty")
	if err != nil {
		return 0, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			logrus.Errorf("failed to close resp body reader")
		}
	}()

	ans := &MemoryDifficultyResponse{}
	if err := json.Unmarshal(body, ans); err != nil {
		return 0, err
	}

	difficulty, err := strconv.ParseUint(ans.Difficulty, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(difficulty), nil
}

type Payload struct {
	HashToVerify string `json:"hash_to_verify"`
	Key          string `json:"key"`
	Account      string `json:"account"`
}

const verifyURL = `http://xenminer.mooo.com/verify`

func Verify(payload Payload) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	logrus.Infof("payload: %s", string(data))

	req, err := http.NewRequest(http.MethodPost, verifyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_ = resp.Body.Close()

	ans := make(map[string]interface{})
	if err = json.Unmarshal(body, &ans); err != nil {
		return err
	}
	logrus.Infof("HTTP Status Code: %d", resp.StatusCode)
	logrus.Infof("resp: %s", marshalToJson(ans))
	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}

	return nil
}
