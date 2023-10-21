package storage

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/hypersdk/crypto/ed25519"
	"github.com/ava-labs/hypersdk/state"
)

// [kycPrefix] + [address]
func AliasAccountKey(al []byte) (k []byte) {
	// return []byte("asad")
	k = make([]byte, 1+KYCAlianChunks+KYCMetadataChunks)
	k[0] = aliasPrefix
	copy(k[1:], al[:])
	binary.BigEndian.PutUint16(k[1+KYCAlianChunks:], KYCMetadataChunks)
	return
}

func SetAccountAlias(
	ctx context.Context,
	mu state.Mutable,
	pk ed25519.PublicKey,
	al []byte,
) error {
	k := AliasAccountKey(al)
	return setAccountAlias(ctx, mu, k, pk)
}

func setAccountAlias(
	ctx context.Context,
	mu state.Mutable,
	key []byte,
	pk ed25519.PublicKey,
) error {
	kd := make([]byte, ed25519.PublicKeyLen)
	copy(kd[:], pk[:])
	return mu.Insert(ctx, key, kd)
}

func GetAccountAlias(
	ctx context.Context,
	im state.Immutable,
	al []byte,
) (ed25519.PublicKey, error) {
	k, pk, _, err := getAccountAlias(ctx, im, al)
	if err != nil {
		err = fmt.Errorf("%s::%w", hex.EncodeToString(k), err)
	}
	return pk, err
}

func getAccountAlias(
	ctx context.Context,
	im state.Immutable,
	al []byte,
) ([]byte, ed25519.PublicKey, bool, error) {
	k := AliasAccountKey(al)
	pk, exists, err := innerGetAccountAlias(im.GetValue(ctx, k))
	return k, pk, exists, err
}

func innerGetAccountAlias(
	v []byte,
	err error,
) (ed25519.PublicKey, bool, error) {
	if errors.Is(err, database.ErrNotFound) {
		return ed25519.EmptyPublicKey, false, nil
	}
	if err != nil {
		return ed25519.EmptyPublicKey, false, err
	}
	return ed25519.PublicKey(v[:]), true, nil
}

// Used to serve RPC queries
func GetAliasFromState(
	ctx context.Context,
	f ReadState,
	al []byte,
) (ed25519.PublicKey, error) {
	values, errs := f(ctx, [][]byte{AliasAccountKey(al)})
	pk, _, err := innerGetAccountAlias(values[0], errs[0])
	return pk, err
}
