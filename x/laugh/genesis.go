package laugh

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/jupiter/x/laugh/keeper"
	"github.com/ltacker/jupiter/x/laugh/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the hoho
	for _, elem := range genState.HohoList {
		k.SetHoho(ctx, *elem)
	}

	// Set hoho count
	k.SetHohoCount(ctx, int64(len(genState.HohoList)))

	// Set all the hihi
	for _, elem := range genState.HihiList {
		k.SetHihi(ctx, *elem)
	}

	// Set hihi count
	k.SetHihiCount(ctx, int64(len(genState.HihiList)))

	// Set all the haha
	for _, elem := range genState.HahaList {
		k.SetHaha(ctx, *elem)
	}

	// Set haha count
	k.SetHahaCount(ctx, int64(len(genState.HahaList)))

	// Set all the hohosent
	for _, elem := range genState.HohosentList {
		k.SetHohosent(ctx, *elem)
	}

	// Set hohosent count
	k.SetHohosentCount(ctx, int64(len(genState.HohosentList)))

	// Set all the hihisent
	for _, elem := range genState.HihisentList {
		k.SetHihisent(ctx, *elem)
	}

	// Set hihisent count
	k.SetHihisentCount(ctx, int64(len(genState.HihisentList)))

	// Set all the hahasent
	for _, elem := range genState.HahasentList {
		k.SetHahasent(ctx, *elem)
	}

	// Set hahasent count
	k.SetHahasentCount(ctx, int64(len(genState.HahasentList)))

	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all hoho
	hohoList := k.GetAllHoho(ctx)
	for _, elem := range hohoList {
		elem := elem
		genesis.HohoList = append(genesis.HohoList, &elem)
	}

	// Get all hihi
	hihiList := k.GetAllHihi(ctx)
	for _, elem := range hihiList {
		elem := elem
		genesis.HihiList = append(genesis.HihiList, &elem)
	}

	// Get all haha
	hahaList := k.GetAllHaha(ctx)
	for _, elem := range hahaList {
		elem := elem
		genesis.HahaList = append(genesis.HahaList, &elem)
	}

	// Get all hohosent
	hohosentList := k.GetAllHohosent(ctx)
	for _, elem := range hohosentList {
		elem := elem
		genesis.HohosentList = append(genesis.HohosentList, &elem)
	}

	// Get all hihisent
	hihisentList := k.GetAllHihisent(ctx)
	for _, elem := range hihisentList {
		elem := elem
		genesis.HihisentList = append(genesis.HihisentList, &elem)
	}

	// Get all hahasent
	hahasentList := k.GetAllHahasent(ctx)
	for _, elem := range hahasentList {
		elem := elem
		genesis.HahasentList = append(genesis.HahasentList, &elem)
	}

	genesis.PortId = k.GetPort(ctx)

	return genesis
}
