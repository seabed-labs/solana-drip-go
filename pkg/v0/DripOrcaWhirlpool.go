// Code generated by https://github.com/dcaf-mocha/anchor-go. DO NOT EDIT.

package drip

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// DripOrcaWhirlpool is the `dripOrcaWhirlpool` instruction.
type DripOrcaWhirlpool struct {

	// [0] = [SIGNER] dripTriggerSource
	//
	// [1] = [WRITE] vault
	//
	// [2] = [] vaultProtoConfig
	//
	// [3] = [] lastVaultPeriod
	//
	// [4] = [WRITE] currentVaultPeriod
	//
	// [5] = [WRITE] vaultTokenAAccount
	//
	// [6] = [WRITE] vaultTokenBAccount
	//
	// [7] = [WRITE] swapTokenAAccount
	//
	// [8] = [WRITE] swapTokenBAccount
	//
	// [9] = [WRITE] dripFeeTokenAAccount
	//
	// [10] = [] tokenProgram
	//
	// [11] = [] associatedTokenProgram
	//
	// [12] = [] whirlpoolProgram
	//
	// [13] = [] systemProgram
	//
	// [14] = [] rent
	//
	// [15] = [WRITE] whirlpool
	//
	// [16] = [WRITE] tickArray0
	//
	// [17] = [WRITE] tickArray1
	//
	// [18] = [WRITE] tickArray2
	//
	// [19] = [] oracle
	ag_solanago.AccountMetaSlice `bin:"-"`
}

type DripOrcaWhirlpoolAccounts struct {
	DripTriggerSource      ag_solanago.PublicKey
	Vault                  ag_solanago.PublicKey
	VaultProtoConfig       ag_solanago.PublicKey
	LastVaultPeriod        ag_solanago.PublicKey
	CurrentVaultPeriod     ag_solanago.PublicKey
	VaultTokenAAccount     ag_solanago.PublicKey
	VaultTokenBAccount     ag_solanago.PublicKey
	SwapTokenAAccount      ag_solanago.PublicKey
	SwapTokenBAccount      ag_solanago.PublicKey
	DripFeeTokenAAccount   ag_solanago.PublicKey
	TokenProgram           ag_solanago.PublicKey
	AssociatedTokenProgram ag_solanago.PublicKey
	WhirlpoolProgram       ag_solanago.PublicKey
	SystemProgram          ag_solanago.PublicKey
	Rent                   ag_solanago.PublicKey
	Whirlpool              ag_solanago.PublicKey
	TickArray0             ag_solanago.PublicKey
	TickArray1             ag_solanago.PublicKey
	TickArray2             ag_solanago.PublicKey
	Oracle                 ag_solanago.PublicKey
}

// NewDripOrcaWhirlpoolInstructionBuilder creates a new `DripOrcaWhirlpool` instruction builder.
func NewDripOrcaWhirlpoolInstructionBuilder() *DripOrcaWhirlpool {
	nd := &DripOrcaWhirlpool{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 20),
	}
	return nd
}

