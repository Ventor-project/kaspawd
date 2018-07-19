// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"bytes"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// TestGenesisBlock tests the genesis block of the main network for validity by
// checking the encoded bytes and hashes.
func TestGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := MainNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), genesisBlockBytes) {
		t.Fatalf("TestGenesisBlock: Genesis block does not appear valid - "+
			"got %v, want %v", spew.Sdump(buf.Bytes()),
			spew.Sdump(genesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := MainNetParams.GenesisBlock.BlockHash()
	if !MainNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestGenesisBlock: Genesis block hash does not "+
			"appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(MainNetParams.GenesisHash))
	}
}

// TestRegTestGenesisBlock tests the genesis block of the regression test
// network for validity by checking the encoded bytes and hashes.
func TestRegTestGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := RegressionNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestRegTestGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), regTestGenesisBlockBytes) {
		t.Fatalf("TestRegTestGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(regTestGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := RegressionNetParams.GenesisBlock.BlockHash()
	if !RegressionNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestRegTestGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(RegressionNetParams.GenesisHash))
	}
}

// TestTestNet3GenesisBlock tests the genesis block of the test network (version
// 3) for validity by checking the encoded bytes and hashes.
func TestTestNet3GenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := TestNet3Params.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestTestNet3GenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), testNet3GenesisBlockBytes) {
		t.Fatalf("TestTestNet3GenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(testNet3GenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := TestNet3Params.GenesisBlock.BlockHash()
	if !TestNet3Params.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestTestNet3GenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(TestNet3Params.GenesisHash))
	}
}

// TestSimNetGenesisBlock tests the genesis block of the simulation test network
// for validity by checking the encoded bytes and hashes.
func TestSimNetGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := SimNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestSimNetGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), simNetGenesisBlockBytes) {
		t.Fatalf("TestSimNetGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(simNetGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := SimNetParams.GenesisBlock.BlockHash()
	if !SimNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestSimNetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(SimNetParams.GenesisHash))
	}
}

// genesisBlockBytes are the wire encoded bytes for the genesis block of the
// main network as of protocol version 60002.
var genesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x3B, 0xA3, 0xED, 0xFD, 0x7A, 0x7B, 0x12, 0xB2, 0x7A, 0xC7, 0x2C,
	0x3E, 0x67, 0x76, 0x8F, 0x61, 0x7F, 0xC8, 0x1B, 0xC3, 0x88, 0x8A, 0x51, 0x32, 0x3A, 0x9F, 0xB8,
	0xAA, 0x4B, 0x1E, 0x5E, 0x4A, 0xC8, 0xC4, 0x28, 0x5B, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0x00,
	0x1E, 0x50, 0x25, 0x19, 0xC0, 0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0x4D,
	0x04, 0xFF, 0xFF, 0x00, 0x1D, 0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6D, 0x65,
	0x73, 0x20, 0x30, 0x33, 0x2F, 0x4A, 0x61, 0x6E, 0x2F, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68,
	0x61, 0x6E, 0x63, 0x65, 0x6C, 0x6C, 0x6F, 0x72, 0x20, 0x6F, 0x6E, 0x20, 0x62, 0x72, 0x69, 0x6E,
	0x6B, 0x20, 0x6F, 0x66, 0x20, 0x73, 0x65, 0x63, 0x6F, 0x6E, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6C,
	0x6F, 0x75, 0x74, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x62, 0x61, 0x6E, 0x6B, 0x73, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x01, 0x00, 0xF2, 0x05, 0x2A, 0x01, 0x00, 0x00, 0x00, 0x43, 0x41,
	0x04, 0x67, 0x8A, 0xFD, 0xB0, 0xFE, 0x55, 0x48, 0x27, 0x19, 0x67, 0xF1, 0xA6, 0x71, 0x30, 0xB7,
	0x10, 0x5C, 0xD6, 0xA8, 0x28, 0xE0, 0x39, 0x09, 0xA6, 0x79, 0x62, 0xE0, 0xEA, 0x1F, 0x61, 0xDE,
	0xB6, 0x49, 0xF6, 0xBC, 0x3F, 0x4C, 0xEF, 0x38, 0xC4, 0xF3, 0x55, 0x04, 0xE5, 0x1E, 0xC1, 0x12,
	0xDE, 0x5C, 0x38, 0x4D, 0xF7, 0xBA, 0x0B, 0x8D, 0x57, 0x8A, 0x4C, 0x70, 0x2B, 0x6B, 0xF1, 0x1D,
	0x5F, 0xAC, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

// regTestGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the regression test network as of protocol version 60002.
var regTestGenesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, /* |.....;..| */
	0xfd, 0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, /* |.z{..z.,| */
	0x3e, 0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, /* |>gv.a...| */
	0xc3, 0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, /* |...Q2:..| */
	0xaa, 0x4b, 0x1e, 0x5e, 0x4a, 0x36, 0xc6, 0x28, /* |.K.^J6.(| */
	0x5b, 0xff, 0xff, 0x7f, 0x20, 0x01, 0x00, 0x00, /* |[... ...| */
	0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, /* |........| */
	0xff, 0xff, 0xff, 0x4d, 0x04, 0xff, 0xff, 0x00, /* |...M....| */
	0x1d, 0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, /* |...EThe | */
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x20, 0x30, 0x33, /* |Times 03| */
	0x2f, 0x4a, 0x61, 0x6e, 0x2f, 0x32, 0x30, 0x30, /* |/Jan/200| */
	0x39, 0x20, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, /* |9 Chance| */
	0x6c, 0x6c, 0x6f, 0x72, 0x20, 0x6f, 0x6e, 0x20, /* |llor on | */
	0x62, 0x72, 0x69, 0x6e, 0x6b, 0x20, 0x6f, 0x66, /* |brink of| */
	0x20, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x20, /* | second | */
	0x62, 0x61, 0x69, 0x6c, 0x6f, 0x75, 0x74, 0x20, /* |bailout | */
	0x66, 0x6f, 0x72, 0x20, 0x62, 0x61, 0x6e, 0x6b, /* |for bank| */
	0x73, 0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0xf2, /* |s.......| */
	0x05, 0x2a, 0x01, 0x00, 0x00, 0x00, 0x43, 0x41, /* |.*....CA| */
	0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, 0x48, /* |.g....UH| */
	0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, 0xb7, /* |'.g..q0.| */
	0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, 0x09, /* |.\..(.9.| */
	0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, 0xde, /* |.yb...a.| */
	0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, 0x38, /* |.I..?L.8| */
	0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, 0x12, /* |..U.....| */
	0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, 0x8d, /* |.\8M....| */
	0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, 0x1d, /* |W.Lp+k..| */
	0x5f, 0xac, 0x00, 0x00, 0x00, 0x00, /* |_.....| */
}

