// Copyright © 2018 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

// Package toa provides methods for computing a LoRaWAN packet's time-on-air
package toa

import (
	"math"
	"time"

	"github.com/TheThingsNetwork/ttn/pkg/errors"
	"github.com/TheThingsNetwork/ttn/pkg/ttnpb"
)

var (
	ErrInvalidCodingRate = &errors.ErrDescriptor{
		MessageFormat: "Invalid coding rate: cannot be different from 4/{5,6,7,8}",
		Code:          1,
		Type:          errors.InvalidArgument,
	}
	ErrUnknownModulation = &errors.ErrDescriptor{
		MessageFormat: "Unknown modulation",
		Code:          2,
		Type:          errors.Unknown,
	}
	ErrInvalidBandwidth = &errors.ErrDescriptor{
		MessageFormat:  "Invalid coding rate: cannot be {bandwidth}",
		Code:           3,
		Type:           errors.InvalidArgument,
		SafeAttributes: []string{"bandwidth"},
	}
	ErrInvalidSpreadingFactor = &errors.ErrDescriptor{
		MessageFormat:  "Invalid spreading factor: cannot be {spreading_factor}",
		Code:           4,
		Type:           errors.InvalidArgument,
		SafeAttributes: []string{"spreading_factor"},
	}
)

func init() {
	ErrInvalidCodingRate.Register()
	ErrUnknownModulation.Register()
	ErrInvalidBandwidth.Register()
	ErrInvalidSpreadingFactor.Register()
}

// Compute the time-on-air from the payload and RF parameters. This function only takes into account the PHY payload.
//
// See http://www.semtech.com/images/datasheet/LoraDesignGuide_STD.pdf, page 7
func Compute(rawPayload []byte, settings ttnpb.TxSettings) (time.Duration, error) {
	switch settings.Modulation {
	case ttnpb.Modulation_LORA:
		d, err := computeLoRa(rawPayload, settings)
		return d, err
	case ttnpb.Modulation_FSK:
		return computeFSK(rawPayload, settings), nil
	default:
		return 0, ErrUnknownModulation.New(nil)
	}
}

func computeLoRa(rawPayload []byte, settings ttnpb.TxSettings) (time.Duration, error) {
	var cr float64
	switch settings.CodingRate {
	case "4/5":
		cr = 1
	case "4/6":
		cr = 2
	case "4/7":
		cr = 3
	case "4/8":
		cr = 4
	default:
		return 0, ErrInvalidCodingRate.New(nil)
	}

	bandwidth := settings.Bandwidth / 1000 // Bandwidth in KHz
	spreadingFactor := settings.SpreadingFactor

	var de float64
	if bandwidth == 0 {
		return 0, ErrInvalidBandwidth.New(errors.Attributes{"bandwidth": 0})
	}
	if spreadingFactor < 7 || spreadingFactor > 12 {
		return 0, ErrInvalidSpreadingFactor.New(errors.Attributes{"spreading_factor": spreadingFactor})
	}
	if bandwidth == 125 && (spreadingFactor == 11 || spreadingFactor == 12) {
		de = 1.0
	}

	pl := float64(len(rawPayload))
	floatBW := float64(bandwidth)
	floatSF := float64(spreadingFactor)
	h := 0.0 // 0 means header is enabled

	tSym := math.Pow(2, floatSF) / floatBW

	payloadNb := 8.0 + math.Max(0.0, math.Ceil((8.0*pl-4.0*floatSF+28.0+16.0-20.0*h)/(4.0*(floatSF-2.0*de)))*(cr+4.0))
	timeOnAir := (payloadNb + 12.25) * tSym * 1000000 // in nanoseconds

	return time.Duration(timeOnAir), nil
}

func computeFSK(rawPayload []byte, settings ttnpb.TxSettings) time.Duration {
	payloadSize := len(rawPayload)
	bitRate := settings.BitRate

	tPkt := int64(time.Second) * (int64(payloadSize) + 5 + 3 + 1 + 2) * 8 / int64(bitRate)

	return time.Duration(tPkt)
}