func (inst *DripOrcaWhirlpool) GetDripOrcaWhirlpoolAccounts() *DripOrcaWhirlpoolAccounts {
	res := &DripOrcaWhirlpoolAccounts{}
	if inst.AccountMetaSlice[0] != nil {
		res.DripTriggerSource = inst.AccountMetaSlice[0].PublicKey
	}
	if inst.AccountMetaSlice[1] != nil {
		res.Vault = inst.AccountMetaSlice[1].PublicKey
	}
	if inst.AccountMetaSlice[2] != nil {
		res.VaultProtoConfig = inst.AccountMetaSlice[2].PublicKey
	}
	if inst.AccountMetaSlice[3] != nil {
		res.LastVaultPeriod = inst.AccountMetaSlice[3].PublicKey
	}
	if inst.AccountMetaSlice[4] != nil {
		res.CurrentVaultPeriod = inst.AccountMetaSlice[4].PublicKey
	}
	if inst.AccountMetaSlice[5] != nil {
		res.VaultTokenAAccount = inst.AccountMetaSlice[5].PublicKey
	}
	if inst.AccountMetaSlice[6] != nil {
		res.VaultTokenBAccount = inst.AccountMetaSlice[6].PublicKey
	}
	if inst.AccountMetaSlice[7] != nil {
		res.SwapTokenAAccount = inst.AccountMetaSlice[7].PublicKey
	}
	if inst.AccountMetaSlice[8] != nil {
		res.SwapTokenBAccount = inst.AccountMetaSlice[8].PublicKey
	}
	if inst.AccountMetaSlice[9] != nil {
		res.DripFeeTokenAAccount = inst.AccountMetaSlice[9].PublicKey
	}
	if inst.AccountMetaSlice[10] != nil {
		res.TokenProgram = inst.AccountMetaSlice[10].PublicKey
	}
	if inst.AccountMetaSlice[11] != nil {
		res.AssociatedTokenProgram = inst.AccountMetaSlice[11].PublicKey
	}
	if inst.AccountMetaSlice[12] != nil {
		res.WhirlpoolProgram = inst.AccountMetaSlice[12].PublicKey
	}
	if inst.AccountMetaSlice[13] != nil {
		res.SystemProgram = inst.AccountMetaSlice[13].PublicKey
	}
	if inst.AccountMetaSlice[14] != nil {
		res.Rent = inst.AccountMetaSlice[14].PublicKey
	}
	if inst.AccountMetaSlice[15] != nil {
		res.Whirlpool = inst.AccountMetaSlice[15].PublicKey
	}
	if inst.AccountMetaSlice[16] != nil {
		res.TickArray0 = inst.AccountMetaSlice[16].PublicKey
	}
	if inst.AccountMetaSlice[17] != nil {
		res.TickArray1 = inst.AccountMetaSlice[17].PublicKey
	}
	if inst.AccountMetaSlice[18] != nil {
		res.TickArray2 = inst.AccountMetaSlice[18].PublicKey
	}
	if inst.AccountMetaSlice[19] != nil {
		res.Oracle = inst.AccountMetaSlice[19].PublicKey
	}
	return res
}

// SetDripTriggerSourceAccount sets the "dripTriggerSource" account.
func (inst *DripOrcaWhirlpool) SetDripTriggerSourceAccount(dripTriggerSource ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(dripTriggerSource).SIGNER()
	return inst
}

// GetDripTriggerSourceAccount gets the "dripTriggerSource" account.
func (inst *DripOrcaWhirlpool) GetDripTriggerSourceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetVaultAccount sets the "vault" account.
func (inst *DripOrcaWhirlpool) SetVaultAccount(vault ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(vault).WRITE()
	return inst
}

// GetVaultAccount gets the "vault" account.
func (inst *DripOrcaWhirlpool) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetVaultProtoConfigAccount sets the "vaultProtoConfig" account.
func (inst *DripOrcaWhirlpool) SetVaultProtoConfigAccount(vaultProtoConfig ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(vaultProtoConfig)
	return inst
}

// GetVaultProtoConfigAccount gets the "vaultProtoConfig" account.
func (inst *DripOrcaWhirlpool) GetVaultProtoConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetLastVaultPeriodAccount sets the "lastVaultPeriod" account.
func (inst *DripOrcaWhirlpool) SetLastVaultPeriodAccount(lastVaultPeriod ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(lastVaultPeriod)
	return inst
}

// GetLastVaultPeriodAccount gets the "lastVaultPeriod" account.
func (inst *DripOrcaWhirlpool) GetLastVaultPeriodAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetCurrentVaultPeriodAccount sets the "currentVaultPeriod" account.
func (inst *DripOrcaWhirlpool) SetCurrentVaultPeriodAccount(currentVaultPeriod ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(currentVaultPeriod).WRITE()
	return inst
}

