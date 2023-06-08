package app

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	adminmodule "github.com/noria-net/module-admin/x/admin"
	adminmoduletypes "github.com/noria-net/module-admin/x/admin/types"
)

const UpgradeName = "add_admin"

func (app WasmApp) RegisterUpgradeHandlers() {
	// Set param key table for params module migration

	app.UpgradeKeeper.SetUpgradeHandler(
		UpgradeName,

		func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

			// Set new params for the admin module
			// Set the initial admin address
			newAdminParams := adminmoduletypes.NewParams("multisig_address_here")

			newAdminGenesis := adminmoduletypes.GenesisState{
				Params: newAdminParams,
			}
			encoded, err := app.appCodec.MarshalJSON(&newAdminGenesis)
			if err != nil {
				return nil, err
			}

			fromVM[adminmoduletypes.ModuleName] = adminmodule.AppModule{}.ConsensusVersion()
			module := app.ModuleManager.Modules[adminmoduletypes.ModuleName].(adminmodule.AppModule)
			module.InitGenesis(ctx, app.appCodec, encoded)

			return app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
		},
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				adminmoduletypes.ModuleName,
			},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
