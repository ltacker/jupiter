package types

import (
	"fmt"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId: PortID,
		// this line is used by starport scaffolding # genesis/types/default
		HohoList:     []*Hoho{},
		HihiList:     []*Hihi{},
		HahaList:     []*Haha{},
		HohosentList: []*Hohosent{},
		HihisentList: []*Hihisent{},
		HahasentList: []*Hahasent{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in hoho
	hohoIdMap := make(map[string]bool)

	for _, elem := range gs.HohoList {
		if _, ok := hohoIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for hoho")
		}
		hohoIdMap[elem.Id] = true
	}
	// Check for duplicated ID in hihi
	hihiIdMap := make(map[string]bool)

	for _, elem := range gs.HihiList {
		if _, ok := hihiIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for hihi")
		}
		hihiIdMap[elem.Id] = true
	}
	// Check for duplicated ID in haha
	hahaIdMap := make(map[string]bool)

	for _, elem := range gs.HahaList {
		if _, ok := hahaIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for haha")
		}
		hahaIdMap[elem.Id] = true
	}
	// Check for duplicated ID in hohosent
	hohosentIdMap := make(map[string]bool)

	for _, elem := range gs.HohosentList {
		if _, ok := hohosentIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for hohosent")
		}
		hohosentIdMap[elem.Id] = true
	}
	// Check for duplicated ID in hihisent
	hihisentIdMap := make(map[string]bool)

	for _, elem := range gs.HihisentList {
		if _, ok := hihisentIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for hihisent")
		}
		hihisentIdMap[elem.Id] = true
	}
	// Check for duplicated ID in hahasent
	hahasentIdMap := make(map[string]bool)

	for _, elem := range gs.HahasentList {
		if _, ok := hahasentIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for hahasent")
		}
		hahasentIdMap[elem.Id] = true
	}

	return nil
}