// GetCurrentVaultPeriodAccount gets the "currentVaultPeriod" account.
func (inst *DripOrcaWhirlpool) GetCurrentVaultPeriodAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetVaultTokenAAccountAccount sets the "vaultTokenAAccount" account.
func (inst *DripOrcaWhirlpool) SetVaultTokenAAccountAccount(vaultTokenAAccount ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(vaultTokenAAccount).WRITE()
	return inst
}

// GetVaultTokenAAccountAccount gets the "vaultTokenAAccount" account.
func (inst *DripOrcaWhirlpool) GetVaultTokenAAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetVaultTokenBAccountAccount sets the "vaultTokenBAccount" account.
func (inst *DripOrcaWhirlpool) SetVaultTokenBAccountAccount(vaultTokenBAccount ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(vaultTokenBAccount).WRITE()
	return inst
}

// GetVaultTokenBAccountAccount gets the "vaultTokenBAccount" account.
func (inst *DripOrcaWhirlpool) GetVaultTokenBAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSwapTokenAAccountAccount sets the "swapTokenAAccount" account.
func (inst *DripOrcaWhirlpool) SetSwapTokenAAccountAccount(swapTokenAAccount ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(swapTokenAAccount).WRITE()
	return inst
}

// GetSwapTokenAAccountAccount gets the "swapTokenAAccount" account.
func (inst *DripOrcaWhirlpool) GetSwapTokenAAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSwapTokenBAccountAccount sets the "swapTokenBAccount" account.
func (inst *DripOrcaWhirlpool) SetSwapTokenBAccountAccount(swapTokenBAccount ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(swapTokenBAccount).WRITE()
	return inst
}

// GetSwapTokenBAccountAccount gets the "swapTokenBAccount" account.
func (inst *DripOrcaWhirlpool) GetSwapTokenBAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetDripFeeTokenAAccountAccount sets the "dripFeeTokenAAccount" account.
func (inst *DripOrcaWhirlpool) SetDripFeeTokenAAccountAccount(dripFeeTokenAAccount ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(dripFeeTokenAAccount).WRITE()
	return inst
}

// GetDripFeeTokenAAccountAccount gets the "dripFeeTokenAAccount" account.
func (inst *DripOrcaWhirlpool) GetDripFeeTokenAAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *DripOrcaWhirlpool) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *DripOrcaWhirlpool) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *DripOrcaWhirlpool) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *DripOrcaWhirlpool) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetWhirlpoolProgramAccount sets the "whirlpoolProgram" account.
func (inst *DripOrcaWhirlpool) SetWhirlpoolProgramAccount(whirlpoolProgram ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(whirlpoolProgram)
	return inst
}

// GetWhirlpoolProgramAccount gets the "whirlpoolProgram" account.
func (inst *DripOrcaWhirlpool) GetWhirlpoolProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *DripOrcaWhirlpool) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *DripOrcaWhirlpool) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetRentAccount sets the "rent" account.
func (inst *DripOrcaWhirlpool) SetRentAccount(rent ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *DripOrcaWhirlpool) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetWhirlpoolAccount sets the "whirlpool" account.
func (inst *DripOrcaWhirlpool) SetWhirlpoolAccount(whirlpool ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(whirlpool).WRITE()
	return inst
}

// GetWhirlpoolAccount gets the "whirlpool" account.
func (inst *DripOrcaWhirlpool) GetWhirlpoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetTickArray0Account sets the "tickArray0" account.
func (inst *DripOrcaWhirlpool) SetTickArray0Account(tickArray0 ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(tickArray0).WRITE()
	return inst
}

// GetTickArray0Account gets the "tickArray0" account.
func (inst *DripOrcaWhirlpool) GetTickArray0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetTickArray1Account sets the "tickArray1" account.
func (inst *DripOrcaWhirlpool) SetTickArray1Account(tickArray1 ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(tickArray1).WRITE()
	return inst
}

// GetTickArray1Account gets the "tickArray1" account.
func (inst *DripOrcaWhirlpool) GetTickArray1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetTickArray2Account sets the "tickArray2" account.
func (inst *DripOrcaWhirlpool) SetTickArray2Account(tickArray2 ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(tickArray2).WRITE()
	return inst
}

