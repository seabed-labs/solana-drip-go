// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drip

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitVaultPeriod is the `initVaultPeriod` instruction.
type InitVaultPeriod struct {
	Params *InitializeVaultPeriodParams

	// [0] = [WRITE] vaultPeriod
	//
	// [1] = [] vault
	//
	// [2] = [] tokenAMint
	//
	// [3] = [] tokenBMint
	//
	// [4] = [] vaultProtoConfig
	//
	// [5] = [WRITE, SIGNER] creator
	//
	// [6] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitVaultPeriodInstructionBuilder creates a new `InitVaultPeriod` instruction builder.
func NewInitVaultPeriodInstructionBuilder() *InitVaultPeriod {
	nd := &InitVaultPeriod{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *InitVaultPeriod) SetParams(params InitializeVaultPeriodParams) *InitVaultPeriod {
	inst.Params = &params
	return inst
}

// SetVaultPeriodAccount sets the "vaultPeriod" account.
func (inst *InitVaultPeriod) SetVaultPeriodAccount(vaultPeriod ag_solanago.PublicKey) *InitVaultPeriod {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(vaultPeriod).WRITE()
	return inst
}

// GetVaultPeriodAccount gets the "vaultPeriod" account.
func (inst *InitVaultPeriod) GetVaultPeriodAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetVaultAccount sets the "vault" account.
func (inst *InitVaultPeriod) SetVaultAccount(vault ag_solanago.PublicKey) *InitVaultPeriod {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(vault)
	return inst
}

// GetVaultAccount gets the "vault" account.
func (inst *InitVaultPeriod) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenAMintAccount sets the "tokenAMint" account.
func (inst *InitVaultPeriod) SetTokenAMintAccount(tokenAMint ag_solanago.PublicKey) *InitVaultPeriod {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenAMint)
	return inst
}

// GetTokenAMintAccount gets the "tokenAMint" account.
func (inst *InitVaultPeriod) GetTokenAMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenBMintAccount sets the "tokenBMint" account.
func (inst *InitVaultPeriod) SetTokenBMintAccount(tokenBMint ag_solanago.PublicKey) *InitVaultPeriod {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenBMint)
	return inst
}

// GetTokenBMintAccount gets the "tokenBMint" account.
func (inst *InitVaultPeriod) GetTokenBMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetVaultProtoConfigAccount sets the "vaultProtoConfig" account.
func (inst *InitVaultPeriod) SetVaultProtoConfigAccount(vaultProtoConfig ag_solanago.PublicKey) *InitVaultPeriod {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(vaultProtoConfig)
	return inst
}

// GetVaultProtoConfigAccount gets the "vaultProtoConfig" account.
func (inst *InitVaultPeriod) GetVaultProtoConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetCreatorAccount sets the "creator" account.
func (inst *InitVaultPeriod) SetCreatorAccount(creator ag_solanago.PublicKey) *InitVaultPeriod {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(creator).WRITE().SIGNER()
	return inst
}

// GetCreatorAccount gets the "creator" account.
func (inst *InitVaultPeriod) GetCreatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitVaultPeriod) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitVaultPeriod {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitVaultPeriod) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst InitVaultPeriod) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitVaultPeriod,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitVaultPeriod) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitVaultPeriod) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.VaultPeriod is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Vault is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenAMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenBMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.VaultProtoConfig is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Creator is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *InitVaultPeriod) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitVaultPeriod")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     vaultPeriod", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           vault", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      tokenAMint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      tokenBMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("vaultProtoConfig", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         creator", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("   systemProgram", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj InitVaultPeriod) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitVaultPeriod) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewInitVaultPeriodInstruction declares a new InitVaultPeriod instruction with the provided parameters and accounts.
func NewInitVaultPeriodInstruction(
	// Parameters:
	params InitializeVaultPeriodParams,
	// Accounts:
	vaultPeriod ag_solanago.PublicKey,
	vault ag_solanago.PublicKey,
	tokenAMint ag_solanago.PublicKey,
	tokenBMint ag_solanago.PublicKey,
	vaultProtoConfig ag_solanago.PublicKey,
	creator ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitVaultPeriod {
	return NewInitVaultPeriodInstructionBuilder().
		SetParams(params).
		SetVaultPeriodAccount(vaultPeriod).
		SetVaultAccount(vault).
		SetTokenAMintAccount(tokenAMint).
		SetTokenBMintAccount(tokenBMint).
		SetVaultProtoConfigAccount(vaultProtoConfig).
		SetCreatorAccount(creator).
		SetSystemProgramAccount(systemProgram)
}
