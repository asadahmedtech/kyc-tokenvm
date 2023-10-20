package storage

type KYCData struct {
	Key          string `json:"key"`
	Exists       bool   `json:"exists"`
	KYCCountry   uint8  `json:"kyccountry"`
	KYCAuthority uint8  `json:"keyauthority"`
	KYCMetadata  []byte `json:"kycmetadata"`
}
