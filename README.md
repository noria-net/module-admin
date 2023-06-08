# Module Admin

## Overview

The admin module is a cosmos sdk module that allows a governance-whitelisted account to perform privileged operations on the blockchain. It is useful to perform state changes with a single transaction instead of having to go through governance.

## Current Features

- [UpdateAlliance](./x/admin/keeper/msg_server_update_alliance.go) (requires the [`Alliance`](https://github.com/terra-money/alliance) module)

## Integration

```go
// go.mod

require (
  github.com/noria-net/module-admin vX.X.X
)
```

```go
// app.go

import (
  adminmodule "github.com/noria-net/module-admin/x/admin"
  adminmodulekeeper "github.com/noria-net/module-admin/x/admin/keeper"
  adminmoduletypes "github.com/noria-net/module-admin/x/admin/types"
)

...

// add the AppModuleBasic to the ModuleBasics Manager
ModuleBasics = module.NewBasicManager(
  ...
  adminmodule.AppModuleBasic{},
)

// Add an AdminKeeper to your app struct
type MyApp struct {
  ...
  AdminKeeper         adminmodulekeeper.Keeper
}

// Instiantiate the AdminKeeper in your app constructor
app.AdminKeeper = *adminmodulekeeper.NewKeeper(
  appCodec,
  keys[adminmoduletypes.StoreKey],
  keys[adminmoduletypes.MemStoreKey],
  app.GetSubspace(adminmoduletypes.ModuleName),

  app.AllianceKeeper,
)
adminModule := adminmodule.NewAppModule(appCodec, app.AdminKeeper, app.AccountKeeper, app.BankKeeper)

// Add the adminModule to the ModuleManager

app.ModuleManager = module.NewManager(
  ...
  adminModule,
  ... // crisis
)

// Add the module name to the begin blockers
app.ModuleManager.SetOrderBeginBlockers(
  ...
  adminmoduletypes.ModuleName,
)

// Add the module name to the end blockers
app.ModuleManager.SetOrderEndBlockers(
  ...
  adminmoduletypes.ModuleName,
)

// Add the module name to the genesis order (init)
app.ModuleManager.SetOrderInitGenesis(
  ...
  adminmoduletypes.ModuleName,
)

// Add the module name to the genesis order (export)
app.ModuleManager.SetOrderExportGenesis(
  ...
  adminmoduletypes.ModuleName,
)

// Add the module name to the params keeper subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
  paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

  ...
  paramsKeeper.Subspace(adminmoduletypes.ModuleName)
  ...
}

```

## Upgrade

When integrating this module into an existing chain, refer to [upgrades.go](./app/upgrades.go) for the upgrade logic.

## Proto

Protobuf files can be found on [buf.build](https://buf.build/github.com/noria-net/admin).

## Params

The module contains a single params named `admin` that equals to the list of whitelisted addresses that can perform privileged operations.
