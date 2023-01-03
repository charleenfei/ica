package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		WhoisList:         []Whois{},
		Testmin:           nil,
		PendingBuyList:    []PendingBuy{},
		PendingSellList:   []PendingSell{},
		CmpHostResultList: []CmpHostResult{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in whois
	whoisIndexMap := make(map[string]struct{})

	for _, elem := range gs.WhoisList {
		index := string(WhoisKey(elem.Index))
		if _, ok := whoisIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for whois")
		}
		whoisIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pendingBuy
	pendingBuyIndexMap := make(map[string]struct{})

	for _, elem := range gs.PendingBuyList {
		index := string(PendingBuyKey(elem.Index))
		if _, ok := pendingBuyIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pendingBuy")
		}
		pendingBuyIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pendingSell
	pendingSellIndexMap := make(map[string]struct{})

	for _, elem := range gs.PendingSellList {
		index := string(PendingSellKey(elem.Name))
		if _, ok := pendingSellIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pendingSell")
		}
		pendingSellIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in cmpHostResult
	cmpHostResultIndexMap := make(map[string]struct{})

	for _, elem := range gs.CmpHostResultList {
		index := string(CmpHostResultKey(elem.Index))
		if _, ok := cmpHostResultIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for cmpHostResult")
		}
		cmpHostResultIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
