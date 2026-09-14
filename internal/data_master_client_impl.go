package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type DataMasterClientImpl struct {
	Host       string
	Sumber     string
	HttpClient *http.Client
}

func NewDataMasterClient(httpClient *http.Client) *DataMasterClientImpl {
	dataMasterHost := os.Getenv("MASTER_DATA_SERVICE_HOST")

	return &DataMasterClientImpl{
		Host:       dataMasterHost,
		Sumber:     "EKAK",
		HttpClient: httpClient,
	}
}

type DataMasterResponse[T any] struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type MappingOpd struct {
	KodeSumber string `json:"kode_sumber"`
	KodeMaster string `json:"kode_master"`
}

type Pegawai struct {
	ID                     uint64  `json:"id"`
	OpdKode                string  `json:"opd_kode"`
	PegawaiNIP             string  `json:"pegawai_nip"`
	PegawaiNama            string  `json:"pegawai_nama"`
	PegawaiJabatanTerakhir *string `json:"pegawai_jabatan_terakhir,omitempty"`
	OpdNama                *string `json:"opd_nama,omitempty"`
}

func (c *DataMasterClientImpl) FindMappingOpd(
	ctx context.Context,
	kodeOpd string,
) (*MappingOpd, error) {
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, fmt.Errorf("parse data master host: %w", err)
	}

	u.Path = "/mapping-opd/find-by-kode-sumber"

	query := u.Query()
	query.Set("sumber", c.Sumber)
	query.Set("kode_sumber", kodeOpd)
	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		u.String(),
		nil,
	)
	log.Printf("URL: %s", u.String())
	if err != nil {
		return nil, fmt.Errorf("create data master request: %w", err)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call data master: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"data master returned unexpected status: %d",
			resp.StatusCode,
		)
	}

	var wrapper DataMasterResponse[MappingOpd]

	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, fmt.Errorf(
			"decode data master response: %w",
			err,
		)
	}

	return &wrapper.Data, nil
}

func (c *DataMasterClientImpl) FindPegawaiByKodeOpd(
	ctx context.Context,
	kodeOpdMaster string,
) ([]Pegawai, error) {
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, fmt.Errorf(
			"parse data master host: %w",
			err,
		)
	}

	u.Path = "/pegawai/find"

	query := u.Query()
	query.Set("kodeOpd", kodeOpdMaster)
	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		u.String(),
		nil,
	)
	log.Printf("URL: %s", u.String())
	if err != nil {
		return nil, fmt.Errorf(
			"create data master request: %w",
			err,
		)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(
			"call data master: %w",
			err,
		)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"data master returned unexpected status: %d",
			resp.StatusCode,
		)
	}

	var wrapper DataMasterResponse[[]Pegawai]

	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, fmt.Errorf(
			"decode data master response: %w",
			err,
		)
	}

	return wrapper.Data, nil
}
