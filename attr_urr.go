package gtp5gnl

import (
	"github.com/khirono/go-nl"
)

const (
	URR_ID = iota + 3
	URR_MEASUREMENT_METHOD
	URR_REPORTING_TRIGGER
	URR_MEASUREMENT_PERIOD
	URR_MEASUREMENT_INFO
	URR_SEID
	URR_VOLUME_THRESHOLD
	URR_VOLUME_QUOTA
)

const (
	URR_VOLUME_QUOTA_FLAG = iota + 1
	URR_VOLUME_QUOTA_TOVOL
	URR_VOLUME_QUOTA_UVOL
	URR_VOLUME_QUOTA_DVOL
)

const (
	URR_VOLUME_THRESHOLD_FLAG = iota + 1
	URR_VOLUME_THRESHOLD_TOVOL
	URR_VOLUME_THRESHOLD_UVOL
	URR_VOLUME_THRESHOLD_DVOL
)

type VolumeThreshold struct {
	flag           uint8
	totalVolume    uint64
	uplinkVolume   uint64
	downlinkVolume uint64
}

type VolumeQuota struct {
	flag           uint8
	totalVolume    uint64
	uplinkVolume   uint64
	downlinkVolume uint64
}

type URR struct {
	ID           uint32
	Method       uint8
	Trigger      uint32
	Period       *uint32
	Info         *uint8
	SEID         *uint64
	VolThreshold *VolumeThreshold
	VolQuota     *VolumeQuota
}

func DecodeURR(b []byte) (*URR, error) {
	urr := new(URR)
	for len(b) > 0 {
		hdr, n, err := nl.DecodeAttrHdr(b)
		if err != nil {
			return nil, err
		}
		switch hdr.MaskedType() {
		case URR_ID:
			urr.ID = native.Uint32(b[n:])
		case URR_MEASUREMENT_METHOD:
			urr.Method = uint8(b[n])
		case URR_REPORTING_TRIGGER:
			urr.Trigger = native.Uint32(b[n:])
		case URR_MEASUREMENT_PERIOD:
			v := native.Uint32(b[n:])
			urr.Period = &v
		case URR_MEASUREMENT_INFO:
			v := uint8(b[n])
			urr.Info = &v
		case URR_SEID:
			v := native.Uint64(b[n:])
			urr.SEID = &v
		case URR_VOLUME_THRESHOLD:
			volthreshold, err := DecodeVolumeThreshold(b[n:])
			if err != nil {
				return nil, err
			}
			urr.VolThreshold = &volthreshold
		case URR_VOLUME_QUOTA:
			volumequota, err := DecodeVolumeQuota(b[n:])
			if err != nil {
				return nil, err
			}
			urr.VolQuota = &volumequota
		}

		b = b[hdr.Len.Align():]
	}
	return urr, nil
}

func DecodeVolumeThreshold(b []byte) (VolumeThreshold, error) {
	var volumethreshold VolumeThreshold

	for len(b) > 0 {
		hdr, n, err := nl.DecodeAttrHdr(b)
		if err != nil {
			return volumethreshold, err
		}
		switch hdr.MaskedType() {
		case URR_VOLUME_THRESHOLD_FLAG:
			v := uint8(b[n])
			volumethreshold.flag = v
		case URR_VOLUME_THRESHOLD_TOVOL:
			v := native.Uint64(b[n:])
			volumethreshold.totalVolume = v
		case URR_VOLUME_THRESHOLD_UVOL:
			v := native.Uint64(b[n:])
			volumethreshold.uplinkVolume = v
		case URR_VOLUME_THRESHOLD_DVOL:
			v := native.Uint64(b[n:])
			volumethreshold.downlinkVolume = v
		}
		b = b[hdr.Len.Align():]
	}
	return volumethreshold, nil
}
func DecodeVolumeQuota(b []byte) (VolumeQuota, error) {
	var volumequota VolumeQuota

	for len(b) > 0 {
		hdr, n, err := nl.DecodeAttrHdr(b)
		if err != nil {
			return volumequota, err
		}
		switch hdr.MaskedType() {
		case URR_VOLUME_QUOTA_FLAG:
			v := uint8(b[n])
			volumequota.flag = v
		case URR_VOLUME_QUOTA_TOVOL:
			v := native.Uint64(b[n:])
			volumequota.totalVolume = v
		case URR_VOLUME_QUOTA_UVOL:
			v := native.Uint64(b[n:])
			volumequota.uplinkVolume = v
		case URR_VOLUME_QUOTA_DVOL:
			v := native.Uint64(b[n:])
			volumequota.downlinkVolume = v
		}
		b = b[hdr.Len.Align():]
	}
	return volumequota, nil
}
