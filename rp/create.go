package rp

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"github.com/tendermint/tmlibs/common"
)

const (
	namespace  = "namespace"
	identifier = "identifier"
	privateKey = "RP_PrivateKey"
)

var (
	CreatePath = fmt.Sprintf("/requests/:%s/:%s", namespace, identifier)
	refs       = map[string]string{}
	reqs       = map[string]Request{}
)

type DataRequestList struct {
	ServiceID string            `json:"service_id"`
	AsID      []string          `json:"as_id_list"`
	Count     int               `json:"count"`
	Params    map[string]string `json:"request_params"`
}

type Request struct {
	ReferenceID     string            `json:"reference_id"`
	IDPList         []string          `json:"idp_list"`
	CallBackURL     string            `json:"callback_url"`
	DataRequestList []DataRequestList `json:"data_request_list"`
	RequestMessage  string            `json:"request_message"`
	MinIal          int               `json:"min_ial"`
	MinAal          int               `json:"min_aal"`
	MinIdp          int               `json:"min_idp"`
	RequestTimeout  int               `json:"request_timeout"`
	Timeout         int               `json:"timeout"`
}

type Response struct {
	RequestID uuid.UUID `json:"request_id"`
}

func Create(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := validateParams(c.Param(namespace), c.Param(identifier))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = validatePayload(req)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}

func validateParams(ns, id string) error {
	return nil
}

func validatePayload(req Request) error {
	return nil
}

func create(req Request) string {
	if v, ok := refs[req.ReferenceID]; ok {
		return v
	}

	nonce := common.RandStr(12)
	reqID, _ := newRequestID(privateKey, nonce, req)
	bcDataList := []blockchainDataRequest{}

	for _, v := range req.DataRequestList {
		bcDataList = append(bcDataList, newBlockchainDataRequest(v))
	}

	if len(req.DataRequestList) != 0 {
		reqs[reqID] = req
	}

	r := blockchainRequest{
		RequestID:       reqID,
		MinIdp:          req.MinIdp,
		MinAal:          req.MinAal,
		MinIal:          req.MinIal,
		Timeout:         req.Timeout,
		DataRequestList: bcDataList,
		MessageHash:     fmt.Sprintf("%s", sha256.Sum256([]byte(req.RequestMessage))),
	}
	path, _ := buildBroadcastPath(r)

	http.Get(path)

	// ...
	return ""
}

func hashJSON(i interface{}) string {
	b, err := json.Marshal(i)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s", sha256.Sum256(b))
}

func newRequestID(pv, nonce string, req Request) (string, error) {
	b, err := json.Marshal(&req)
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("Concat_with_nonce_%s(%s)", nonce, base64.StdEncoding.EncodeToString(b))
	return fmt.Sprintf("%s", sha256.Sum256([]byte(s))), nil
}

type blockchainDataRequest struct {
	ServiceID         string
	AsID              []string
	Count             int
	RequestParamsHash string
}

func newBlockchainDataRequest(req DataRequestList) blockchainDataRequest {
	return blockchainDataRequest{
		ServiceID:         req.ServiceID,
		AsID:              req.AsID,
		Count:             req.Count,
		RequestParamsHash: hashJSON(req.Params),
	}
}

type blockchainRequest struct {
	RequestID       string                  `json:"request_id"`
	MinIal          int                     `json:"min_ial"`
	MinAal          int                     `json:"min_aal"`
	MinIdp          int                     `json:"min_idp"`
	Timeout         int                     `json:"timeout"`
	DataRequestList []blockchainDataRequest `json:"data_request_list"`
	MessageHash     string                  `json:"message_hash"`
}

func buildBroadcastPath(req blockchainRequest) (path string, err error) {
	nonce := common.RandStr(12)
	fn := "CreateRequest"
	b, err := json.Marshal(&req)
	if err != nil {
		return
	}

	tx := fmt.Sprintf("%s|%s|%s", fn, b, nonce)

	path = "/broadcast_tx_commit?tx=" + base64.StdEncoding.EncodeToString([]byte(tx))
	return
}
