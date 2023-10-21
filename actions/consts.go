// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package actions

// Note: Registry will error during initialization if a duplicate ID is assigned. We explicitly assign IDs to avoid accidental remapping.
const (
	burnAssetID   uint8 = 0
	closeOrderID  uint8 = 1
	createAssetID uint8 = 2
	exportAssetID uint8 = 3
	importAssetID uint8 = 4
	createOrderID uint8 = 5
	fillOrderID   uint8 = 6
	mintAssetID   uint8 = 7
	transferID    uint8 = 8
	createKYCID   uint8 = 9
)

const (
	// TODO: tune this
	BurnComputeUnits        = 2
	CloseOrderComputeUnits  = 5
	CreateAssetComputeUnits = 10
	ExportAssetComputeUnits = 10
	ImportAssetComputeUnits = 10
	CreateOrderComputeUnits = 5
	NoFillOrderComputeUnits = 5
	FillOrderComputeUnits   = 15
	MintAssetComputeUnits   = 2
	TransferComputeUnits    = 1

	MaxSymbolSize   = 8
	MaxMemoSize     = 256
	MaxMetadataSize = 256
	MaxDecimals     = 9

	MaxKYCSize            = 5
	CreateKYCComputeUnits = 20
	CreateKYCAliasUnits   = 32
)

const (
	KYCCountryInd uint8 = 1
	KYCCountrySG  uint8 = 2
	KYCCountryUS  uint8 = 3
	KYCCountryUK  uint8 = 4
	KYCCountryAus uint8 = 5
)

const (
	KYCAuthorityIndivual uint8 = 1
	KYCAuthorityGov      uint8 = 2
	KYCAuthorityCompany  uint8 = 3
)

type KYCDetails struct {
	Code uint8  `json:"code"`
	Name string `json:"name"`
}

var KYCCountryList = []KYCDetails{
	KYCDetails{KYCCountryInd, "India"},
	KYCDetails{KYCCountrySG, "Signapore"},
	KYCDetails{KYCCountryUS, "USA"},
	KYCDetails{KYCCountryUK, "UK"},
	KYCDetails{KYCCountryAus, "Australia"},
}

var KYCAuthorityList = []KYCDetails{
	KYCDetails{KYCAuthorityIndivual, "Indivual"},
	KYCDetails{KYCAuthorityGov, "Government"},
	KYCDetails{KYCAuthorityCompany, "Company"},
}