// GetTickArray2Account gets the "tickArray2" account.
func (inst *DripOrcaWhirlpool) GetTickArray2Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

// SetOracleAccount sets the "oracle" account.
func (inst *DripOrcaWhirlpool) SetOracleAccount(oracle ag_solanago.PublicKey) *DripOrcaWhirlpool {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(oracle)
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *DripOrcaWhirlpool) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(19)
}

func (inst DripOrcaWhirlpool) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DripOrcaWhirlpool,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DripOrcaWhirlpool) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DripOrcaWhirlpool) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.DripTriggerSource is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Vault is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.VaultProtoConfig is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.LastVaultPeriod is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.CurrentVaultPeriod is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.VaultTokenAAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.VaultTokenBAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SwapTokenAAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SwapTokenBAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.DripFeeTokenAAccount is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.WhirlpoolProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.Whirlpool is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.TickArray0 is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.TickArray1 is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.TickArray2 is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.Oracle is not set")
		}
	}
	return nil
}

func (inst *DripOrcaWhirlpool) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DripOrcaWhirlpool")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=20]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     dripTriggerSource", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                 vault", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      vaultProtoConfig", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       lastVaultPeriod", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    currentVaultPeriod", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           vaultTokenA", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           vaultTokenB", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("            swapTokenA", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("            swapTokenB", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("         dripFeeTokenA", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("      whirlpoolProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("             whirlpool", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("            tickArray0", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("            tickArray1", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("            tickArray2", inst.AccountMetaSlice.Get(18)))
						accountsBranch.Child(ag_format.Meta("                oracle", inst.AccountMetaSlice.Get(19)))
					})
				})
		})
}

func (obj DripOrcaWhirlpool) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *DripOrcaWhirlpool) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewDripOrcaWhirlpoolInstruction declares a new DripOrcaWhirlpool instruction with the provided parameters and accounts.
func NewDripOrcaWhirlpoolInstruction(
	// Accounts:
	dripTriggerSource ag_solanago.PublicKey,
	vault ag_solanago.PublicKey,
	vaultProtoConfig ag_solanago.PublicKey,
	lastVaultPeriod ag_solanago.PublicKey,
	currentVaultPeriod ag_solanago.PublicKey,
	vaultTokenAAccount ag_solanago.PublicKey,
	vaultTokenBAccount ag_solanago.PublicKey,
	swapTokenAAccount ag_solanago.PublicKey,
	swapTokenBAccount ag_solanago.PublicKey,
	dripFeeTokenAAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	whirlpoolProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	whirlpool ag_solanago.PublicKey,
	tickArray0 ag_solanago.PublicKey,
	tickArray1 ag_solanago.PublicKey,
	tickArray2 ag_solanago.PublicKey,
	oracle ag_solanago.PublicKey) *DripOrcaWhirlpool {
	return NewDripOrcaWhirlpoolInstructionBuilder().
		SetDripTriggerSourceAccount(dripTriggerSource).
		SetVaultAccount(vault).
		SetVaultProtoConfigAccount(vaultProtoConfig).
		SetLastVaultPeriodAccount(lastVaultPeriod).
		SetCurrentVaultPeriodAccount(currentVaultPeriod).
		SetVaultTokenAAccountAccount(vaultTokenAAccount).
		SetVaultTokenBAccountAccount(vaultTokenBAccount).
		SetSwapTokenAAccountAccount(swapTokenAAccount).
		SetSwapTokenBAccountAccount(swapTokenBAccount).
		SetDripFeeTokenAAccountAccount(dripFeeTokenAAccount).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetWhirlpoolProgramAccount(whirlpoolProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetWhirlpoolAccount(whirlpool).
		SetTickArray0Account(tickArray0).
		SetTickArray1Account(tickArray1).
		SetTickArray2Account(tickArray2).
		SetOracleAccount(oracle)
}
