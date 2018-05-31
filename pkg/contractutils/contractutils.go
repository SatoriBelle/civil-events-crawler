// Package contractutils contains utilities related to smart contracts
// and testing smart contracts
package contractutils

import (
	"encoding/hex"
	log "github.com/golang/glog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joincivil/civil-events-crawler/pkg/generated/eip20"
	"github.com/joincivil/civil-events-crawler/pkg/generated/government"
	"github.com/joincivil/civil-events-crawler/pkg/generated/newsroom"
	"github.com/joincivil/civil-events-crawler/pkg/generated/parameterizer"
	"github.com/joincivil/civil-events-crawler/pkg/generated/plcr"
	"github.com/joincivil/civil-events-crawler/pkg/generated/tcr"
)

const (
	rinkebyAddress = "wss://rinkeby.infura.io/ws"

	minDeposit                       = 10
	pMinDeposit                      = 100
	applyStageLength                 = 18000
	pApplyStageLength                = 120
	commitStageLength                = 18000
	pCommitStageLength               = 120
	revealStageLength                = 18000
	pRevealStageLength               = 120
	dispensationPct                  = 50
	pDispensationPct                 = 50
	voteQuorum                       = 50
	pVoteQuorum                      = 50
	pProcessBy                       = 18000
	appealFeeAmount                  = 1000
	challengeAppealLength            = 18000
	requestAppealPhaseLength         = 36000
	judgeAppealPhaseLength           = 36000
	appealChallengeCommitStageLength = 16000
	appealChallengeRevealStageLength = 16000
	appealSupermajorityPercentage    = 66
)

// SetupRinkebyClient returns an instance of the ethclient setup on Rinkeby
func SetupRinkebyClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(rinkebyAddress)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// SetupSimulatedClient returns an an instance of the simulated backend.
func SetupSimulatedClient() (*backends.SimulatedBackend, *bind.TransactOpts) {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	genAlloc := make(core.GenesisAlloc)
	genAlloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(9223372036854775807)}
	sim := backends.NewSimulatedBackend(genAlloc)
	return sim, auth
}

// AllTestContracts contains the values returned from SetupAllTestContracts
type AllTestContracts struct {
	TokenAddr        common.Address
	TokenContract    *eip20.EIP20Contract
	PlcrAddr         common.Address
	PlcrContract     *plcr.PLCRVotingContract
	ParamAddr        common.Address
	ParamContract    *parameterizer.ParameterizerContract
	GovtAddr         common.Address
	GovtContract     *government.GovernmentContract
	CivilTcrAddr     common.Address
	CivilTcrContract *tcr.CivilTCRContract
	NewsroomAddr     common.Address
	NewsroomContract *newsroom.NewsroomContract
	Client           *backends.SimulatedBackend
	Auth             *bind.TransactOpts
}

// SetupAllTestContracts returns a struct with all the test contracts deployed to the
// simulated backend.
func SetupAllTestContracts() (*AllTestContracts, error) {
	client, auth := SetupSimulatedClient()
	tokenAddr, eip20, err := setupTestEIP20Contract(client, auth)
	if err != nil {
		log.Fatalf("Unable to deploy a test token: %v", err)
	}

	client.Commit()

	balance, _ := eip20.BalanceOf(&bind.CallOpts{}, auth.From)
	approveOpts := &bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}

	plcrAddr, plcr, err := setupTestPLCRVotingContract(client, auth, tokenAddr)
	if err != nil {
		log.Fatalf("Unable to deploy a test PLCR voting contract: err: %v", err)
		return nil, err
	}

	_, err = eip20.Approve(approveOpts, plcrAddr, balance)
	if err != nil {
		log.Errorf("Unable to approve user for PLCR: %v", err)
		return nil, err
	}

	client.Commit()

	paramAddr, param, err := setupTestParameterizerContract(client, auth, tokenAddr, plcrAddr)
	if err != nil {
		log.Fatalf("Unable to deploy a test parameterizer contract: err: %v", err)
		return nil, err
	}

	_, err = eip20.Approve(approveOpts, paramAddr, balance)
	if err != nil {
		log.Errorf("Unable to approve user for parameterizer: %v", err)
		return nil, err
	}

	client.Commit()

	govtAddr, govt, err := setupTestGovernmentContract(client, auth, auth.From, auth.From)
	if err != nil {
		log.Fatalf("Unable to deploy a test government contract: err: %v", err)
		return nil, err
	}

	_, err = eip20.Approve(approveOpts, govtAddr, balance)
	if err != nil {
		log.Errorf("Unable to approve user for govt: %v", err)
		return nil, err
	}

	client.Commit()

	civilTcrAddr, civilTcr, err := setupTestCivilTCRContract(client, auth, tokenAddr,
		plcrAddr, paramAddr, govtAddr)
	if err != nil {
		log.Fatalf("Unable to deploy a test civil tcr contract: err: %v", err)
		return nil, err
	}

	_, err = eip20.Approve(approveOpts, civilTcrAddr, balance)
	if err != nil {
		log.Errorf("Unable to approve user for tcr: %v", err)
		return nil, err
	}

	client.Commit()

	newsroomAddr, newsroom, err := setupTestNewsroomContract(client, auth)
	if err != nil {
		log.Fatalf("Unable to deploy a test newsroom contract: err: %v", err)
		return nil, err
	}

	_, err = eip20.Approve(approveOpts, newsroomAddr, balance)
	if err != nil {
		log.Errorf("Unable to approve user for newsroom: err: %v", err)
		return nil, err
	}

	client.Commit()

	return &AllTestContracts{
		TokenAddr:        tokenAddr,
		TokenContract:    eip20,
		PlcrAddr:         plcrAddr,
		PlcrContract:     plcr,
		ParamAddr:        paramAddr,
		ParamContract:    param,
		GovtAddr:         govtAddr,
		GovtContract:     govt,
		CivilTcrAddr:     civilTcrAddr,
		CivilTcrContract: civilTcr,
		NewsroomAddr:     newsroomAddr,
		NewsroomContract: newsroom,
		Client:           client,
		Auth:             auth,
	}, nil

}

