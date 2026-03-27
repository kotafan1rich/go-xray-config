package repository

import (
	"encoding/json"
	"fmt"
	"go-xray-config/internal/models"
	"os"

	"github.com/google/uuid"
)

type XrayRepository interface {
	Add() (string, error)
	Write(config models.XRayConfig) error
	Delete(id uuid.UUID) error
	GetAll() (models.XRayConfig, error)
}

type xrayRepository struct {
	xrayPath   string
	host       string
	publicKey  string
	serverName string
}

func NewXrayRepository(xrayPath string, host string, publicKey string, serverName string) XrayRepository {
	return &xrayRepository{
		xrayPath:   xrayPath,
		host:       host,
		publicKey:  publicKey,
		serverName: serverName,
	}
}

func (r *xrayRepository) GetAll() (models.XRayConfig, error) {
	file, err := os.ReadFile(r.xrayPath)
	var xrayConfig models.XRayConfig
	if err != nil || len(file) == 0 {
		return xrayConfig, fmt.Errorf("failed to read xray config: %v", err)
	} else {
		err = json.Unmarshal(file, &xrayConfig)
		if err != nil {
			return xrayConfig, err
		}
	}
	return xrayConfig, nil
}

func (r *xrayRepository) Write(config models.XRayConfig) error {
	file, err := os.Create(r.xrayPath)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, _ := json.MarshalIndent(config, "", "  ")
	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func (r *xrayRepository) Add() (string, error) {
	xrayConfig, err := r.GetAll()
	if err != nil {
		return "", err
	}

	for i := range xrayConfig.Inbounds {
		inbound := &xrayConfig.Inbounds[i]
		if inbound.Protocol == "vless" {
			newClient := models.XrayClient{
				ID:   uuid.New(),
				Flow: "xtls-rprx-vision",
			}
			inbound.Settings.Clients = append(inbound.Settings.Clients, newClient)
			err = r.Write(xrayConfig)
			if err != nil {
				return "", err
			}
			return newClient.Str(*inbound, r.host, r.publicKey, r.serverName), nil
		}
	}
	return "", nil
}

func (r *xrayRepository) Delete(id uuid.UUID) error {
	xrayConfig, err := r.GetAll()
	if err != nil {
		return err
	}

	for i := range xrayConfig.Inbounds {
		inbound := &xrayConfig.Inbounds[i]
		if inbound.Protocol == "vless" {
			for j := range inbound.Settings.Clients {
				client := &inbound.Settings.Clients[j]
				if client.ID == id {
					inbound.Settings.Clients = append(inbound.Settings.Clients[:j], inbound.Settings.Clients[j+1:]...)
					return r.Write(xrayConfig)
				}
			}
			return fmt.Errorf("client with id %s not found", id)
		}
	}
	return fmt.Errorf("client with id %s not found", id)
}
