// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package drip

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitVaultProtoConfig is the `initVaultProtoConfig` instruction.
type InitVaultProtoConfig struct {
	Params *InitializeVaultProtoConfigParams

	// [0] = [WRITE, SIGNER] creator
	//
	// [1] = [WRITE, SIGNER] vaultProtoConfig
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitVaultProtoConfigInstructionBuilder creates a new `InitVaultProtoConfig` instruction builder.
func NewInitVaultProtoConfigInstructionBuilder() *InitVaultProtoConfig {
	nd := &InitVaultProtoConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *InitVaultProtoConfig) SetParams(params InitializeVaultProtoConfigParams) *InitVaultProtoConfig {
	inst.Params = &params
	return inst
}

// SetCreatorAccount sets the "creator" account.
func (inst *InitVaultProtoConfig) SetCreatorAccount(creator ag_solanago.PublicKey) *InitVaultProtoConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(creator).WRITE().SIGNER()
	return inst
}

// GetCreatorAccount gets the "creator" account.
func (inst *InitVaultProtoConfig) GetCreatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetVaultProtoConfigAccount sets the "vaultProtoConfig" account.
func (inst *InitVaultProtoConfig) SetVaultProtoConfigAccount(vaultProtoConfig ag_solanago.PublicKey) *InitVaultProtoConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(vaultProtoConfig).WRITE().SIGNER()
	return inst
}

// GetVaultProtoConfigAccount gets the "vaultProtoConfig" account.
func (inst *InitVaultProtoConfig) GetVaultProtoConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitVaultProtoConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitVaultProtoConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitVaultProtoConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst InitVaultProtoConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitVaultProtoConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitVaultProtoConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitVaultProtoConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Creator is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.VaultProtoConfig is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *InitVaultProtoConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitVaultProtoConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         creator", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("vaultProtoConfig", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   systemProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj InitVaultProtoConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitVaultProtoConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewInitVaultProtoConfigInstruction declares a new InitVaultProtoConfig instruction with the provided parameters and accounts.
func NewInitVaultProtoConfigInstruction(
	// Parameters:
	params InitializeVaultProtoConfigParams,
	// Accounts:
	creator ag_solanago.PublicKey,
	vaultProtoConfig ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitVaultProtoConfig {
	return NewInitVaultProtoConfigInstructionBuilder().
		SetParams(params).
		SetCreatorAccount(creator).
		SetVaultProtoConfigAccount(vaultProtoConfig).
		SetSystemProgramAccount(systemProgram)
}
