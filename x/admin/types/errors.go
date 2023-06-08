package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/admin module sentinel errors
var (
	ErrUpdateAlliance      = sdkerrors.Register(ModuleName, 1100, "error updating alliance")
	ErrAdminNotInitialized = sdkerrors.Register(ModuleName, 1101, "error admin not initialized")
)