// testNet3GenesisBlockBytes are the wire encoded bytes for the genesis block of
// the test network (version 3) as of protocol version 60002.
var testNet3GenesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, /* |.....;..| */
	0xfd, 0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, /* |.z{..z.,| */
	0x3e, 0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, /* |>gv.a...| */
	0xc3, 0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, /* |...Q2:..| */
	0xaa, 0x4b, 0x1e, 0x5e, 0x4a, 0x06, 0xc7, 0x28, /* |.K.^J..(| */
	0x5b, 0xff, 0xff, 0x00, 0x1e, 0x3b, 0x1b, 0x2f, /* |[....;./| */
	0x80, 0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, /* |........| */
	0xff, 0xff, 0xff, 0x4d, 0x04, 0xff, 0xff, 0x00, /* |...M....| */
	0x1d, 0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, /* |...EThe | */
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x20, 0x30, 0x33, /* |Times 03| */
	0x2f, 0x4a, 0x61, 0x6e, 0x2f, 0x32, 0x30, 0x30, /* |/Jan/200| */
	0x39, 0x20, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, /* |9 Chance| */
	0x6c, 0x6c, 0x6f, 0x72, 0x20, 0x6f, 0x6e, 0x20, /* |llor on | */
	0x62, 0x72, 0x69, 0x6e, 0x6b, 0x20, 0x6f, 0x66, /* |brink of| */
	0x20, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x20, /* | second | */
	0x62, 0x61, 0x69, 0x6c, 0x6f, 0x75, 0x74, 0x20, /* |bailout | */
	0x66, 0x6f, 0x72, 0x20, 0x62, 0x61, 0x6e, 0x6b, /* |for bank| */
	0x73, 0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0xf2, /* |s.......| */
	0x05, 0x2a, 0x01, 0x00, 0x00, 0x00, 0x43, 0x41, /* |.*....CA| */
	0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, 0x48, /* |.g....UH| */
	0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, 0xb7, /* |'.g..q0.| */
	0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, 0x09, /* |.\..(.9.| */
	0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, 0xde, /* |.yb...a.| */
	0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, 0x38, /* |.I..?L.8| */
	0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, 0x12, /* |..U.....| */
	0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, 0x8d, /* |.\8M....| */
	0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, 0x1d, /* |W.Lp+k..| */
	0x5f, 0xac, 0x00, 0x00, 0x00, 0x00, /* |_.....| */
}

// simNetGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the simulation test network as of protocol version 70002.
var simNetGenesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, /* |.....;..| */
	0xfd, 0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, /* |.z{..z.,| */
	0x3e, 0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, /* |>gv.a...| */
	0xc3, 0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, /* |...Q2:..| */
	0xaa, 0x4b, 0x1e, 0x5e, 0x4a, 0xec, 0xc7, 0x28, /* |.K.^J..(| */
	0x5b, 0xff, 0xff, 0x7f, 0x20, 0xfb, 0xff, 0xff, /* |[... ...| */
	0x9f, 0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, /* |........| */
	0xff, 0xff, 0xff, 0x4d, 0x04, 0xff, 0xff, 0x00, /* |...M....| */
	0x1d, 0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, /* |...EThe | */
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x20, 0x30, 0x33, /* |Times 03| */
	0x2f, 0x4a, 0x61, 0x6e, 0x2f, 0x32, 0x30, 0x30, /* |/Jan/200| */
	0x39, 0x20, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, /* |9 Chance| */
	0x6c, 0x6c, 0x6f, 0x72, 0x20, 0x6f, 0x6e, 0x20, /* |llor on | */
	0x62, 0x72, 0x69, 0x6e, 0x6b, 0x20, 0x6f, 0x66, /* |brink of| */
	0x20, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x20, /* | second | */
	0x62, 0x61, 0x69, 0x6c, 0x6f, 0x75, 0x74, 0x20, /* |bailout | */
	0x66, 0x6f, 0x72, 0x20, 0x62, 0x61, 0x6e, 0x6b, /* |for bank| */
	0x73, 0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0xf2, /* |s.......| */
	0x05, 0x2a, 0x01, 0x00, 0x00, 0x00, 0x43, 0x41, /* |.*....CA| */
	0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, 0x48, /* |.g....UH| */
	0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, 0xb7, /* |'.g..q0.| */
	0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, 0x09, /* |.\..(.9.| */
	0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, 0xde, /* |.yb...a.| */
	0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, 0x38, /* |.I..?L.8| */
	0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, 0x12, /* |..U.....| */
	0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, 0x8d, /* |.\8M....| */
	0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, 0x1d, /* |W.Lp+k..| */
	0x5f, 0xac, 0x00, 0x00, 0x00, 0x00, /* |_.....| */
}
