package wgconfig

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Write(t *testing.T) {
	cfg := Config{
		Interface: Interface{
			ListenPort: 5559,
			Address:    "10.77.24.1/24",
			DNS:        []string{"8.8.8.8", "8.8.4.4"},
			MTU:        800,
			PrivateKey: "4CwbPHW85Y/xdgB/zD/P0bZdM3XVNpi85H45FMscB1A=",
			PostUp:     "iptables -A FORWARD -i %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE",
			PostDown:   "iptables -D FORWARD -i %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE",
		},
		Peers: Peers{
			{
				Comment:    "Comment1",
				PublicKey:  "gIIbPCSRw7qQnW/aS3g1PjZTEXnTBqSjo8sS9MADows=",
				AllowedIPs: "10.77.24.22/32",
				Endpoint:   "endpoint",
			},
			{
				Comment:    "Comment",
				PublicKey:  "KKJfVUC8awDEa4H7Pa5lRCvnq3cdrLMHpZVNF7YkgVA=",
				AllowedIPs: "10.77.24.24/32",
				Endpoint:   "endpoint2",
			},
			{
				PublicKey:           "IJgEGy5QPRbwuf7yY1+bbirFeHoNwdYzIfrWMNFEG30=",
				AllowedIPs:          "10.77.24.26/32",
				PersistentKeepalive: 30,
			},
			{
				PublicKey:  "NafllWlCPqa4Jhv10Rjbk38pxyWiWcpkwRYwcd47qic=",
				AllowedIPs: "10.77.24.28/32",
			},
		},
	}

	expected := `[Interface]
PrivateKey = 4CwbPHW85Y/xdgB/zD/P0bZdM3XVNpi85H45FMscB1A=
Address    = 10.77.24.1/24
ListenPort = 5559
DNS        = 8.8.8.8,8.8.4.4
MTU        = 800
PostUp     = iptables -A FORWARD -i %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
PostDown   = iptables -D FORWARD -i %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE

; Comment1
[Peer]
PublicKey  = gIIbPCSRw7qQnW/aS3g1PjZTEXnTBqSjo8sS9MADows=
AllowedIPs = 10.77.24.22/32
Endpoint   = endpoint

; Comment
[Peer]
PublicKey  = KKJfVUC8awDEa4H7Pa5lRCvnq3cdrLMHpZVNF7YkgVA=
AllowedIPs = 10.77.24.24/32
Endpoint   = endpoint2

[Peer]
PublicKey           = IJgEGy5QPRbwuf7yY1+bbirFeHoNwdYzIfrWMNFEG30=
AllowedIPs          = 10.77.24.26/32
PersistentKeepalive = 30

[Peer]
PublicKey  = NafllWlCPqa4Jhv10Rjbk38pxyWiWcpkwRYwcd47qic=
AllowedIPs = 10.77.24.28/32
`

	data := new(bytes.Buffer)
	_, err := cfg.Write(data)

	assert.NoError(t, err)
	assert.Equal(t, expected, data.String())
}

func TestConfig_Read(t *testing.T) {
	input := []byte(`[Interface]
Address = 10.8.16.1/24
ListenPort = 51820
PrivateKey = 4CwbPHW85Y/xdgB/zD/P0bZdM3XVNpi85H45FMscB1A=
DNS = 1.1.1.1,1.1.0.0
MTU = 1500
PreUp = echo "UP"
PostUp = iptables -A FORWARD -i %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
PreDown = echo "DOWN"
PostDown = iptables -D FORWARD -i %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE

# Peer 1
[Peer]
PublicKey = HHQSHN5TG6d0f3Wo0zeJM74v6073rQhc1+Yc8cwQ32Q=
AllowedIPs = 10.8.16.2/32
Endpoint = https://example.com:9800

; Peer 2
[Peer]
PublicKey = ttHzRDWUmVHWn+CXBGj04fYwdeb51wIUt0iC8ejP2wo=
AllowedIPs = 10.8.16.3/32
PersistentKeepalive = 20

[Peer]
PublicKey = 064r3zzmeaCGCEwXlfj+2tNV6tTnxbFiZalk1XIY7wI=
AllowedIPs = 10.8.16.4/32`)

	expected := Config{
		Interface: Interface{
			ListenPort: 51820,
			Address:    "10.8.16.1/24",
			PrivateKey: "4CwbPHW85Y/xdgB/zD/P0bZdM3XVNpi85H45FMscB1A=",
			DNS:        []string{"1.1.1.1", "1.1.0.0"},
			MTU:        1500,
			PreUp:      "echo \"UP\"",
			PostUp:     "iptables -A FORWARD -i %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE",
			PreDown:    "echo \"DOWN\"",
			PostDown:   "iptables -D FORWARD -i %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE",
		},
		Peers: Peers{
			{
				Comment:    "Peer 1",
				PublicKey:  "HHQSHN5TG6d0f3Wo0zeJM74v6073rQhc1+Yc8cwQ32Q=",
				AllowedIPs: "10.8.16.2/32",
				Endpoint:   "https://example.com:9800",
			},
			{
				Comment:             "Peer 2",
				PublicKey:           "ttHzRDWUmVHWn+CXBGj04fYwdeb51wIUt0iC8ejP2wo=",
				AllowedIPs:          "10.8.16.3/32",
				PersistentKeepalive: 20,
			},
			{
				PublicKey:  "064r3zzmeaCGCEwXlfj+2tNV6tTnxbFiZalk1XIY7wI=",
				AllowedIPs: "10.8.16.4/32",
			},
		},
	}

	cfg := Config{}
	err := cfg.Read(bytes.NewBuffer(input))

	assert.NoError(t, err)
	assert.Equal(t, expected, cfg)
}

func TestConfig_ReadInlineComment(t *testing.T) {
	input := []byte(`[Interface]
PostUp = iptables -A FORWARD -i %i -j ACCEPT ; comment
`)

	expected := Config{
		Interface: Interface{
			PostUp: "iptables -A FORWARD -i %i -j ACCEPT",
		},
	}
	cfg := Config{}
	err := cfg.Read(bytes.NewBuffer(input))

	assert.NoError(t, err)
	assert.Equal(t, expected, cfg)
}
