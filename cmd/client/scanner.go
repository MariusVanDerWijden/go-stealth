package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	gostealth "github.com/MariusVanDerWijden/go-stealth"
	"github.com/MariusVanDerWijden/go-stealth/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ViewScan(client *ethclient.Client, contract *bindings.ERC5564Announcer, scanningKey *ecdsa.PrivateKey, spendingKey *ecdsa.PublicKey) error {
	sc := scanner{
		scanningKey:       scanningKey,
		spendingPublicKey: spendingKey,
		client:            client,
	}
	return sc.scan(0, contract)
}

func Scan(client *ethclient.Client, contract *bindings.ERC5564Announcer, scanningKey *ecdsa.PrivateKey, spendingKey *ecdsa.PrivateKey) error {
	sc := scanner{
		scanningKey: scanningKey,
		spendingKey: spendingKey,
		client:      client,
	}
	return sc.scan(0, contract)
}

type scanner struct {
	scanningKey       *ecdsa.PrivateKey
	spendingKey       *ecdsa.PrivateKey
	spendingPublicKey *ecdsa.PublicKey
	client            *ethclient.Client
}

func (s *scanner) scan(start uint64, contract *bindings.ERC5564Announcer) error {
	filterOpts := bind.FilterOpts{Start: start, End: nil, Context: context.Background()}
	schemeIDs := []*big.Int{new(big.Int)} // Secp256k1 has scheme id 0
	it, err := contract.FilterAnnouncement(&filterOpts, schemeIDs, nil, nil)
	if err != nil {
		return err
	}
	defer it.Close()
	for it.Next() {
		if err := s.handleEvent(it.Event); err != nil {
			return err
		}
	}
	return it.Error()
}

func (s *scanner) wait(start uint64, addresses []common.Address, callers []common.Address, contract *bindings.ERC5564Announcer) error {
	schemeIDs := []*big.Int{new(big.Int)} // Secp256k1 has scheme id 0
	sink := make(chan *bindings.ERC5564AnnouncerAnnouncement)
	sub, err := contract.WatchAnnouncement(&bind.WatchOpts{}, sink, schemeIDs, addresses, callers)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	for {
		ev := <-sink
		if err := s.handleEvent(ev); err != nil {
			return err
		}
	}
}

func (s *scanner) handleEvent(event *bindings.ERC5564AnnouncerAnnouncement) error {
	fmt.Printf("Found event: scheme: %v stealthAddr: %v caller: %v epheremeralPubKey: %v metadata: %v \n", event.SchemeId, event.StealthAddress, event.Caller, event.EphemeralPubKey, event.Metadata)
	ephemeralPK, err := crypto.DecompressPubkey(event.EphemeralPubKey)
	if err != nil {
		fmt.Println("could not decompress ephemeral public key")
		return nil
	}
	addr := event.StealthAddress
	if s.spendingKey != nil {
		secretKey, err := gostealth.ParseEvent(ephemeralPK, event.StealthAddress, gostealth.ViewTag(event.Metadata[0]), s.scanningKey, s.spendingKey)
		if err != nil {
			fmt.Println("Event was not meant for us")
			return nil
		}
		fmt.Printf("Stealth address found: %v spendable with %v\n", addr, secretKey)
	} else {
		dhSecret, err := gostealth.ParseEventView(ephemeralPK, event.StealthAddress, gostealth.ViewTag(event.Metadata[0]), s.scanningKey, s.spendingPublicKey)
		if err != nil {
			fmt.Println("Event was not meant for us")
			return nil
		}
		fmt.Printf("Stealth address found: %v spendable with secret %v and the spending sk\n", addr, dhSecret)
	}
	return checkAddress(addr, s.client)
}

func checkAddress(addr common.Address, client *ethclient.Client) error {
	// check if address has funds
	bal, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}
	if bal.Cmp(new(big.Int)) != 0 {
		fmt.Printf("Found a stealth address: %v with balance %v\n", addr, bal)
	}
	return nil
}
