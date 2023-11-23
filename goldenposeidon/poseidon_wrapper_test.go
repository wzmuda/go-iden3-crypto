package poseidon

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type testVector struct {
	bytes        string
	expectedHash string
}

// Test cases below are taken from goldenposeidon/poseidon_test.go.
// The same set of input parameters was used:
//
//	const prime uint64 = 18446744069414584321
//	b0 := uint64(0)
//	b1 := uint64(1)
//	bm1 := prime - 1
//	bM := prime
//
// The following arrays were generated and stringified:
//   - {b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0}
//   - {b1, b1, b1, b1, b1, b1, b1, b1, b1, b1, b1, b1}
//   - {bm1, bm1, bm1, bm1, bm1, bm1, bm1, bm1, bm1, bm1, bm1, bm1}
//   - {bM, bM, bM, bM, bM, bM, bM, bM, b0, b0, b0, b0}
//   - {uint64(923978),
//     uint64(235763497586),
//     uint64(9827635653498),
//     uint64(112870),
//     uint64(289273673480943876),
//     uint64(230295874986745876),
//     uint64(6254867324987),
//     uint64(2087),
//     b0, b0, b0, b0}
//
// Expected hashes were also taken from goldenposeidon/poseidon_test.go and strngified.
var testVectors = []testVector{
	{
		bytes:        "00",
		expectedHash: "3c18a9786cb0b359c4055e3364a246c37953db0ab48808f4c71603f33a1144ca",
	},
	{
		bytes:        "000000000000000100000000000000010000000000000001000000000000000100000000000000010000000000000001000000000000000100000000000000010000000000000001000000000000000100000000000000010000000000000001",
		expectedHash: "e3fd1ad5743c4d77b94b3adc599d563009783d643dd45102a89f8f921605bbc8",
	},
	{
		bytes:        "ffffffff00000000ffffffff00000000ffffffff00000000ffffffff00000000ffffffff00000000ffffffff00000000ffffffff00000000ffffffff00000000ffffffff00000000ffffffff00000000ffffffff00000000ffffffff00000000",
		expectedHash: "be0085cfc57a8357d95af71847d05c09cf55a13d33c1c95395803a74f4530e82",
	},
	{
		bytes:        "ffffffff00000001ffffffff00000001ffffffff00000001ffffffff00000001ffffffff00000001ffffffff00000001ffffffff00000001ffffffff000000010000000000000000000000000000000000000000000000000000000000000000",
		expectedHash: "3c18a9786cb0b359c4055e3364a246c37953db0ab48808f4c71603f33a1144ca",
	},
	{
		bytes:        "00000000000e194a00000036e4997a72000008f02cbb6b7a000000000001b8e60403b4df96b4390403322ce0cdd2a014000005b05325203b00000000000008270000000000000000000000000000000000000000000000000000000000000000",
		expectedHash: "1a425786411bff9f0daa79475c066f986d29cb4f444a1af271437c9714fa79ff",
	},
}

func TestPoseidonWrapperSum(t *testing.T) {
	for i, vector := range testVectors {
		t.Run(fmt.Sprintf("test vector %d", i), func(t *testing.T) {
			inputBytes, err := hex.DecodeString(vector.bytes)
			require.NoError(t, err)

			hasher, err := New()
			require.NoError(t, err)
			hasher.Write(inputBytes)
			res := hasher.Sum(nil)

			require.NotEmpty(t, res)
			require.Equal(t, vector.expectedHash, hex.EncodeToString(res))
		})
	}
}

func TestPoseidonNewPoseidon(t *testing.T) {
	for i, vector := range testVectors {
		t.Run(fmt.Sprintf("test vector %d", i), func(t *testing.T) {
			inputBytes, err := hex.DecodeString(vector.bytes)
			require.NoError(t, err)

			res := Sum(inputBytes)

			require.NotEmpty(t, res)
			require.Equal(t, vector.expectedHash, hex.EncodeToString(res))
		})
	}
}
