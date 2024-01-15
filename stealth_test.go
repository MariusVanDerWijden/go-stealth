package gostealth

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestStaticTestCase(t *testing.T) {
	// Testdata taken from https://github.com/nerolation/executable-stealth-address-specs/blob/main/test.ipynb
	// Sender
	aPrivNum, _ := new(big.Int).SetString("31582853143040820948875942041653389873450407831047855470517498178324574486065", 10)
	aPriv, err := crypto.HexToECDSA(aPrivNum.Text(16))
	if err != nil {
		t.Fatal(err)
	}
	wantPrivX, _ := new(big.Int).SetString("99931485108758068354634100015529707565438847495649276196131125998359569029703", 10)
	wantPrivY, _ := new(big.Int).SetString("4744375390796532504618795785909610189099640957761399522523575349957196497592", 10)

	if aPriv.PublicKey.X.Cmp(wantPrivX) != 0 || aPriv.PublicKey.Y.Cmp(wantPrivY) != 0 {
		t.Fatal("PK does not match")
	}

	// Recipient
	bPrivSpNum, _ := new(big.Int).SetString("30787322447577792890566286485782027903969759412226064433999487819529647462924", 10)
	bPrivSp, _ := crypto.HexToECDSA(bPrivSpNum.Text(16))
	bPrivScNum, _ := new(big.Int).SetString("50431308649801251425320023123245644035351225602185776979597242007527042324186", 10)
	bPrivSc, _ := crypto.HexToECDSA(bPrivScNum.Text(16))

	wantStealthMA := "st:eth:0x03e017e9d9dbcb9ce5771acfce74c95bc0eafb5db37ef4b1ac62375f8e7a4c8aef021ba1833a9575bd2ad924440a20a80417437f77b0539cbc3f5bbaeeb2881efe04"
	stealthMA := StealthMetaAddress(&bPrivSp.PublicKey, &bPrivSc.PublicKey)

	if stealthMA != wantStealthMA {
		t.Fatalf("stealthMA does not match, got %v want %v", stealthMA, wantStealthMA)
	}

	spTest, scTest, err := ParseStealthAddress(stealthMA)
	if err != nil {
		t.Fatal("failed to parse stealth address")
	}
	if !spTest.Equal(&bPrivSp.PublicKey) || !scTest.Equal(&bPrivSc.PublicKey) {
		t.Fatal("parsed keys do not match")
	}

	// Sender computes stealth address and necessary information
	stAddress, ephemeralPK, viewTag, err := GenStealthAddress(stealthMA, aPriv)
	if err != nil {
		t.Fatal(err)
	}
	wantStAddr := "0xc7c1bbf258340e551061e7d561798555aa871c0d"
	wantEphPKX, _ := new(big.Int).SetString("99931485108758068354634100015529707565438847495649276196131125998359569029703", 10)
	wantEphPKY, _ := new(big.Int).SetString("4744375390796532504618795785909610189099640957761399522523575349957196497592", 10)
	wantViewTag := 61

	if strings.ToLower(stAddress.String()) != wantStAddr {
		t.Fatal("invalid stAddr")
	}
	if ephemeralPK.X.Cmp(wantEphPKX) != 0 || ephemeralPK.Y.Cmp(wantEphPKY) != 0 {
		t.Fatal("invalid ephemeral pk")
	}
	if viewTag != ViewTag(wantViewTag) {
		t.Fatal("invalid view tag")
	}

	dhSecret, err := ParseEventView(ephemeralPK, *stAddress, viewTag, bPrivSc, &bPrivSp.PublicKey)
	if err != nil {
		t.Fatal(err)
	}

	wantStealthSK, _ := new(big.Int).SetString("0x81c527d561a196132fe18f2242385e4cdac91990657021cd0cee71a24d55242e", 0)
	stealthSK, _, err := stealthAddrPrivKey(dhSecret, bPrivSp)
	if err != nil {
		t.Fatal(err)
	}
	if stealthSK.Cmp(wantStealthSK) != 0 {
		t.Fatal("invalid stealth private key")
	}
}