// setupTestNewsroomContract deploys a test newsroom contract to the given ContractBackend.
func setupTestNewsroomContract(client bind.ContractBackend, auth *bind.TransactOpts) (common.Address,
	*newsroom.NewsroomContract, error) {
	address, _, contract, err := newsroom.DeployNewsroomContract(
		auth,
		client,
		"Your Newsroom Here",
		"newsroom.com/charter",
		[32]byte{},
	)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, contract, nil
}

// setupTestEIP20Contract deploys a test token contract to the given ContractBackend.
func setupTestEIP20Contract(client bind.ContractBackend, auth *bind.TransactOpts) (common.Address,
	*eip20.EIP20Contract, error) {
	address, _, contract, err := eip20.DeployEIP20Contract(
		auth,
		client,
		big.NewInt(9223372036854775807),
		"CivilTokenTest",
		18,
		"CVLT",
	)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, contract, nil
}

// setupTestPLCRVotingContract deploys a test PLCR voting contract to the given ContractBackend.
func setupTestPLCRVotingContract(client bind.ContractBackend, auth *bind.TransactOpts,
	tokenAddr common.Address) (common.Address, *plcr.PLCRVotingContract, error) {
	address, _, contract, err := plcr.DeployPLCRVotingContract(
		auth,
		client,
		tokenAddr,
	)
	if err != nil {
		return common.Address{}, nil, err
	}

	return address, contract, nil
}

// setupTestParameterizerContract deploys a test parameterizer voting contract to the given ContractBackend.
func setupTestParameterizerContract(client bind.ContractBackend, auth *bind.TransactOpts,
	tokenAddr common.Address, plcrAddr common.Address) (common.Address, *parameterizer.ParameterizerContract, error) {
	params := [16]*big.Int{
		big.NewInt(minDeposit),
		big.NewInt(pMinDeposit),
		big.NewInt(applyStageLength),
		big.NewInt(pApplyStageLength),
		big.NewInt(commitStageLength),
		big.NewInt(pCommitStageLength),
		big.NewInt(revealStageLength),
		big.NewInt(pRevealStageLength),
		big.NewInt(dispensationPct),
		big.NewInt(pDispensationPct),
		big.NewInt(voteQuorum),
		big.NewInt(pVoteQuorum),
		big.NewInt(pProcessBy),
		big.NewInt(challengeAppealLength),
		big.NewInt(appealChallengeCommitStageLength),
		big.NewInt(appealChallengeRevealStageLength),
	}
	address, _, contract, err := parameterizer.DeployParameterizerContract(
		auth,
		client,
		tokenAddr,
		plcrAddr,
		params,
	)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, contract, nil
}

// setupTestGovernmentContract deploys a test government contract to the given ContractBackend.
func setupTestGovernmentContract(client bind.ContractBackend, auth *bind.TransactOpts,
	appellateAddr common.Address, governmentControllerAddr common.Address) (common.Address,
	*government.GovernmentContract, error) {

	hash := sha3.NewKeccak256()
	var buf []byte
	var fixedBuf [32]byte
	decoded, _ := hex.DecodeString("Constitution: Be Bad.")
	hash.Write(decoded)
	buf = hash.Sum(buf)
	copy(fixedBuf[:], buf)

	address, _, contract, err := government.DeployGovernmentContract(
		auth,
		client,
		appellateAddr,
		governmentControllerAddr,
		big.NewInt(appealFeeAmount),
		big.NewInt(requestAppealPhaseLength),
		big.NewInt(judgeAppealPhaseLength),
		big.NewInt(appealSupermajorityPercentage),
		fixedBuf,
		"http://madeupURL.com",
	)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, contract, nil
}

// setupTestCivilTCRContract deploys a test Civil TCR contract to the given ContractBackend.
func setupTestCivilTCRContract(client bind.ContractBackend, auth *bind.TransactOpts,
	tokenAddr common.Address, plcrAddr common.Address, paramAddr common.Address,
	govtAddr common.Address) (common.Address, *tcr.CivilTCRContract, error) {
	address, _, contract, err := tcr.DeployCivilTCRContract(
		auth,
		client,
		tokenAddr,
		plcrAddr,
		paramAddr,
		govtAddr,
	)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, contract, nil
}
