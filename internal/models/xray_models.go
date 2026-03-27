package models

import (
	"fmt"

	"github.com/google/uuid"
)

type XrayLog struct {
	LogLevel string `json:"logLevel"`
}

type XrayClient struct {
	ID   uuid.UUID `json:"id"`
	Flow string    `json:"flow" default:"xtls-rprx-vision"`
}

func (xc *XrayClient) Str(clientInbound XrayInbound, host string, publicKey string, serverName string) string {
	return fmt.Sprintf(
		"%s://%s@%s:%d?encryption=%s&security=%s&sni=%s&fp=chrome&pbk=%s&sid=%s&type=%s&flow=%s#%s",
		clientInbound.Protocol,
		xc.ID.String(),
		host,
		clientInbound.Port,
		clientInbound.Settings.Decryption,
		clientInbound.StreamSettings.Security,
		clientInbound.StreamSettings.RealitySettings.ServerNames[0],
		publicKey,
		clientInbound.StreamSettings.RealitySettings.ShortIDs[0],
		clientInbound.StreamSettings.Network,
		xc.Flow,
		serverName,
	)
}

type XraySettings struct {
	Clients    []XrayClient `json:"clients"`
	Decryption string       `json:"decryption" default:"none"`
	Auth       string       `json:"auth omiempty"`
	Udp        bool         `json:"udp omitempty"`
}

type XrayRealitySettings struct {
	PrivateKey  string   `json:"privateKey" required:"true"`
	ShortIDs    []string `json:"shortIds" required:"true"`
	ServerNames []string `json:"serverNames" required:"true"`
	Dest        string   `json:"dest" default:"www.cloudflare.com:443"`
}

type XrayStreamSettings struct {
	Network         string              `json:"network" default:"tcp"`
	Security        string              `json:"security" default:"reality"`
	RealitySettings XrayRealitySettings `json:"realitySettings"`
}

type XrayInbound struct {
	Port           int                `json:"port"`
	Protocol       string             `json:"protocol"`
	Settings       XraySettings       `json:"settings"`
	StreamSettings XrayStreamSettings `json:"streamSettings"`
}

type XRayConfig struct {
	Log       XrayLog       `json:"log"`
	Inbounds  []XrayInbound `json:"inbounds"`
	Outbounds []any         `json:"outbounds"`
}
