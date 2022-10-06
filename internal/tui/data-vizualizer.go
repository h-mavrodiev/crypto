package tui

import (
	"crypto/internal/gate"
	"crypto/internal/stex"
)

type VisualizeInfo struct {
	Gate gate.GateInfo `json:"GateInfo"`
	Stex stex.StexInfo `json:"StexInfo"`
}

func (v *VisualizeInfo) UpdateGateInfoFields(gateInfo *gate.GateInfo) {
	v.Gate = *gateInfo
}

func (v *VisualizeInfo) UpdateStexInfoFields(stexInfo *stex.StexInfo) {
	v.Stex = *stexInfo
}
