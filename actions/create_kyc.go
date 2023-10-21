// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package actions

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/consts"
	"github.com/ava-labs/hypersdk/examples/tokenvm/auth"
	"github.com/ava-labs/hypersdk/examples/tokenvm/storage"
	"github.com/ava-labs/hypersdk/state"
	"github.com/ava-labs/hypersdk/utils"
)

var _ chain.Action = (*CreateKYC)(nil)

type CreateKYC struct {
	KYCCheck     bool   `json:"kyccheck"`
	KYCCountry   uint8  `json:"kyccountry"`
	KYCAuthority uint8  `json:"keyauthority"`
	KYCMetadata  []byte `json:"kycmetadata"`
	KYCAlias     []byte `json:"kycalias"`
}

func (*CreateKYC) GetTypeID() uint8 {
	return createKYCID
}

func (k *CreateKYC) StateKeys(rauth chain.Auth, txID ids.ID) []string {
	return []string{
		string(storage.KYCAccountKey(auth.GetActor(rauth))),
		string(storage.AliasAccountKey(k.KYCAlias)),
	}
}

func (*CreateKYC) StateKeysMaxChunks() []uint16 {
	return []uint16{storage.KYCChucks, storage.KYCAlianChunks}
}

func (*CreateKYC) OutputsWarpMessage() bool {
	return false
}

func (c *CreateKYC) Execute(
	ctx context.Context,
	_ chain.Rules,
	mu state.Mutable,
	_ int64,
	rauth chain.Auth,
	txID ids.ID,
	_ bool,
) (bool, uint64, []byte, *warp.UnsignedMessage, error) {
	actor := auth.GetActor(rauth)
	if len(c.KYCMetadata) == 0 {
		return false, CreateKYCComputeUnits, OutputSymbolEmpty, nil, nil
	}

	if c.KYCAuthority != KYCAuthorityCompany && c.KYCAuthority != KYCAuthorityGov && c.KYCAuthority != KYCAuthorityIndivual {
		return false, CreateKYCComputeUnits, OutputSymbolEmpty, nil, nil
	}

	// It should only be possible to overwrite an existing asset if there is
	// a hash collision.
	if err := storage.SetAccountKYC(ctx, mu, actor, c.KYCCountry, c.KYCAuthority, c.KYCMetadata, c.KYCAlias); err != nil {
		return false, CreateKYCComputeUnits, utils.ErrBytes(err), nil, nil
	}

	// if err := storage.SetAccountAlias(ctx, mu, actor, c.KYCAlias); err != nil {
	// 	return false, CreateKYCComputeUnits, utils.ErrBytes(err), nil, nil
	// }

	// if kyc, err := storage.GetAccountKYC(ctx, mu, actor); err != nil {
	// 	err = fmt.Errorf("GetAccountKYC %w", err)
	// 	return false, CreateKYCComputeUnits, utils.ErrBytes(err), nil, nil
	// }

	return true, CreateKYCComputeUnits, OutputKYCCreated, nil, nil
}

func (*CreateKYC) MaxComputeUnits(chain.Rules) uint64 {
	return CreateKYCComputeUnits
}

func (c *CreateKYC) Size() int {
	// TODO: add small bytes (smaller int prefix)
	return consts.Uint16Len
}

func (c *CreateKYC) Marshal(p *codec.Packer) {
	// p.PackInt(int(c.KYCCountry))
	// p.PackInt(int(c.KYCCountry))
	p.PackByte(c.KYCCountry)
	p.PackByte(c.KYCAuthority)
	p.PackBytes([]byte(c.KYCAlias))
	p.PackBytes([]byte(c.KYCMetadata))
}

func UnmarshalCreateKYC(p *codec.Packer, _ *warp.Message) (chain.Action, error) {
	var create CreateKYC
	fmt.Println("TEst")
	fmt.Println(string(p.Bytes()))

	// create.KYCCountry = uint8(p.UnpackInt(true))
	// create.KYCAuthority = uint8(p.UnpackInt(true))

	create.KYCCountry = uint8(p.UnpackByte())
	create.KYCAuthority = uint8(p.UnpackByte())

	// var kc, ka []byte
	// p.UnpackBytes(1, true, &kc)
	// create.KYCCountry = kc[0]

	// p.UnpackBytes(1, true, &ka)
	// create.KYCAuthority = ka[0]

	// create.KYCAuthority = p.UnpackByte()
	p.UnpackBytes(CreateKYCAliasUnits, true, &create.KYCAlias)
	p.UnpackBytes(CreateKYCComputeUnits, true, &create.KYCMetadata)

	fmt.Println("TEst", create.KYCCountry, create.KYCAuthority, string(create.KYCAlias), string(create.KYCMetadata))
	return &create, p.Err()
}

func (*CreateKYC) ValidRange(chain.Rules) (int64, int64) {
	// Returning -1, -1 means that the action is always valid.
	return -1, -1
}
