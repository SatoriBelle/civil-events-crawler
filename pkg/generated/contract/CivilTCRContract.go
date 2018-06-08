// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// CivilTCRContractABI is the input ABI used to generate the binding from.
const CivilTCRContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"government\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"string\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"appWasMade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challengeRequestAppealExpiries\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"listings\",\"outputs\":[{\"name\":\"applicationExpiry\",\"type\":\"uint256\"},{\"name\":\"isWhitelisted\",\"type\":\"bool\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"unstakedDeposit\",\"type\":\"uint256\"},{\"name\":\"challengeID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"challengeExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"exitListing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"challengeID\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challenges\",\"outputs\":[{\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"name\":\"challenger\",\"type\":\"address\"},{\"name\":\"resolved\",\"type\":\"bool\"},{\"name\":\"stake\",\"type\":\"uint256\"},{\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"appeals\",\"outputs\":[{\"name\":\"requester\",\"type\":\"address\"},{\"name\":\"appealFeePaid\",\"type\":\"uint256\"},{\"name\":\"appealPhaseExpiry\",\"type\":\"uint256\"},{\"name\":\"appealGranted\",\"type\":\"bool\"},{\"name\":\"appealOpenToChallengeExpiry\",\"type\":\"uint256\"},{\"name\":\"appealChallengeID\",\"type\":\"uint256\"},{\"name\":\"overturned\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"challengeID\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"hasClaimedTokens\",\"outputs\":[{\"name\":\"hasClaimedTokens\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"challengeID\",\"type\":\"uint256\"}],\"name\":\"determineReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"canBeWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parameterizer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voting\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"plcrAddr\",\"type\":\"address\"},{\"name\":\"paramsAddr\",\"type\":\"address\"},{\"name\":\"govtAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challengeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"appealFeePaid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"requester\",\"type\":\"address\"}],\"name\":\"_AppealRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challengeID\",\"type\":\"uint256\"}],\"name\":\"_AppealGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challengeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"_FailedChallengeOverturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challengeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"_SuccessfulChallengeOverturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challengeID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"appealChallengeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"string\"}],\"name\":\"_GrantedAppealChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challengeID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"appealChallengeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"_GrantedAppealOverturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challengeID\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"appealChallengeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"_GrantedAppealConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newGovernment\",\"type\":\"address\"}],\"name\":\"_GovernmentTransfered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"appEndDate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"applicant\",\"type\":\"address\"}],\"name\":\"_Application\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challengeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"commitEndDate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealEndDate\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"_Challenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"added\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newTotal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"_Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"withdrew\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newTotal\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"_Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"_ApplicationWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"_ApplicationRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"_ListingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"_ListingWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"_TouchAndRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challengeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"_ChallengeFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"listingAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challengeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokens\",\"type\":\"uint256\"}],\"name\":\"_ChallengeSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"challengeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"_RewardClaimed\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"requestAppeal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"grantAppeal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"transferGovernment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"updateStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"string\"}],\"name\":\"challenge\",\"outputs\":[{\"name\":\"challengeID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"string\"}],\"name\":\"challengeGrantedAppeal\",\"outputs\":[{\"name\":\"challengeID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"},{\"name\":\"challengeID\",\"type\":\"uint256\"},{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"voterReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"challengeCanBeResolved\",\"outputs\":[{\"name\":\"canBeResolved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"appealCanBeResolved\",\"outputs\":[{\"name\":\"canBeResolved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listingAddress\",\"type\":\"address\"}],\"name\":\"appealChallengeCanBeResolved\",\"outputs\":[{\"name\":\"canBeResolved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CivilTCRContractBin is the compiled bytecode used for deploying new contracts.
const CivilTCRContractBin = `0x608060405234801561001057600080fd5b50604051608080620061658339810180604052810190808051906020019092919080519060200190929190805190602001909291908051906020019092919050505083838382828282828282600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050505050505080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050615fe9806200017c6000396000f30060806040526004361061016a576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806301162b931461016f5780630aac4543146101b2578063120c40c61461020d57806325ecef04146102aa5780632672f526146103055780632ea9b6961461035c578063432261bb146103b757806347e7ef24146103fa57806355246b9c146104475780635b5d4e73146104da57806361a9a8ca1461051d57806364c373181461057857806365d96c82146105b95780636eefcab91461065c5780637d3c5d01146106b75780637fd57e0b146106fa57806386bb8f371461073d5780638f1d377614610774578063a7aad3db14610801578063acff86871461086c578063bc4b010f1461090b578063be78533f146109a8578063c8187cf114610a0d578063dd4e7cfd14610a4e578063e1e3f91514610aa9578063f3fef3a314610b00578063fc0c546a14610b4d578063fce1ccca14610ba4575b600080fd5b34801561017b57600080fd5b506101b0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bfb565b005b3480156101be57600080fd5b506101f3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c76565b604051808215151515815260200191505060405180910390f35b34801561021957600080fd5b50610294600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610dc5565b6040518082815260200191505060405180910390f35b3480156102b657600080fd5b506102eb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611545565b604051808215151515815260200191505060405180910390f35b34801561031157600080fd5b5061031a611615565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561036857600080fd5b5061039d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061163b565b604051808215151515815260200191505060405180910390f35b3480156103c357600080fd5b506103f8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116de565b005b34801561040657600080fd5b50610445600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611cba565b005b34801561045357600080fd5b506104d8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611f22565b005b3480156104e657600080fd5b5061051b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612012565b005b34801561052957600080fd5b5061055e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506121b3565b604051808215151515815260200191505060405180910390f35b34801561058457600080fd5b506105a360048036038101908080359060200190929190505050612201565b6040518082815260200191505060405180910390f35b3480156105c557600080fd5b506105fa600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612219565b60405180868152602001851515151581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019550505050505060405180910390f35b34801561066857600080fd5b5061069d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061227c565b604051808215151515815260200191505060405180910390f35b3480156106c357600080fd5b506106f8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061233e565b005b34801561070657600080fd5b5061073b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612636565b005b34801561074957600080fd5b5061077260048036038101908080359060200190929190803590602001909291905050506127c3565b005b34801561078057600080fd5b5061079f600480360381019080803590602001909291905050506127e9565b604051808681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001841515151581526020018381526020018281526020019550505050505060405180910390f35b34801561080d57600080fd5b50610856600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919050505061284c565b6040518082815260200191505060405180910390f35b34801561087857600080fd5b5061089760048036038101908080359060200190929190505050612afe565b604051808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001878152602001868152602001851515151581526020018481526020018381526020018215151515815260200197505050505050505060405180910390f35b34801561091757600080fd5b50610992600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050612b7a565b6040518082815260200191505060405180910390f35b3480156109b457600080fd5b506109f360048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612eb5565b604051808215151515815260200191505060405180910390f35b348015610a1957600080fd5b50610a3860048036038101908080359060200190929190505050612f1f565b6040518082815260200191505060405180910390f35b348015610a5a57600080fd5b50610a8f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612fde565b604051808215151515815260200191505060405180910390f35b348015610ab557600080fd5b50610abe6130ba565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610b0c57600080fd5b50610b4b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506130e0565b005b348015610b5957600080fd5b50610b62613438565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610bb057600080fd5b50610bb961345e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610c0481612fde565b15610c1757610c1281613484565b610c73565b610c208161163b565b15610c3357610c2e81613549565b610c72565b610c3c81611545565b15610c4f57610c4a816139a0565b610c71565b610c5881610c76565b15610c6b57610c6681613bbe565b610c70565b600080fd5b5b5b5b50565b6000806000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301549150600760008381526020019081526020016000209050600081600501541415610ceb5760009250610dbe565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ee68483082600501546040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015610d8057600080fd5b505af1158015610d94573d6000803e3d6000fd5b505050506040513d6020811015610daa57600080fd5b810190808051906020019092919050505092505b5050919050565b6000806000806000600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020935060076000856003015481526020019081526020016000209250600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825260148152602001807f61707065616c566f746550657263656e74616765000000000000000000000000815250602001915050602060405180830381600087803b158015610ee957600080fd5b505af1158015610efd573d6000803e3d6000fd5b505050506040513d6020811015610f1357600080fd5b810190808051906020019092919050505091508260030160009054906101000a900460ff161515610f4357600080fd5b60008360050154141515610f5657600080fd5b428360040154111515610f6857600080fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166332ed3d6083600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825260188152602001807f6368616c6c656e676541707065616c436f6d6d69744c656e0000000000000000815250602001915050602060405180830381600087803b15801561106957600080fd5b505af115801561107d573d6000803e3d6000fd5b505050506040513d602081101561109357600080fd5b8101908080519060200190929190505050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825260188152602001807f6368616c6c656e676541707065616c52657665616c4c656e0000000000000000815250602001915050602060405180830381600087803b15801561116657600080fd5b505af115801561117a573d6000803e3d6000fd5b505050506040513d602081101561119057600080fd5b81019080805190602001909291905050506040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808481526020018381526020018281526020019350505050602060405180830381600087803b15801561120157600080fd5b505af1158015611215573d6000803e3d6000fd5b505050506040513d602081101561122b57600080fd5b8101908080519060200190929190505050905060a06040519081016040528060648560010154856064030281151561125f57fe5b0481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020018460010154815260200160008152506000808381526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff0219169083151502179055506060820151816002015560808201518160030155905050808360050181905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd333086600101546040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561143957600080fd5b505af115801561144d573d6000803e3d6000fd5b505050506040513d602081101561146357600080fd5b8101908080519060200190929190505050151561147f57600080fd5b8084600301548873ffffffffffffffffffffffffffffffffffffffff167fedfe36bf00610fb3b5474f1efd2de0d52ffb9a50b056ee37c33cea805fd44161896040518080602001828103825283818151815260200191508051906020019080838360005b838110156114fe5780820151818401526020810190506114e3565b50505050905090810190601f16801561152b5780820380516001836020036101000a031916815260200191505b509250505060405180910390a48094505050505092915050565b6000806000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015491506007600083815260200190815260200160002090506115ad8461227c565b15156115b857600080fd5b6000816002015414156115ce576000925061160e565b8060030160009054906101000a900460ff1615156115f45742816002015410925061160e565b42816004015410801561160b575060008160050154145b92505b5050919050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030154905061168c8361227c565b151561169757600080fd5b42600660008381526020019081526020016000205411156116bb57600091506116d8565b600060076000838152602001908152602001600020600201541491505b50919050565b6000806000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ee68483084600301546040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1580156117b957600080fd5b505af11580156117cd573d6000803e3d6000fd5b505050506040513d60208110156117e357600080fd5b810190808051906020019092919050505015156117ff57600080fd5b4260066000856003015481526020019081526020016000205411151561182457600080fd5b6000600760008560030154815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561188357600080fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825260098152602001807f61707065616c4665650000000000000000000000000000000000000000000000815250602001915050602060405180830381600087803b15801561194557600080fd5b505af1158015611959573d6000803e3d6000fd5b505050506040513d602081101561196f57600080fd5b8101908080519060200190929190505050915060076000846003015481526020019081526020016000209050338160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550818160010181905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252600e8152602001807f6a7564676541707065616c4c656e000000000000000000000000000000000000815250602001915050602060405180830381600087803b158015611aa957600080fd5b505af1158015611abd573d6000803e3d6000fd5b505050506040513d6020811015611ad357600080fd5b810190808051906020019092919050505042018160020181905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015611be757600080fd5b505af1158015611bfb573d6000803e3d6000fd5b505050506040513d6020811015611c1157600080fd5b81019080805190602001909291905050501515611c2d57600080fd5b82600301548473ffffffffffffffffffffffffffffffffffffffff167f7e8074c49cf8258160ac15b3e0fad39069cc31cec0b58a1b428d65e6a4e2ed468433604051808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a350505050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090503373ffffffffffffffffffffffffffffffffffffffff168160010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515611d5b57600080fd5b818160020160008282540192505081905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015611e6657600080fd5b505af1158015611e7a573d6000803e3d6000fd5b505050506040513d6020811015611e9057600080fd5b81019080805190602001909291905050501515611eac57600080fd5b3373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167ffc2e298800eb7bcacdea96213f53962a6bdf27d2a560f831d4e42301492e8f6a848460020154604051808381526020018281526020019250505060405180910390a3505050565b8260008190503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015611fa357600080fd5b505af1158015611fb7573d6000803e3d6000fd5b505050506040513d6020811015611fcd57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1614151561200057600080fd5b61200b8585856141ae565b5050505050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635793b9cf6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561209857600080fd5b505af11580156120ac573d6000803e3d6000fd5b505050506040513d60208110156120c257600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561210c57600080fd5b80600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f016b4781993f669e6eac42012fead2d96f8381769b4efbb4ad686cca6031ea8881604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b600080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154119050919050565b60066020528060005260406000206000915090505481565b60016020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16908060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154905085565b600080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015490506000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030154118015612336575060008082815260200190815260200160002060010160149054906101000a900460ff16155b915050919050565b600080600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166356e1fb886040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156123c757600080fd5b505af11580156123db573d6000803e3d6000fd5b505050506040513d60208110156123f157600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561243b57600080fd5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209150600760008360030154815260200190815260200160002090504281600201541115156124a757600080fd5b8060030160009054906101000a900460ff161515156124c557600080fd5b60018160030160006101000a81548160ff021916908315150217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825260128152602001807f6368616c6c656e676541707065616c4c656e0000000000000000000000000000815250602001915050602060405180830381600087803b1580156125a457600080fd5b505af11580156125b8573d6000803e3d6000fd5b505050506040513d60208110156125ce57600080fd5b81019080805190602001909291905050504201816004018190555081600301548373ffffffffffffffffffffffffffffffffffffffff167f4142d097c264b24b2e32e9965e184ba773286757b5af93907d7b62022e1333af60405160405180910390a3505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090503373ffffffffffffffffffffffffffffffffffffffff168160010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156126d757600080fd5b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff16151561273257600080fd5b60008160030154148061276857506000808260030154815260200190815260200160002060010160149054906101000a900460ff165b151561277357600080fd5b61277c826141d6565b8173ffffffffffffffffffffffffffffffffffffffff167f09a024f7311a15ac363521bddca1d9937c4244ab9a25e6c968e610b35ecc750360405160405180910390a25050565b600080600084815260200190815260200160002090506127e48383836144a2565b505050565b60006020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160149054906101000a900460ff16908060020154908060030154905085565b60008060008060008060008060008a81526020019081526020016000209550600760008a8152602001908152602001600020945085600301549350856000015492508460030160009054906101000a900460ff1680156128bb57508460060160009054906101000a900460ff16155b91506000905081156129d657600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636afa97a88b8b8b6040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050602060405180830381600087803b15801561299457600080fd5b505af11580156129a8573d6000803e3d6000fd5b505050506040513d60208110156129be57600080fd5b81019080805190602001909291905050509050612ae1565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b43bd0698b8b8b6040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050602060405180830381600087803b158015612aa357600080fd5b505af1158015612ab7573d6000803e3d6000fd5b505050506040513d6020811015612acd57600080fd5b810190808051906020019092919050505090505b83838202811515612aee57fe5b0496505050505050509392505050565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030160009054906101000a900460ff16908060040154908060050154908060060160009054906101000a900460ff16905087565b6000806000612b898585614723565b91506000821115612eaa57600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825260108152602001807f7265717565737441707065616c4c656e00000000000000000000000000000000815250602001915050602060405180830381600087803b158015612c5657600080fd5b505af1158015612c6a573d6000803e3d6000fd5b505050506040513d6020811015612c8057600080fd5b8101908080519060200190929190505050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252600e8152602001807f72657665616c53746167654c656e000000000000000000000000000000000000815250602001915050602060405180830381600087803b158015612d5357600080fd5b505af1158015612d67573d6000803e3d6000fd5b505050506040513d6020811015612d7d57600080fd5b8101908080519060200190929190505050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252600e8152602001807f636f6d6d697453746167654c656e000000000000000000000000000000000000815250602001915050602060405180830381600087803b158015612e5057600080fd5b505af1158015612e64573d6000803e3d6000fd5b505050506040513d6020811015612e7a57600080fd5b81019080805190602001909291905050500101905080420160066000848152602001908152602001600020819055505b819250505092915050565b600080600084815260200190815260200160002060040160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000612fd760008084815260200190815260200160002060a06040519081016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff1615151515815260200160028201548152602001600382015481525050836151ee565b9050919050565b600080600080600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209250826003015491506000808381526020019081526020016000209050613049856121b3565b80156130585750428360000154105b801561307357508260010160009054906101000a900460ff16155b801561309f5750600082148061309e5750600115158160010160149054906101000a900460ff161515145b5b156130ad57600193506130b2565b600093505b505050919050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090503373ffffffffffffffffffffffffffffffffffffffff168160010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561318157600080fd5b8060020154821115151561319457600080fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252600a8152602001807f6d696e4465706f73697400000000000000000000000000000000000000000000815250602001915050602060405180830381600087803b15801561325657600080fd5b505af115801561326a573d6000803e3d6000fd5b505050506040513d602081101561328057600080fd5b810190808051906020019092919050505082826002015403101515156132a557600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561336a57600080fd5b505af115801561337e573d6000803e3d6000fd5b505050506040513d602081101561339457600080fd5b810190808051906020019092919050505015156133b057600080fd5b8181600201600082825403925050819055503373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f7b7771adeec078e4a338f627a52f307a7fd66d915cb133b5ace441bed26abc0b848460020154604051808381526020018281526020019250505060405180910390a3505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002091508160010160009054906101000a900460ff16905060018260010160006101000a81548160ff021916908315150217905550801515613544578273ffffffffffffffffffffffffffffffffffffffff167fb268dc7f76f496fd307b40e0a3b57631a7e46123d9f708b1573bd4efbba3bdba60405160405180910390a25b505050565b600080600080600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002093508360030154925060008084815260200190815260200160002091506135b483612f1f565b905060018260010160146101000a81548160ff021916908315150217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663053e71a6846040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561366457600080fd5b505af1158015613678573d6000803e3d6000fd5b505050506040513d602081101561368e57600080fd5b81019080805190602001909291905050508260030181905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166349403183846040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561373857600080fd5b505af115801561374c573d6000803e3d6000fd5b505050506040513d602081101561376257600080fd5b81019080805190602001909291905050501561391457613781856141d6565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8360010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561386a57600080fd5b505af115801561387e573d6000803e3d6000fd5b505050506040513d602081101561389457600080fd5b810190808051906020019092919050505015156138b057600080fd5b828573ffffffffffffffffffffffffffffffffffffffff167fe86031b52c5a57c0768c3f196b63abf60b5ed012de77ce1bb88fc63588f7603a84600001548560030154604051808381526020018281526020019250505060405180910390a3613999565b61391d85613484565b80846002016000828254019250508190555060008460030181905550828573ffffffffffffffffffffffffffffffffffffffff167f3639145ca0a6a8008912a972730b5c8634e1fd1555ea44a257a8de53c30d72fb84600001548560030154604051808381526020018281526020019250505060405180910390a35b5050505050565b600080600080600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209350600760008560030154815260200190815260200160002092508260030160009054906101000a900460ff1615613b5857613a20856153cf565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685600101546040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015613b0d57600080fd5b505af1158015613b21573d6000803e3d6000fd5b505050506040513d6020811015613b3757600080fd5b81019080805190602001909291905050501515613b5357600080fd5b613bb7565b60008085600301548152602001908152602001600020915060028360010154811515613b8057fe5b049050808260000160008282540192505081905550808360010154038260020160008282540192505081905550613bb685613549565b5b5050505050565b600080600080600080600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209550600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015494506007600087600301548152602001908152602001600020935083600501549250600080856005015481526020019081526020016000209150613c8e83612f1f565b905060018260010160146101000a81548160ff021916908315150217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663053e71a6866040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015613d3e57600080fd5b505af1158015613d52573d6000803e3d6000fd5b505050506040513d6020811015613d6857600080fd5b81019080805190602001909291905050508260030181905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166349403183846040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015613e1257600080fd5b505af1158015613e26573d6000803e3d6000fd5b505050506040513d6020811015613e3c57600080fd5b81019080805190602001909291905050501561400c57613e5b87613549565b60018460060160006101000a81548160ff021916908315150217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8360010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015613f6157600080fd5b505af1158015613f75573d6000803e3d6000fd5b505050506040513d6020811015613f8b57600080fd5b81019080805190602001909291905050501515613fa757600080fd5b82858873ffffffffffffffffffffffffffffffffffffffff167fc49556ab8bcbdd0403e98b6dac260ac86008640cda2a5a229c895353b87f2feb85600001548660030154604051808381526020018281526020019250505060405180910390a46141a5565b614015876153cf565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156140fe57600080fd5b505af1158015614112573d6000803e3d6000fd5b505050506040513d602081101561412857600080fd5b8101908080519060200190929190505050151561414457600080fd5b82858873ffffffffffffffffffffffffffffffffffffffff167f8a7e8d1076fec4f93e4d57111b034ab3975009f601977350c4542e15d2e8b0f685600001548660030154604051808381526020018281526020019250505060405180910390a45b50505050505050565b826000813b90506000811115156141c457600080fd5b6141cf85858561581c565b5050505050565b600080600080600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002093508360010160009054906101000a900460ff1692508360010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16915083600201549050600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000808201600090556001820160006101000a81549060ff02191690556001820160016101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600282016000905560038201600090555050600081111561440957600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156143c257600080fd5b505af11580156143d6573d6000803e3d6000fd5b505050506040513d60208110156143ec57600080fd5b8101908080519060200190929190505050151561440857600080fd5b5b8215614457578473ffffffffffffffffffffffffffffffffffffffff167f5aebb54d5afe29103adbc464fd4e0313af619f4d19f10a0209128b76cd9d0b0760405160405180910390a261449b565b8473ffffffffffffffffffffffffffffffffffffffff167f8ad9ca8735c55207756208e8f59c7693e83d234fc6c8af2713f266468edff87160405160405180910390a25b5050505050565b600080600015158360040160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514151561450657600080fd5b600115158360010160149054906101000a900460ff16151514151561452a57600080fd5b614535338686615d23565b915061454233868661284c565b905081836003016000828254039250508190555080836000016000828254039250508190555060018360040160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561468757600080fd5b505af115801561469b573d6000803e3d6000fd5b505050506040513d60208110156146b157600080fd5b810190808051906020019092919050505015156146cd57600080fd5b3373ffffffffffffffffffffffffffffffffffffffff16857f6f4c982acc31b0af2cf1dc1556f21c0325d893782d65e83c68a5534a33f59957836040518082815260200191505060405180910390a35050505050565b600080600080600080600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209450600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252600a8152602001807f6d696e4465706f73697400000000000000000000000000000000000000000000815250602001915050602060405180830381600087803b15801561482f57600080fd5b505af1158015614843573d6000803e3d6000fd5b505050506040513d602081101561485957600080fd5b81019080805190602001909291905050509350614875886121b3565b151561488057600080fd5b6000856003015414151561489357600080fd5b83856002015410156148f4576148a8886141d6565b8773ffffffffffffffffffffffffffffffffffffffff167fc88462fa6972b64560d1dd34c4d66f2ff9841b2f974bdb0fab0827133d692f6460405160405180910390a2600095506151e3565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166332ed3d60600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252600a8152602001807f766f746551756f72756d00000000000000000000000000000000000000000000815250602001915050602060405180830381600087803b1580156149f457600080fd5b505af1158015614a08573d6000803e3d6000fd5b505050506040513d6020811015614a1e57600080fd5b8101908080519060200190929190505050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252600e8152602001807f636f6d6d697453746167654c656e000000000000000000000000000000000000815250602001915050602060405180830381600087803b158015614af157600080fd5b505af1158015614b05573d6000803e3d6000fd5b505050506040513d6020811015614b1b57600080fd5b8101908080519060200190929190505050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252600e8152602001807f72657665616c53746167654c656e000000000000000000000000000000000000815250602001915050602060405180830381600087803b158015614bee57600080fd5b505af1158015614c02573d6000803e3d6000fd5b505050506040513d6020811015614c1857600080fd5b81019080805190602001909291905050506040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808481526020018381526020018281526020019350505050602060405180830381600087803b158015614c8957600080fd5b505af1158015614c9d573d6000803e3d6000fd5b505050506040513d6020811015614cb357600080fd5b8101908080519060200190929190505050925060a060405190810160405280606486600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252600f8152602001807f64697370656e736174696f6e5063740000000000000000000000000000000000815250602001915050602060405180830381600087803b158015614d9757600080fd5b505af1158015614dab573d6000803e3d6000fd5b505050506040513d6020811015614dc157600080fd5b810190808051906020019092919050505060640302811515614ddf57fe5b0481526020013373ffffffffffffffffffffffffffffffffffffffff16815260200160001515815260200185815260200160008152506000808581526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff0219169083151502179055506060820151816002015560808201518160030155905050828560030181905550838560020160008282540392505081905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330876040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015614fc357600080fd5b505af1158015614fd7573d6000803e3d6000fd5b505050506040513d6020811015614fed57600080fd5b8101908080519060200190929190505050151561500957600080fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636148fed5846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060a060405180830381600087803b15801561509a57600080fd5b505af11580156150ae573d6000803e3d6000fd5b505050506040513d60a08110156150c457600080fd5b810190808051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050505050915091503373ffffffffffffffffffffffffffffffffffffffff16838973ffffffffffffffffffffffffffffffffffffffff167f9a8e3864cbacafc5547b8567796b4d12d51217a879192b61932f5151ce5813108a86866040518080602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156151a3578082015181840152602081019050615188565b50505050905090810190601f1680156151d05780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a48295505b505050505092915050565b600082604001511580156152ca5750600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ee684830836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561528e57600080fd5b505af11580156152a2573d6000803e3d6000fd5b505050506040513d60208110156152b857600080fd5b81019080805190602001909291905050505b15156152d557600080fd5b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663053e71a6846040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561536857600080fd5b505af115801561537c573d6000803e3d6000fd5b505050506040513d602081101561539257600080fd5b810190808051906020019092919050505014156153b857826060015160020290506153c9565b826000015183606001516002020390505b92915050565b600080600080600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020935083600301549250600080848152602001908152602001600020915061543a83612f1f565b905060018260010160146101000a81548160ff021916908315150217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8cfa3f0846040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1580156154ea57600080fd5b505af11580156154fe573d6000803e3d6000fd5b505050506040513d602081101561551457600080fd5b81019080805190602001909291905050508260030181905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166349403183846040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1580156155be57600080fd5b505af11580156155d2573d6000803e3d6000fd5b505050506040513d60208110156155e857600080fd5b81019080805190602001909291905050501561567d5761560785613484565b808460020160008282540192505081905550828573ffffffffffffffffffffffffffffffffffffffff167f72506b3ce4d8f0cf8cf6ccb7cd5281af2b0d020121fb20abfa073eeb3f6d296e84600001548560030154604051808381526020018281526020019250505060405180910390a3615815565b615686856141d6565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8360010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561576f57600080fd5b505af1158015615783573d6000803e3d6000fd5b505050506040513d602081101561579957600080fd5b810190808051906020019092919050505015156157b557600080fd5b828573ffffffffffffffffffffffffffffffffffffffff167f446922bbfdaa528d4a969857cd0894d6bf8bbff52226624e752b3f1be7513b0a84600001548560030154604051808381526020018281526020019250505060405180910390a35b5050505050565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060010160009054906101000a900460ff1615151561587d57600080fd5b615886846121b3565b15151561589257600080fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252600a8152602001807f6d696e4465706f73697400000000000000000000000000000000000000000000815250602001915050602060405180830381600087803b15801561595457600080fd5b505af1158015615968573d6000803e3d6000fd5b505050506040513d602081101561597e57600080fd5b8101908080519060200190929190505050831015151561599d57600080fd5b338160010160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550615aef600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252600d8152602001807f6170706c7953746167654c656e00000000000000000000000000000000000000815250602001915050602060405180830381600087803b158015615aa557600080fd5b505af1158015615ab9573d6000803e3d6000fd5b505050506040513d6020811015615acf57600080fd5b810190808051906020019092919050505042615fa190919063ffffffff16565b8160000181905550828160020181905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015615bf957600080fd5b505af1158015615c0d573d6000803e3d6000fd5b505050506040513d6020811015615c2357600080fd5b81019080805190602001909291905050501515615c3f57600080fd5b3373ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f09cd8dcaf170a50a26316b5fe0727dd9fb9581a688d65e758b16a1650da65c0b858460000154866040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015615ce1578082015181840152602081019050615cc6565b50505050905090810190601f168015615d0e5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a350505050565b6000806007600085815260200190815260200160002060030160009054906101000a900460ff168015615d7757506007600085815260200190815260200160002060060160009054906101000a900460ff16155b90508015615e8e57600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636afa97a88686866040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050602060405180830381600087803b158015615e4c57600080fd5b505af1158015615e60573d6000803e3d6000fd5b505050506040513d6020811015615e7657600080fd5b81019080805190602001909291905050509150615f99565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b43bd0698686866040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018281526020019350505050602060405180830381600087803b158015615f5b57600080fd5b505af1158015615f6f573d6000803e3d6000fd5b505050506040513d6020811015615f8557600080fd5b810190808051906020019092919050505091505b509392505050565b60008183019050828110151515615fb457fe5b809050929150505600a165627a7a72305820d3672b7262f66b2d26cdf4018be15170706731d7d730dea358dabf162db4d6b40029`

// DeployCivilTCRContract deploys a new Ethereum contract, binding an instance of CivilTCRContract to it.
func DeployCivilTCRContract(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddr common.Address, plcrAddr common.Address, paramsAddr common.Address, govtAddr common.Address) (common.Address, *types.Transaction, *CivilTCRContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CivilTCRContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CivilTCRContractBin), backend, tokenAddr, plcrAddr, paramsAddr, govtAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CivilTCRContract{CivilTCRContractCaller: CivilTCRContractCaller{contract: contract}, CivilTCRContractTransactor: CivilTCRContractTransactor{contract: contract}, CivilTCRContractFilterer: CivilTCRContractFilterer{contract: contract}}, nil
}

// CivilTCRContract is an auto generated Go binding around an Ethereum contract.
type CivilTCRContract struct {
	CivilTCRContractCaller     // Read-only binding to the contract
	CivilTCRContractTransactor // Write-only binding to the contract
	CivilTCRContractFilterer   // Log filterer for contract events
}

// CivilTCRContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CivilTCRContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CivilTCRContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CivilTCRContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CivilTCRContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CivilTCRContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CivilTCRContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CivilTCRContractSession struct {
	Contract     *CivilTCRContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CivilTCRContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CivilTCRContractCallerSession struct {
	Contract *CivilTCRContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CivilTCRContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CivilTCRContractTransactorSession struct {
	Contract     *CivilTCRContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CivilTCRContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CivilTCRContractRaw struct {
	Contract *CivilTCRContract // Generic contract binding to access the raw methods on
}

// CivilTCRContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CivilTCRContractCallerRaw struct {
	Contract *CivilTCRContractCaller // Generic read-only contract binding to access the raw methods on
}

// CivilTCRContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CivilTCRContractTransactorRaw struct {
	Contract *CivilTCRContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCivilTCRContract creates a new instance of CivilTCRContract, bound to a specific deployed contract.
func NewCivilTCRContract(address common.Address, backend bind.ContractBackend) (*CivilTCRContract, error) {
	contract, err := bindCivilTCRContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContract{CivilTCRContractCaller: CivilTCRContractCaller{contract: contract}, CivilTCRContractTransactor: CivilTCRContractTransactor{contract: contract}, CivilTCRContractFilterer: CivilTCRContractFilterer{contract: contract}}, nil
}

// NewCivilTCRContractCaller creates a new read-only instance of CivilTCRContract, bound to a specific deployed contract.
func NewCivilTCRContractCaller(address common.Address, caller bind.ContractCaller) (*CivilTCRContractCaller, error) {
	contract, err := bindCivilTCRContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractCaller{contract: contract}, nil
}

// NewCivilTCRContractTransactor creates a new write-only instance of CivilTCRContract, bound to a specific deployed contract.
func NewCivilTCRContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CivilTCRContractTransactor, error) {
	contract, err := bindCivilTCRContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractTransactor{contract: contract}, nil
}

// NewCivilTCRContractFilterer creates a new log filterer instance of CivilTCRContract, bound to a specific deployed contract.
func NewCivilTCRContractFilterer(address common.Address, filterer bind.ContractFilterer) (*CivilTCRContractFilterer, error) {
	contract, err := bindCivilTCRContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractFilterer{contract: contract}, nil
}

// bindCivilTCRContract binds a generic wrapper to an already deployed contract.
func bindCivilTCRContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CivilTCRContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CivilTCRContract *CivilTCRContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CivilTCRContract.Contract.CivilTCRContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CivilTCRContract *CivilTCRContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.CivilTCRContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CivilTCRContract *CivilTCRContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.CivilTCRContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CivilTCRContract *CivilTCRContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CivilTCRContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CivilTCRContract *CivilTCRContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CivilTCRContract *CivilTCRContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.contract.Transact(opts, method, params...)
}

// AppWasMade is a free data retrieval call binding the contract method 0x61a9a8ca.
//
// Solidity: function appWasMade(listingAddress address) constant returns(bool)
func (_CivilTCRContract *CivilTCRContractCaller) AppWasMade(opts *bind.CallOpts, listingAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "appWasMade", listingAddress)
	return *ret0, err
}

// AppWasMade is a free data retrieval call binding the contract method 0x61a9a8ca.
//
// Solidity: function appWasMade(listingAddress address) constant returns(bool)
func (_CivilTCRContract *CivilTCRContractSession) AppWasMade(listingAddress common.Address) (bool, error) {
	return _CivilTCRContract.Contract.AppWasMade(&_CivilTCRContract.CallOpts, listingAddress)
}

// AppWasMade is a free data retrieval call binding the contract method 0x61a9a8ca.
//
// Solidity: function appWasMade(listingAddress address) constant returns(bool)
func (_CivilTCRContract *CivilTCRContractCallerSession) AppWasMade(listingAddress common.Address) (bool, error) {
	return _CivilTCRContract.Contract.AppWasMade(&_CivilTCRContract.CallOpts, listingAddress)
}

// AppealCanBeResolved is a free data retrieval call binding the contract method 0x25ecef04.
//
// Solidity: function appealCanBeResolved(listingAddress address) constant returns(canBeResolved bool)
func (_CivilTCRContract *CivilTCRContractCaller) AppealCanBeResolved(opts *bind.CallOpts, listingAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "appealCanBeResolved", listingAddress)
	return *ret0, err
}

// AppealCanBeResolved is a free data retrieval call binding the contract method 0x25ecef04.
//
// Solidity: function appealCanBeResolved(listingAddress address) constant returns(canBeResolved bool)
func (_CivilTCRContract *CivilTCRContractSession) AppealCanBeResolved(listingAddress common.Address) (bool, error) {
	return _CivilTCRContract.Contract.AppealCanBeResolved(&_CivilTCRContract.CallOpts, listingAddress)
}

// AppealCanBeResolved is a free data retrieval call binding the contract method 0x25ecef04.
//
// Solidity: function appealCanBeResolved(listingAddress address) constant returns(canBeResolved bool)
func (_CivilTCRContract *CivilTCRContractCallerSession) AppealCanBeResolved(listingAddress common.Address) (bool, error) {
	return _CivilTCRContract.Contract.AppealCanBeResolved(&_CivilTCRContract.CallOpts, listingAddress)
}

// AppealChallengeCanBeResolved is a free data retrieval call binding the contract method 0x0aac4543.
//
// Solidity: function appealChallengeCanBeResolved(listingAddress address) constant returns(canBeResolved bool)
func (_CivilTCRContract *CivilTCRContractCaller) AppealChallengeCanBeResolved(opts *bind.CallOpts, listingAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "appealChallengeCanBeResolved", listingAddress)
	return *ret0, err
}

// AppealChallengeCanBeResolved is a free data retrieval call binding the contract method 0x0aac4543.
//
// Solidity: function appealChallengeCanBeResolved(listingAddress address) constant returns(canBeResolved bool)
func (_CivilTCRContract *CivilTCRContractSession) AppealChallengeCanBeResolved(listingAddress common.Address) (bool, error) {
	return _CivilTCRContract.Contract.AppealChallengeCanBeResolved(&_CivilTCRContract.CallOpts, listingAddress)
}

// AppealChallengeCanBeResolved is a free data retrieval call binding the contract method 0x0aac4543.
//
// Solidity: function appealChallengeCanBeResolved(listingAddress address) constant returns(canBeResolved bool)
func (_CivilTCRContract *CivilTCRContractCallerSession) AppealChallengeCanBeResolved(listingAddress common.Address) (bool, error) {
	return _CivilTCRContract.Contract.AppealChallengeCanBeResolved(&_CivilTCRContract.CallOpts, listingAddress)
}

// Appeals is a free data retrieval call binding the contract method 0xacff8687.
//
// Solidity: function appeals( uint256) constant returns(requester address, appealFeePaid uint256, appealPhaseExpiry uint256, appealGranted bool, appealOpenToChallengeExpiry uint256, appealChallengeID uint256, overturned bool)
func (_CivilTCRContract *CivilTCRContractCaller) Appeals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Requester                   common.Address
	AppealFeePaid               *big.Int
	AppealPhaseExpiry           *big.Int
	AppealGranted               bool
	AppealOpenToChallengeExpiry *big.Int
	AppealChallengeID           *big.Int
	Overturned                  bool
}, error) {
	ret := new(struct {
		Requester                   common.Address
		AppealFeePaid               *big.Int
		AppealPhaseExpiry           *big.Int
		AppealGranted               bool
		AppealOpenToChallengeExpiry *big.Int
		AppealChallengeID           *big.Int
		Overturned                  bool
	})
	out := ret
	err := _CivilTCRContract.contract.Call(opts, out, "appeals", arg0)
	return *ret, err
}

// Appeals is a free data retrieval call binding the contract method 0xacff8687.
//
// Solidity: function appeals( uint256) constant returns(requester address, appealFeePaid uint256, appealPhaseExpiry uint256, appealGranted bool, appealOpenToChallengeExpiry uint256, appealChallengeID uint256, overturned bool)
func (_CivilTCRContract *CivilTCRContractSession) Appeals(arg0 *big.Int) (struct {
	Requester                   common.Address
	AppealFeePaid               *big.Int
	AppealPhaseExpiry           *big.Int
	AppealGranted               bool
	AppealOpenToChallengeExpiry *big.Int
	AppealChallengeID           *big.Int
	Overturned                  bool
}, error) {
	return _CivilTCRContract.Contract.Appeals(&_CivilTCRContract.CallOpts, arg0)
}

// Appeals is a free data retrieval call binding the contract method 0xacff8687.
//
// Solidity: function appeals( uint256) constant returns(requester address, appealFeePaid uint256, appealPhaseExpiry uint256, appealGranted bool, appealOpenToChallengeExpiry uint256, appealChallengeID uint256, overturned bool)
func (_CivilTCRContract *CivilTCRContractCallerSession) Appeals(arg0 *big.Int) (struct {
	Requester                   common.Address
	AppealFeePaid               *big.Int
	AppealPhaseExpiry           *big.Int
	AppealGranted               bool
	AppealOpenToChallengeExpiry *big.Int
	AppealChallengeID           *big.Int
	Overturned                  bool
}, error) {
	return _CivilTCRContract.Contract.Appeals(&_CivilTCRContract.CallOpts, arg0)
}

// CanBeWhitelisted is a free data retrieval call binding the contract method 0xdd4e7cfd.
//
// Solidity: function canBeWhitelisted(listingAddress address) constant returns(bool)
func (_CivilTCRContract *CivilTCRContractCaller) CanBeWhitelisted(opts *bind.CallOpts, listingAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "canBeWhitelisted", listingAddress)
	return *ret0, err
}

// CanBeWhitelisted is a free data retrieval call binding the contract method 0xdd4e7cfd.
//
// Solidity: function canBeWhitelisted(listingAddress address) constant returns(bool)
func (_CivilTCRContract *CivilTCRContractSession) CanBeWhitelisted(listingAddress common.Address) (bool, error) {
	return _CivilTCRContract.Contract.CanBeWhitelisted(&_CivilTCRContract.CallOpts, listingAddress)
}

// CanBeWhitelisted is a free data retrieval call binding the contract method 0xdd4e7cfd.
//
// Solidity: function canBeWhitelisted(listingAddress address) constant returns(bool)
func (_CivilTCRContract *CivilTCRContractCallerSession) CanBeWhitelisted(listingAddress common.Address) (bool, error) {
	return _CivilTCRContract.Contract.CanBeWhitelisted(&_CivilTCRContract.CallOpts, listingAddress)
}

// ChallengeCanBeResolved is a free data retrieval call binding the contract method 0x2ea9b696.
//
// Solidity: function challengeCanBeResolved(listingAddress address) constant returns(canBeResolved bool)
func (_CivilTCRContract *CivilTCRContractCaller) ChallengeCanBeResolved(opts *bind.CallOpts, listingAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "challengeCanBeResolved", listingAddress)
	return *ret0, err
}

// ChallengeCanBeResolved is a free data retrieval call binding the contract method 0x2ea9b696.
//
// Solidity: function challengeCanBeResolved(listingAddress address) constant returns(canBeResolved bool)
func (_CivilTCRContract *CivilTCRContractSession) ChallengeCanBeResolved(listingAddress common.Address) (bool, error) {
	return _CivilTCRContract.Contract.ChallengeCanBeResolved(&_CivilTCRContract.CallOpts, listingAddress)
}

// ChallengeCanBeResolved is a free data retrieval call binding the contract method 0x2ea9b696.
//
// Solidity: function challengeCanBeResolved(listingAddress address) constant returns(canBeResolved bool)
func (_CivilTCRContract *CivilTCRContractCallerSession) ChallengeCanBeResolved(listingAddress common.Address) (bool, error) {
	return _CivilTCRContract.Contract.ChallengeCanBeResolved(&_CivilTCRContract.CallOpts, listingAddress)
}

// ChallengeExists is a free data retrieval call binding the contract method 0x6eefcab9.
//
// Solidity: function challengeExists(listingAddress address) constant returns(bool)
func (_CivilTCRContract *CivilTCRContractCaller) ChallengeExists(opts *bind.CallOpts, listingAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "challengeExists", listingAddress)
	return *ret0, err
}

// ChallengeExists is a free data retrieval call binding the contract method 0x6eefcab9.
//
// Solidity: function challengeExists(listingAddress address) constant returns(bool)
func (_CivilTCRContract *CivilTCRContractSession) ChallengeExists(listingAddress common.Address) (bool, error) {
	return _CivilTCRContract.Contract.ChallengeExists(&_CivilTCRContract.CallOpts, listingAddress)
}

// ChallengeExists is a free data retrieval call binding the contract method 0x6eefcab9.
//
// Solidity: function challengeExists(listingAddress address) constant returns(bool)
func (_CivilTCRContract *CivilTCRContractCallerSession) ChallengeExists(listingAddress common.Address) (bool, error) {
	return _CivilTCRContract.Contract.ChallengeExists(&_CivilTCRContract.CallOpts, listingAddress)
}

// ChallengeRequestAppealExpiries is a free data retrieval call binding the contract method 0x64c37318.
//
// Solidity: function challengeRequestAppealExpiries( uint256) constant returns(uint256)
func (_CivilTCRContract *CivilTCRContractCaller) ChallengeRequestAppealExpiries(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "challengeRequestAppealExpiries", arg0)
	return *ret0, err
}

// ChallengeRequestAppealExpiries is a free data retrieval call binding the contract method 0x64c37318.
//
// Solidity: function challengeRequestAppealExpiries( uint256) constant returns(uint256)
func (_CivilTCRContract *CivilTCRContractSession) ChallengeRequestAppealExpiries(arg0 *big.Int) (*big.Int, error) {
	return _CivilTCRContract.Contract.ChallengeRequestAppealExpiries(&_CivilTCRContract.CallOpts, arg0)
}

// ChallengeRequestAppealExpiries is a free data retrieval call binding the contract method 0x64c37318.
//
// Solidity: function challengeRequestAppealExpiries( uint256) constant returns(uint256)
func (_CivilTCRContract *CivilTCRContractCallerSession) ChallengeRequestAppealExpiries(arg0 *big.Int) (*big.Int, error) {
	return _CivilTCRContract.Contract.ChallengeRequestAppealExpiries(&_CivilTCRContract.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges( uint256) constant returns(rewardPool uint256, challenger address, resolved bool, stake uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractCaller) Challenges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RewardPool  *big.Int
	Challenger  common.Address
	Resolved    bool
	Stake       *big.Int
	TotalTokens *big.Int
}, error) {
	ret := new(struct {
		RewardPool  *big.Int
		Challenger  common.Address
		Resolved    bool
		Stake       *big.Int
		TotalTokens *big.Int
	})
	out := ret
	err := _CivilTCRContract.contract.Call(opts, out, "challenges", arg0)
	return *ret, err
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges( uint256) constant returns(rewardPool uint256, challenger address, resolved bool, stake uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractSession) Challenges(arg0 *big.Int) (struct {
	RewardPool  *big.Int
	Challenger  common.Address
	Resolved    bool
	Stake       *big.Int
	TotalTokens *big.Int
}, error) {
	return _CivilTCRContract.Contract.Challenges(&_CivilTCRContract.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges( uint256) constant returns(rewardPool uint256, challenger address, resolved bool, stake uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractCallerSession) Challenges(arg0 *big.Int) (struct {
	RewardPool  *big.Int
	Challenger  common.Address
	Resolved    bool
	Stake       *big.Int
	TotalTokens *big.Int
}, error) {
	return _CivilTCRContract.Contract.Challenges(&_CivilTCRContract.CallOpts, arg0)
}

// DetermineReward is a free data retrieval call binding the contract method 0xc8187cf1.
//
// Solidity: function determineReward(challengeID uint256) constant returns(uint256)
func (_CivilTCRContract *CivilTCRContractCaller) DetermineReward(opts *bind.CallOpts, challengeID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "determineReward", challengeID)
	return *ret0, err
}

// DetermineReward is a free data retrieval call binding the contract method 0xc8187cf1.
//
// Solidity: function determineReward(challengeID uint256) constant returns(uint256)
func (_CivilTCRContract *CivilTCRContractSession) DetermineReward(challengeID *big.Int) (*big.Int, error) {
	return _CivilTCRContract.Contract.DetermineReward(&_CivilTCRContract.CallOpts, challengeID)
}

// DetermineReward is a free data retrieval call binding the contract method 0xc8187cf1.
//
// Solidity: function determineReward(challengeID uint256) constant returns(uint256)
func (_CivilTCRContract *CivilTCRContractCallerSession) DetermineReward(challengeID *big.Int) (*big.Int, error) {
	return _CivilTCRContract.Contract.DetermineReward(&_CivilTCRContract.CallOpts, challengeID)
}

// Government is a free data retrieval call binding the contract method 0x2672f526.
//
// Solidity: function government() constant returns(address)
func (_CivilTCRContract *CivilTCRContractCaller) Government(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "government")
	return *ret0, err
}

// Government is a free data retrieval call binding the contract method 0x2672f526.
//
// Solidity: function government() constant returns(address)
func (_CivilTCRContract *CivilTCRContractSession) Government() (common.Address, error) {
	return _CivilTCRContract.Contract.Government(&_CivilTCRContract.CallOpts)
}

// Government is a free data retrieval call binding the contract method 0x2672f526.
//
// Solidity: function government() constant returns(address)
func (_CivilTCRContract *CivilTCRContractCallerSession) Government() (common.Address, error) {
	return _CivilTCRContract.Contract.Government(&_CivilTCRContract.CallOpts)
}

// HasClaimedTokens is a free data retrieval call binding the contract method 0xbe78533f.
//
// Solidity: function hasClaimedTokens(challengeID uint256, voter address) constant returns(hasClaimedTokens bool)
func (_CivilTCRContract *CivilTCRContractCaller) HasClaimedTokens(opts *bind.CallOpts, challengeID *big.Int, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "hasClaimedTokens", challengeID, voter)
	return *ret0, err
}

// HasClaimedTokens is a free data retrieval call binding the contract method 0xbe78533f.
//
// Solidity: function hasClaimedTokens(challengeID uint256, voter address) constant returns(hasClaimedTokens bool)
func (_CivilTCRContract *CivilTCRContractSession) HasClaimedTokens(challengeID *big.Int, voter common.Address) (bool, error) {
	return _CivilTCRContract.Contract.HasClaimedTokens(&_CivilTCRContract.CallOpts, challengeID, voter)
}

// HasClaimedTokens is a free data retrieval call binding the contract method 0xbe78533f.
//
// Solidity: function hasClaimedTokens(challengeID uint256, voter address) constant returns(hasClaimedTokens bool)
func (_CivilTCRContract *CivilTCRContractCallerSession) HasClaimedTokens(challengeID *big.Int, voter common.Address) (bool, error) {
	return _CivilTCRContract.Contract.HasClaimedTokens(&_CivilTCRContract.CallOpts, challengeID, voter)
}

// Listings is a free data retrieval call binding the contract method 0x65d96c82.
//
// Solidity: function listings( address) constant returns(applicationExpiry uint256, isWhitelisted bool, owner address, unstakedDeposit uint256, challengeID uint256)
func (_CivilTCRContract *CivilTCRContractCaller) Listings(opts *bind.CallOpts, arg0 common.Address) (struct {
	ApplicationExpiry *big.Int
	IsWhitelisted     bool
	Owner             common.Address
	UnstakedDeposit   *big.Int
	ChallengeID       *big.Int
}, error) {
	ret := new(struct {
		ApplicationExpiry *big.Int
		IsWhitelisted     bool
		Owner             common.Address
		UnstakedDeposit   *big.Int
		ChallengeID       *big.Int
	})
	out := ret
	err := _CivilTCRContract.contract.Call(opts, out, "listings", arg0)
	return *ret, err
}

// Listings is a free data retrieval call binding the contract method 0x65d96c82.
//
// Solidity: function listings( address) constant returns(applicationExpiry uint256, isWhitelisted bool, owner address, unstakedDeposit uint256, challengeID uint256)
func (_CivilTCRContract *CivilTCRContractSession) Listings(arg0 common.Address) (struct {
	ApplicationExpiry *big.Int
	IsWhitelisted     bool
	Owner             common.Address
	UnstakedDeposit   *big.Int
	ChallengeID       *big.Int
}, error) {
	return _CivilTCRContract.Contract.Listings(&_CivilTCRContract.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0x65d96c82.
//
// Solidity: function listings( address) constant returns(applicationExpiry uint256, isWhitelisted bool, owner address, unstakedDeposit uint256, challengeID uint256)
func (_CivilTCRContract *CivilTCRContractCallerSession) Listings(arg0 common.Address) (struct {
	ApplicationExpiry *big.Int
	IsWhitelisted     bool
	Owner             common.Address
	UnstakedDeposit   *big.Int
	ChallengeID       *big.Int
}, error) {
	return _CivilTCRContract.Contract.Listings(&_CivilTCRContract.CallOpts, arg0)
}

// Parameterizer is a free data retrieval call binding the contract method 0xe1e3f915.
//
// Solidity: function parameterizer() constant returns(address)
func (_CivilTCRContract *CivilTCRContractCaller) Parameterizer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "parameterizer")
	return *ret0, err
}

// Parameterizer is a free data retrieval call binding the contract method 0xe1e3f915.
//
// Solidity: function parameterizer() constant returns(address)
func (_CivilTCRContract *CivilTCRContractSession) Parameterizer() (common.Address, error) {
	return _CivilTCRContract.Contract.Parameterizer(&_CivilTCRContract.CallOpts)
}

// Parameterizer is a free data retrieval call binding the contract method 0xe1e3f915.
//
// Solidity: function parameterizer() constant returns(address)
func (_CivilTCRContract *CivilTCRContractCallerSession) Parameterizer() (common.Address, error) {
	return _CivilTCRContract.Contract.Parameterizer(&_CivilTCRContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CivilTCRContract *CivilTCRContractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CivilTCRContract *CivilTCRContractSession) Token() (common.Address, error) {
	return _CivilTCRContract.Contract.Token(&_CivilTCRContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_CivilTCRContract *CivilTCRContractCallerSession) Token() (common.Address, error) {
	return _CivilTCRContract.Contract.Token(&_CivilTCRContract.CallOpts)
}

// VoterReward is a free data retrieval call binding the contract method 0xa7aad3db.
//
// Solidity: function voterReward(voter address, challengeID uint256, salt uint256) constant returns(uint256)
func (_CivilTCRContract *CivilTCRContractCaller) VoterReward(opts *bind.CallOpts, voter common.Address, challengeID *big.Int, salt *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "voterReward", voter, challengeID, salt)
	return *ret0, err
}

// VoterReward is a free data retrieval call binding the contract method 0xa7aad3db.
//
// Solidity: function voterReward(voter address, challengeID uint256, salt uint256) constant returns(uint256)
func (_CivilTCRContract *CivilTCRContractSession) VoterReward(voter common.Address, challengeID *big.Int, salt *big.Int) (*big.Int, error) {
	return _CivilTCRContract.Contract.VoterReward(&_CivilTCRContract.CallOpts, voter, challengeID, salt)
}

// VoterReward is a free data retrieval call binding the contract method 0xa7aad3db.
//
// Solidity: function voterReward(voter address, challengeID uint256, salt uint256) constant returns(uint256)
func (_CivilTCRContract *CivilTCRContractCallerSession) VoterReward(voter common.Address, challengeID *big.Int, salt *big.Int) (*big.Int, error) {
	return _CivilTCRContract.Contract.VoterReward(&_CivilTCRContract.CallOpts, voter, challengeID, salt)
}

// Voting is a free data retrieval call binding the contract method 0xfce1ccca.
//
// Solidity: function voting() constant returns(address)
func (_CivilTCRContract *CivilTCRContractCaller) Voting(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CivilTCRContract.contract.Call(opts, out, "voting")
	return *ret0, err
}

// Voting is a free data retrieval call binding the contract method 0xfce1ccca.
//
// Solidity: function voting() constant returns(address)
func (_CivilTCRContract *CivilTCRContractSession) Voting() (common.Address, error) {
	return _CivilTCRContract.Contract.Voting(&_CivilTCRContract.CallOpts)
}

// Voting is a free data retrieval call binding the contract method 0xfce1ccca.
//
// Solidity: function voting() constant returns(address)
func (_CivilTCRContract *CivilTCRContractCallerSession) Voting() (common.Address, error) {
	return _CivilTCRContract.Contract.Voting(&_CivilTCRContract.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0x55246b9c.
//
// Solidity: function apply(listingAddress address, amount uint256, data string) returns()
func (_CivilTCRContract *CivilTCRContractTransactor) Apply(opts *bind.TransactOpts, listingAddress common.Address, amount *big.Int, data string) (*types.Transaction, error) {
	return _CivilTCRContract.contract.Transact(opts, "apply", listingAddress, amount, data)
}

// Apply is a paid mutator transaction binding the contract method 0x55246b9c.
//
// Solidity: function apply(listingAddress address, amount uint256, data string) returns()
func (_CivilTCRContract *CivilTCRContractSession) Apply(listingAddress common.Address, amount *big.Int, data string) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.Apply(&_CivilTCRContract.TransactOpts, listingAddress, amount, data)
}

// Apply is a paid mutator transaction binding the contract method 0x55246b9c.
//
// Solidity: function apply(listingAddress address, amount uint256, data string) returns()
func (_CivilTCRContract *CivilTCRContractTransactorSession) Apply(listingAddress common.Address, amount *big.Int, data string) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.Apply(&_CivilTCRContract.TransactOpts, listingAddress, amount, data)
}

// Challenge is a paid mutator transaction binding the contract method 0xbc4b010f.
//
// Solidity: function challenge(listingAddress address, data string) returns(challengeID uint256)
func (_CivilTCRContract *CivilTCRContractTransactor) Challenge(opts *bind.TransactOpts, listingAddress common.Address, data string) (*types.Transaction, error) {
	return _CivilTCRContract.contract.Transact(opts, "challenge", listingAddress, data)
}

// Challenge is a paid mutator transaction binding the contract method 0xbc4b010f.
//
// Solidity: function challenge(listingAddress address, data string) returns(challengeID uint256)
func (_CivilTCRContract *CivilTCRContractSession) Challenge(listingAddress common.Address, data string) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.Challenge(&_CivilTCRContract.TransactOpts, listingAddress, data)
}

// Challenge is a paid mutator transaction binding the contract method 0xbc4b010f.
//
// Solidity: function challenge(listingAddress address, data string) returns(challengeID uint256)
func (_CivilTCRContract *CivilTCRContractTransactorSession) Challenge(listingAddress common.Address, data string) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.Challenge(&_CivilTCRContract.TransactOpts, listingAddress, data)
}

// ChallengeGrantedAppeal is a paid mutator transaction binding the contract method 0x120c40c6.
//
// Solidity: function challengeGrantedAppeal(listingAddress address, data string) returns(challengeID uint256)
func (_CivilTCRContract *CivilTCRContractTransactor) ChallengeGrantedAppeal(opts *bind.TransactOpts, listingAddress common.Address, data string) (*types.Transaction, error) {
	return _CivilTCRContract.contract.Transact(opts, "challengeGrantedAppeal", listingAddress, data)
}

// ChallengeGrantedAppeal is a paid mutator transaction binding the contract method 0x120c40c6.
//
// Solidity: function challengeGrantedAppeal(listingAddress address, data string) returns(challengeID uint256)
func (_CivilTCRContract *CivilTCRContractSession) ChallengeGrantedAppeal(listingAddress common.Address, data string) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.ChallengeGrantedAppeal(&_CivilTCRContract.TransactOpts, listingAddress, data)
}

// ChallengeGrantedAppeal is a paid mutator transaction binding the contract method 0x120c40c6.
//
// Solidity: function challengeGrantedAppeal(listingAddress address, data string) returns(challengeID uint256)
func (_CivilTCRContract *CivilTCRContractTransactorSession) ChallengeGrantedAppeal(listingAddress common.Address, data string) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.ChallengeGrantedAppeal(&_CivilTCRContract.TransactOpts, listingAddress, data)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x86bb8f37.
//
// Solidity: function claimReward(challengeID uint256, salt uint256) returns()
func (_CivilTCRContract *CivilTCRContractTransactor) ClaimReward(opts *bind.TransactOpts, challengeID *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _CivilTCRContract.contract.Transact(opts, "claimReward", challengeID, salt)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x86bb8f37.
//
// Solidity: function claimReward(challengeID uint256, salt uint256) returns()
func (_CivilTCRContract *CivilTCRContractSession) ClaimReward(challengeID *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.ClaimReward(&_CivilTCRContract.TransactOpts, challengeID, salt)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x86bb8f37.
//
// Solidity: function claimReward(challengeID uint256, salt uint256) returns()
func (_CivilTCRContract *CivilTCRContractTransactorSession) ClaimReward(challengeID *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.ClaimReward(&_CivilTCRContract.TransactOpts, challengeID, salt)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(listingAddress address, amount uint256) returns()
func (_CivilTCRContract *CivilTCRContractTransactor) Deposit(opts *bind.TransactOpts, listingAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CivilTCRContract.contract.Transact(opts, "deposit", listingAddress, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(listingAddress address, amount uint256) returns()
func (_CivilTCRContract *CivilTCRContractSession) Deposit(listingAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.Deposit(&_CivilTCRContract.TransactOpts, listingAddress, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(listingAddress address, amount uint256) returns()
func (_CivilTCRContract *CivilTCRContractTransactorSession) Deposit(listingAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.Deposit(&_CivilTCRContract.TransactOpts, listingAddress, amount)
}

// ExitListing is a paid mutator transaction binding the contract method 0x7fd57e0b.
//
// Solidity: function exitListing(listingAddress address) returns()
func (_CivilTCRContract *CivilTCRContractTransactor) ExitListing(opts *bind.TransactOpts, listingAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.contract.Transact(opts, "exitListing", listingAddress)
}

// ExitListing is a paid mutator transaction binding the contract method 0x7fd57e0b.
//
// Solidity: function exitListing(listingAddress address) returns()
func (_CivilTCRContract *CivilTCRContractSession) ExitListing(listingAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.ExitListing(&_CivilTCRContract.TransactOpts, listingAddress)
}

// ExitListing is a paid mutator transaction binding the contract method 0x7fd57e0b.
//
// Solidity: function exitListing(listingAddress address) returns()
func (_CivilTCRContract *CivilTCRContractTransactorSession) ExitListing(listingAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.ExitListing(&_CivilTCRContract.TransactOpts, listingAddress)
}

// GrantAppeal is a paid mutator transaction binding the contract method 0x7d3c5d01.
//
// Solidity: function grantAppeal(listingAddress address) returns()
func (_CivilTCRContract *CivilTCRContractTransactor) GrantAppeal(opts *bind.TransactOpts, listingAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.contract.Transact(opts, "grantAppeal", listingAddress)
}

// GrantAppeal is a paid mutator transaction binding the contract method 0x7d3c5d01.
//
// Solidity: function grantAppeal(listingAddress address) returns()
func (_CivilTCRContract *CivilTCRContractSession) GrantAppeal(listingAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.GrantAppeal(&_CivilTCRContract.TransactOpts, listingAddress)
}

// GrantAppeal is a paid mutator transaction binding the contract method 0x7d3c5d01.
//
// Solidity: function grantAppeal(listingAddress address) returns()
func (_CivilTCRContract *CivilTCRContractTransactorSession) GrantAppeal(listingAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.GrantAppeal(&_CivilTCRContract.TransactOpts, listingAddress)
}

// RequestAppeal is a paid mutator transaction binding the contract method 0x432261bb.
//
// Solidity: function requestAppeal(listingAddress address) returns()
func (_CivilTCRContract *CivilTCRContractTransactor) RequestAppeal(opts *bind.TransactOpts, listingAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.contract.Transact(opts, "requestAppeal", listingAddress)
}

// RequestAppeal is a paid mutator transaction binding the contract method 0x432261bb.
//
// Solidity: function requestAppeal(listingAddress address) returns()
func (_CivilTCRContract *CivilTCRContractSession) RequestAppeal(listingAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.RequestAppeal(&_CivilTCRContract.TransactOpts, listingAddress)
}

// RequestAppeal is a paid mutator transaction binding the contract method 0x432261bb.
//
// Solidity: function requestAppeal(listingAddress address) returns()
func (_CivilTCRContract *CivilTCRContractTransactorSession) RequestAppeal(listingAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.RequestAppeal(&_CivilTCRContract.TransactOpts, listingAddress)
}

// TransferGovernment is a paid mutator transaction binding the contract method 0x5b5d4e73.
//
// Solidity: function transferGovernment(newAddress address) returns()
func (_CivilTCRContract *CivilTCRContractTransactor) TransferGovernment(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.contract.Transact(opts, "transferGovernment", newAddress)
}

// TransferGovernment is a paid mutator transaction binding the contract method 0x5b5d4e73.
//
// Solidity: function transferGovernment(newAddress address) returns()
func (_CivilTCRContract *CivilTCRContractSession) TransferGovernment(newAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.TransferGovernment(&_CivilTCRContract.TransactOpts, newAddress)
}

// TransferGovernment is a paid mutator transaction binding the contract method 0x5b5d4e73.
//
// Solidity: function transferGovernment(newAddress address) returns()
func (_CivilTCRContract *CivilTCRContractTransactorSession) TransferGovernment(newAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.TransferGovernment(&_CivilTCRContract.TransactOpts, newAddress)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x01162b93.
//
// Solidity: function updateStatus(listingAddress address) returns()
func (_CivilTCRContract *CivilTCRContractTransactor) UpdateStatus(opts *bind.TransactOpts, listingAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.contract.Transact(opts, "updateStatus", listingAddress)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x01162b93.
//
// Solidity: function updateStatus(listingAddress address) returns()
func (_CivilTCRContract *CivilTCRContractSession) UpdateStatus(listingAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.UpdateStatus(&_CivilTCRContract.TransactOpts, listingAddress)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x01162b93.
//
// Solidity: function updateStatus(listingAddress address) returns()
func (_CivilTCRContract *CivilTCRContractTransactorSession) UpdateStatus(listingAddress common.Address) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.UpdateStatus(&_CivilTCRContract.TransactOpts, listingAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(listingAddress address, amount uint256) returns()
func (_CivilTCRContract *CivilTCRContractTransactor) Withdraw(opts *bind.TransactOpts, listingAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CivilTCRContract.contract.Transact(opts, "withdraw", listingAddress, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(listingAddress address, amount uint256) returns()
func (_CivilTCRContract *CivilTCRContractSession) Withdraw(listingAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.Withdraw(&_CivilTCRContract.TransactOpts, listingAddress, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(listingAddress address, amount uint256) returns()
func (_CivilTCRContract *CivilTCRContractTransactorSession) Withdraw(listingAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CivilTCRContract.Contract.Withdraw(&_CivilTCRContract.TransactOpts, listingAddress, amount)
}

// CivilTCRContractAppealGrantedIterator is returned from FilterAppealGranted and is used to iterate over the raw logs and unpacked data for AppealGranted events raised by the CivilTCRContract contract.
type CivilTCRContractAppealGrantedIterator struct {
	Event *CivilTCRContractAppealGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractAppealGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractAppealGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractAppealGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractAppealGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractAppealGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractAppealGranted represents a AppealGranted event raised by the CivilTCRContract contract.
type CivilTCRContractAppealGranted struct {
	ListingAddress common.Address
	ChallengeID    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAppealGranted is a free log retrieval operation binding the contract event 0x4142d097c264b24b2e32e9965e184ba773286757b5af93907d7b62022e1333af.
//
// Solidity: e _AppealGranted(listingAddress indexed address, challengeID indexed uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterAppealGranted(opts *bind.FilterOpts, listingAddress []common.Address, challengeID []*big.Int) (*CivilTCRContractAppealGrantedIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_AppealGranted", listingAddressRule, challengeIDRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractAppealGrantedIterator{contract: _CivilTCRContract.contract, event: "_AppealGranted", logs: logs, sub: sub}, nil
}

// WatchAppealGranted is a free log subscription operation binding the contract event 0x4142d097c264b24b2e32e9965e184ba773286757b5af93907d7b62022e1333af.
//
// Solidity: e _AppealGranted(listingAddress indexed address, challengeID indexed uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchAppealGranted(opts *bind.WatchOpts, sink chan<- *CivilTCRContractAppealGranted, listingAddress []common.Address, challengeID []*big.Int) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_AppealGranted", listingAddressRule, challengeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractAppealGranted)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_AppealGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractAppealRequestedIterator is returned from FilterAppealRequested and is used to iterate over the raw logs and unpacked data for AppealRequested events raised by the CivilTCRContract contract.
type CivilTCRContractAppealRequestedIterator struct {
	Event *CivilTCRContractAppealRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractAppealRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractAppealRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractAppealRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractAppealRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractAppealRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractAppealRequested represents a AppealRequested event raised by the CivilTCRContract contract.
type CivilTCRContractAppealRequested struct {
	ListingAddress common.Address
	ChallengeID    *big.Int
	AppealFeePaid  *big.Int
	Requester      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAppealRequested is a free log retrieval operation binding the contract event 0x7e8074c49cf8258160ac15b3e0fad39069cc31cec0b58a1b428d65e6a4e2ed46.
//
// Solidity: e _AppealRequested(listingAddress indexed address, challengeID indexed uint256, appealFeePaid uint256, requester address)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterAppealRequested(opts *bind.FilterOpts, listingAddress []common.Address, challengeID []*big.Int) (*CivilTCRContractAppealRequestedIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_AppealRequested", listingAddressRule, challengeIDRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractAppealRequestedIterator{contract: _CivilTCRContract.contract, event: "_AppealRequested", logs: logs, sub: sub}, nil
}

// WatchAppealRequested is a free log subscription operation binding the contract event 0x7e8074c49cf8258160ac15b3e0fad39069cc31cec0b58a1b428d65e6a4e2ed46.
//
// Solidity: e _AppealRequested(listingAddress indexed address, challengeID indexed uint256, appealFeePaid uint256, requester address)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchAppealRequested(opts *bind.WatchOpts, sink chan<- *CivilTCRContractAppealRequested, listingAddress []common.Address, challengeID []*big.Int) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_AppealRequested", listingAddressRule, challengeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractAppealRequested)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_AppealRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractApplicationIterator is returned from FilterApplication and is used to iterate over the raw logs and unpacked data for Application events raised by the CivilTCRContract contract.
type CivilTCRContractApplicationIterator struct {
	Event *CivilTCRContractApplication // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractApplicationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractApplication)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractApplication)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractApplicationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractApplicationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractApplication represents a Application event raised by the CivilTCRContract contract.
type CivilTCRContractApplication struct {
	ListingAddress common.Address
	Deposit        *big.Int
	AppEndDate     *big.Int
	Data           string
	Applicant      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterApplication is a free log retrieval operation binding the contract event 0x09cd8dcaf170a50a26316b5fe0727dd9fb9581a688d65e758b16a1650da65c0b.
//
// Solidity: e _Application(listingAddress indexed address, deposit uint256, appEndDate uint256, data string, applicant indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterApplication(opts *bind.FilterOpts, listingAddress []common.Address, applicant []common.Address) (*CivilTCRContractApplicationIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	var applicantRule []interface{}
	for _, applicantItem := range applicant {
		applicantRule = append(applicantRule, applicantItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_Application", listingAddressRule, applicantRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractApplicationIterator{contract: _CivilTCRContract.contract, event: "_Application", logs: logs, sub: sub}, nil
}

// WatchApplication is a free log subscription operation binding the contract event 0x09cd8dcaf170a50a26316b5fe0727dd9fb9581a688d65e758b16a1650da65c0b.
//
// Solidity: e _Application(listingAddress indexed address, deposit uint256, appEndDate uint256, data string, applicant indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchApplication(opts *bind.WatchOpts, sink chan<- *CivilTCRContractApplication, listingAddress []common.Address, applicant []common.Address) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	var applicantRule []interface{}
	for _, applicantItem := range applicant {
		applicantRule = append(applicantRule, applicantItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_Application", listingAddressRule, applicantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractApplication)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_Application", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractApplicationRemovedIterator is returned from FilterApplicationRemoved and is used to iterate over the raw logs and unpacked data for ApplicationRemoved events raised by the CivilTCRContract contract.
type CivilTCRContractApplicationRemovedIterator struct {
	Event *CivilTCRContractApplicationRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractApplicationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractApplicationRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractApplicationRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractApplicationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractApplicationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractApplicationRemoved represents a ApplicationRemoved event raised by the CivilTCRContract contract.
type CivilTCRContractApplicationRemoved struct {
	ListingAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterApplicationRemoved is a free log retrieval operation binding the contract event 0x8ad9ca8735c55207756208e8f59c7693e83d234fc6c8af2713f266468edff871.
//
// Solidity: e _ApplicationRemoved(listingAddress indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterApplicationRemoved(opts *bind.FilterOpts, listingAddress []common.Address) (*CivilTCRContractApplicationRemovedIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_ApplicationRemoved", listingAddressRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractApplicationRemovedIterator{contract: _CivilTCRContract.contract, event: "_ApplicationRemoved", logs: logs, sub: sub}, nil
}

// WatchApplicationRemoved is a free log subscription operation binding the contract event 0x8ad9ca8735c55207756208e8f59c7693e83d234fc6c8af2713f266468edff871.
//
// Solidity: e _ApplicationRemoved(listingAddress indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchApplicationRemoved(opts *bind.WatchOpts, sink chan<- *CivilTCRContractApplicationRemoved, listingAddress []common.Address) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_ApplicationRemoved", listingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractApplicationRemoved)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_ApplicationRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractApplicationWhitelistedIterator is returned from FilterApplicationWhitelisted and is used to iterate over the raw logs and unpacked data for ApplicationWhitelisted events raised by the CivilTCRContract contract.
type CivilTCRContractApplicationWhitelistedIterator struct {
	Event *CivilTCRContractApplicationWhitelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractApplicationWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractApplicationWhitelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractApplicationWhitelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractApplicationWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractApplicationWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractApplicationWhitelisted represents a ApplicationWhitelisted event raised by the CivilTCRContract contract.
type CivilTCRContractApplicationWhitelisted struct {
	ListingAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterApplicationWhitelisted is a free log retrieval operation binding the contract event 0xb268dc7f76f496fd307b40e0a3b57631a7e46123d9f708b1573bd4efbba3bdba.
//
// Solidity: e _ApplicationWhitelisted(listingAddress indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterApplicationWhitelisted(opts *bind.FilterOpts, listingAddress []common.Address) (*CivilTCRContractApplicationWhitelistedIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_ApplicationWhitelisted", listingAddressRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractApplicationWhitelistedIterator{contract: _CivilTCRContract.contract, event: "_ApplicationWhitelisted", logs: logs, sub: sub}, nil
}

// WatchApplicationWhitelisted is a free log subscription operation binding the contract event 0xb268dc7f76f496fd307b40e0a3b57631a7e46123d9f708b1573bd4efbba3bdba.
//
// Solidity: e _ApplicationWhitelisted(listingAddress indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchApplicationWhitelisted(opts *bind.WatchOpts, sink chan<- *CivilTCRContractApplicationWhitelisted, listingAddress []common.Address) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_ApplicationWhitelisted", listingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractApplicationWhitelisted)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_ApplicationWhitelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractChallengeIterator is returned from FilterChallenge and is used to iterate over the raw logs and unpacked data for Challenge events raised by the CivilTCRContract contract.
type CivilTCRContractChallengeIterator struct {
	Event *CivilTCRContractChallenge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractChallenge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractChallenge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractChallenge represents a Challenge event raised by the CivilTCRContract contract.
type CivilTCRContractChallenge struct {
	ListingAddress common.Address
	ChallengeID    *big.Int
	Data           string
	CommitEndDate  *big.Int
	RevealEndDate  *big.Int
	Challenger     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallenge is a free log retrieval operation binding the contract event 0x9a8e3864cbacafc5547b8567796b4d12d51217a879192b61932f5151ce581310.
//
// Solidity: e _Challenge(listingAddress indexed address, challengeID indexed uint256, data string, commitEndDate uint256, revealEndDate uint256, challenger indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterChallenge(opts *bind.FilterOpts, listingAddress []common.Address, challengeID []*big.Int, challenger []common.Address) (*CivilTCRContractChallengeIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_Challenge", listingAddressRule, challengeIDRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractChallengeIterator{contract: _CivilTCRContract.contract, event: "_Challenge", logs: logs, sub: sub}, nil
}

// WatchChallenge is a free log subscription operation binding the contract event 0x9a8e3864cbacafc5547b8567796b4d12d51217a879192b61932f5151ce581310.
//
// Solidity: e _Challenge(listingAddress indexed address, challengeID indexed uint256, data string, commitEndDate uint256, revealEndDate uint256, challenger indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchChallenge(opts *bind.WatchOpts, sink chan<- *CivilTCRContractChallenge, listingAddress []common.Address, challengeID []*big.Int, challenger []common.Address) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_Challenge", listingAddressRule, challengeIDRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractChallenge)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_Challenge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractChallengeFailedIterator is returned from FilterChallengeFailed and is used to iterate over the raw logs and unpacked data for ChallengeFailed events raised by the CivilTCRContract contract.
type CivilTCRContractChallengeFailedIterator struct {
	Event *CivilTCRContractChallengeFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractChallengeFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractChallengeFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractChallengeFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractChallengeFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractChallengeFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractChallengeFailed represents a ChallengeFailed event raised by the CivilTCRContract contract.
type CivilTCRContractChallengeFailed struct {
	ListingAddress common.Address
	ChallengeID    *big.Int
	RewardPool     *big.Int
	TotalTokens    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeFailed is a free log retrieval operation binding the contract event 0x3639145ca0a6a8008912a972730b5c8634e1fd1555ea44a257a8de53c30d72fb.
//
// Solidity: e _ChallengeFailed(listingAddress indexed address, challengeID indexed uint256, rewardPool uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterChallengeFailed(opts *bind.FilterOpts, listingAddress []common.Address, challengeID []*big.Int) (*CivilTCRContractChallengeFailedIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_ChallengeFailed", listingAddressRule, challengeIDRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractChallengeFailedIterator{contract: _CivilTCRContract.contract, event: "_ChallengeFailed", logs: logs, sub: sub}, nil
}

// WatchChallengeFailed is a free log subscription operation binding the contract event 0x3639145ca0a6a8008912a972730b5c8634e1fd1555ea44a257a8de53c30d72fb.
//
// Solidity: e _ChallengeFailed(listingAddress indexed address, challengeID indexed uint256, rewardPool uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchChallengeFailed(opts *bind.WatchOpts, sink chan<- *CivilTCRContractChallengeFailed, listingAddress []common.Address, challengeID []*big.Int) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_ChallengeFailed", listingAddressRule, challengeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractChallengeFailed)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_ChallengeFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractChallengeSucceededIterator is returned from FilterChallengeSucceeded and is used to iterate over the raw logs and unpacked data for ChallengeSucceeded events raised by the CivilTCRContract contract.
type CivilTCRContractChallengeSucceededIterator struct {
	Event *CivilTCRContractChallengeSucceeded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractChallengeSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractChallengeSucceeded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractChallengeSucceeded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractChallengeSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractChallengeSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractChallengeSucceeded represents a ChallengeSucceeded event raised by the CivilTCRContract contract.
type CivilTCRContractChallengeSucceeded struct {
	ListingAddress common.Address
	ChallengeID    *big.Int
	RewardPool     *big.Int
	TotalTokens    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeSucceeded is a free log retrieval operation binding the contract event 0xe86031b52c5a57c0768c3f196b63abf60b5ed012de77ce1bb88fc63588f7603a.
//
// Solidity: e _ChallengeSucceeded(listingAddress indexed address, challengeID indexed uint256, rewardPool uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterChallengeSucceeded(opts *bind.FilterOpts, listingAddress []common.Address, challengeID []*big.Int) (*CivilTCRContractChallengeSucceededIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_ChallengeSucceeded", listingAddressRule, challengeIDRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractChallengeSucceededIterator{contract: _CivilTCRContract.contract, event: "_ChallengeSucceeded", logs: logs, sub: sub}, nil
}

// WatchChallengeSucceeded is a free log subscription operation binding the contract event 0xe86031b52c5a57c0768c3f196b63abf60b5ed012de77ce1bb88fc63588f7603a.
//
// Solidity: e _ChallengeSucceeded(listingAddress indexed address, challengeID indexed uint256, rewardPool uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchChallengeSucceeded(opts *bind.WatchOpts, sink chan<- *CivilTCRContractChallengeSucceeded, listingAddress []common.Address, challengeID []*big.Int) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_ChallengeSucceeded", listingAddressRule, challengeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractChallengeSucceeded)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_ChallengeSucceeded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the CivilTCRContract contract.
type CivilTCRContractDepositIterator struct {
	Event *CivilTCRContractDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractDeposit represents a Deposit event raised by the CivilTCRContract contract.
type CivilTCRContractDeposit struct {
	ListingAddress common.Address
	Added          *big.Int
	NewTotal       *big.Int
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xfc2e298800eb7bcacdea96213f53962a6bdf27d2a560f831d4e42301492e8f6a.
//
// Solidity: e _Deposit(listingAddress indexed address, added uint256, newTotal uint256, owner indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterDeposit(opts *bind.FilterOpts, listingAddress []common.Address, owner []common.Address) (*CivilTCRContractDepositIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_Deposit", listingAddressRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractDepositIterator{contract: _CivilTCRContract.contract, event: "_Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xfc2e298800eb7bcacdea96213f53962a6bdf27d2a560f831d4e42301492e8f6a.
//
// Solidity: e _Deposit(listingAddress indexed address, added uint256, newTotal uint256, owner indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CivilTCRContractDeposit, listingAddress []common.Address, owner []common.Address) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_Deposit", listingAddressRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractDeposit)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractFailedChallengeOverturnedIterator is returned from FilterFailedChallengeOverturned and is used to iterate over the raw logs and unpacked data for FailedChallengeOverturned events raised by the CivilTCRContract contract.
type CivilTCRContractFailedChallengeOverturnedIterator struct {
	Event *CivilTCRContractFailedChallengeOverturned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractFailedChallengeOverturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractFailedChallengeOverturned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractFailedChallengeOverturned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractFailedChallengeOverturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractFailedChallengeOverturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractFailedChallengeOverturned represents a FailedChallengeOverturned event raised by the CivilTCRContract contract.
type CivilTCRContractFailedChallengeOverturned struct {
	ListingAddress common.Address
	ChallengeID    *big.Int
	RewardPool     *big.Int
	TotalTokens    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFailedChallengeOverturned is a free log retrieval operation binding the contract event 0x446922bbfdaa528d4a969857cd0894d6bf8bbff52226624e752b3f1be7513b0a.
//
// Solidity: e _FailedChallengeOverturned(listingAddress indexed address, challengeID indexed uint256, rewardPool uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterFailedChallengeOverturned(opts *bind.FilterOpts, listingAddress []common.Address, challengeID []*big.Int) (*CivilTCRContractFailedChallengeOverturnedIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_FailedChallengeOverturned", listingAddressRule, challengeIDRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractFailedChallengeOverturnedIterator{contract: _CivilTCRContract.contract, event: "_FailedChallengeOverturned", logs: logs, sub: sub}, nil
}

// WatchFailedChallengeOverturned is a free log subscription operation binding the contract event 0x446922bbfdaa528d4a969857cd0894d6bf8bbff52226624e752b3f1be7513b0a.
//
// Solidity: e _FailedChallengeOverturned(listingAddress indexed address, challengeID indexed uint256, rewardPool uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchFailedChallengeOverturned(opts *bind.WatchOpts, sink chan<- *CivilTCRContractFailedChallengeOverturned, listingAddress []common.Address, challengeID []*big.Int) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_FailedChallengeOverturned", listingAddressRule, challengeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractFailedChallengeOverturned)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_FailedChallengeOverturned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractGovernmentTransferedIterator is returned from FilterGovernmentTransfered and is used to iterate over the raw logs and unpacked data for GovernmentTransfered events raised by the CivilTCRContract contract.
type CivilTCRContractGovernmentTransferedIterator struct {
	Event *CivilTCRContractGovernmentTransfered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractGovernmentTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractGovernmentTransfered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractGovernmentTransfered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractGovernmentTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractGovernmentTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractGovernmentTransfered represents a GovernmentTransfered event raised by the CivilTCRContract contract.
type CivilTCRContractGovernmentTransfered struct {
	NewGovernment common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGovernmentTransfered is a free log retrieval operation binding the contract event 0x016b4781993f669e6eac42012fead2d96f8381769b4efbb4ad686cca6031ea88.
//
// Solidity: e _GovernmentTransfered(newGovernment address)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterGovernmentTransfered(opts *bind.FilterOpts) (*CivilTCRContractGovernmentTransferedIterator, error) {

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_GovernmentTransfered")
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractGovernmentTransferedIterator{contract: _CivilTCRContract.contract, event: "_GovernmentTransfered", logs: logs, sub: sub}, nil
}

// WatchGovernmentTransfered is a free log subscription operation binding the contract event 0x016b4781993f669e6eac42012fead2d96f8381769b4efbb4ad686cca6031ea88.
//
// Solidity: e _GovernmentTransfered(newGovernment address)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchGovernmentTransfered(opts *bind.WatchOpts, sink chan<- *CivilTCRContractGovernmentTransfered) (event.Subscription, error) {

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_GovernmentTransfered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractGovernmentTransfered)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_GovernmentTransfered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractGrantedAppealChallengedIterator is returned from FilterGrantedAppealChallenged and is used to iterate over the raw logs and unpacked data for GrantedAppealChallenged events raised by the CivilTCRContract contract.
type CivilTCRContractGrantedAppealChallengedIterator struct {
	Event *CivilTCRContractGrantedAppealChallenged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractGrantedAppealChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractGrantedAppealChallenged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractGrantedAppealChallenged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractGrantedAppealChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractGrantedAppealChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractGrantedAppealChallenged represents a GrantedAppealChallenged event raised by the CivilTCRContract contract.
type CivilTCRContractGrantedAppealChallenged struct {
	ListingAddress    common.Address
	ChallengeID       *big.Int
	AppealChallengeID *big.Int
	Data              string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGrantedAppealChallenged is a free log retrieval operation binding the contract event 0xedfe36bf00610fb3b5474f1efd2de0d52ffb9a50b056ee37c33cea805fd44161.
//
// Solidity: e _GrantedAppealChallenged(listingAddress indexed address, challengeID indexed uint256, appealChallengeID indexed uint256, data string)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterGrantedAppealChallenged(opts *bind.FilterOpts, listingAddress []common.Address, challengeID []*big.Int, appealChallengeID []*big.Int) (*CivilTCRContractGrantedAppealChallengedIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}
	var appealChallengeIDRule []interface{}
	for _, appealChallengeIDItem := range appealChallengeID {
		appealChallengeIDRule = append(appealChallengeIDRule, appealChallengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_GrantedAppealChallenged", listingAddressRule, challengeIDRule, appealChallengeIDRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractGrantedAppealChallengedIterator{contract: _CivilTCRContract.contract, event: "_GrantedAppealChallenged", logs: logs, sub: sub}, nil
}

// WatchGrantedAppealChallenged is a free log subscription operation binding the contract event 0xedfe36bf00610fb3b5474f1efd2de0d52ffb9a50b056ee37c33cea805fd44161.
//
// Solidity: e _GrantedAppealChallenged(listingAddress indexed address, challengeID indexed uint256, appealChallengeID indexed uint256, data string)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchGrantedAppealChallenged(opts *bind.WatchOpts, sink chan<- *CivilTCRContractGrantedAppealChallenged, listingAddress []common.Address, challengeID []*big.Int, appealChallengeID []*big.Int) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}
	var appealChallengeIDRule []interface{}
	for _, appealChallengeIDItem := range appealChallengeID {
		appealChallengeIDRule = append(appealChallengeIDRule, appealChallengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_GrantedAppealChallenged", listingAddressRule, challengeIDRule, appealChallengeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractGrantedAppealChallenged)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_GrantedAppealChallenged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractGrantedAppealConfirmedIterator is returned from FilterGrantedAppealConfirmed and is used to iterate over the raw logs and unpacked data for GrantedAppealConfirmed events raised by the CivilTCRContract contract.
type CivilTCRContractGrantedAppealConfirmedIterator struct {
	Event *CivilTCRContractGrantedAppealConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractGrantedAppealConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractGrantedAppealConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractGrantedAppealConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractGrantedAppealConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractGrantedAppealConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractGrantedAppealConfirmed represents a GrantedAppealConfirmed event raised by the CivilTCRContract contract.
type CivilTCRContractGrantedAppealConfirmed struct {
	ListingAddress    common.Address
	ChallengeID       *big.Int
	AppealChallengeID *big.Int
	RewardPool        *big.Int
	TotalTokens       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGrantedAppealConfirmed is a free log retrieval operation binding the contract event 0x8a7e8d1076fec4f93e4d57111b034ab3975009f601977350c4542e15d2e8b0f6.
//
// Solidity: e _GrantedAppealConfirmed(listingAddress indexed address, challengeID indexed uint256, appealChallengeID indexed uint256, rewardPool uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterGrantedAppealConfirmed(opts *bind.FilterOpts, listingAddress []common.Address, challengeID []*big.Int, appealChallengeID []*big.Int) (*CivilTCRContractGrantedAppealConfirmedIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}
	var appealChallengeIDRule []interface{}
	for _, appealChallengeIDItem := range appealChallengeID {
		appealChallengeIDRule = append(appealChallengeIDRule, appealChallengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_GrantedAppealConfirmed", listingAddressRule, challengeIDRule, appealChallengeIDRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractGrantedAppealConfirmedIterator{contract: _CivilTCRContract.contract, event: "_GrantedAppealConfirmed", logs: logs, sub: sub}, nil
}

// WatchGrantedAppealConfirmed is a free log subscription operation binding the contract event 0x8a7e8d1076fec4f93e4d57111b034ab3975009f601977350c4542e15d2e8b0f6.
//
// Solidity: e _GrantedAppealConfirmed(listingAddress indexed address, challengeID indexed uint256, appealChallengeID indexed uint256, rewardPool uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchGrantedAppealConfirmed(opts *bind.WatchOpts, sink chan<- *CivilTCRContractGrantedAppealConfirmed, listingAddress []common.Address, challengeID []*big.Int, appealChallengeID []*big.Int) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}
	var appealChallengeIDRule []interface{}
	for _, appealChallengeIDItem := range appealChallengeID {
		appealChallengeIDRule = append(appealChallengeIDRule, appealChallengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_GrantedAppealConfirmed", listingAddressRule, challengeIDRule, appealChallengeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractGrantedAppealConfirmed)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_GrantedAppealConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractGrantedAppealOverturnedIterator is returned from FilterGrantedAppealOverturned and is used to iterate over the raw logs and unpacked data for GrantedAppealOverturned events raised by the CivilTCRContract contract.
type CivilTCRContractGrantedAppealOverturnedIterator struct {
	Event *CivilTCRContractGrantedAppealOverturned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractGrantedAppealOverturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractGrantedAppealOverturned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractGrantedAppealOverturned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractGrantedAppealOverturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractGrantedAppealOverturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractGrantedAppealOverturned represents a GrantedAppealOverturned event raised by the CivilTCRContract contract.
type CivilTCRContractGrantedAppealOverturned struct {
	ListingAddress    common.Address
	ChallengeID       *big.Int
	AppealChallengeID *big.Int
	RewardPool        *big.Int
	TotalTokens       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGrantedAppealOverturned is a free log retrieval operation binding the contract event 0xc49556ab8bcbdd0403e98b6dac260ac86008640cda2a5a229c895353b87f2feb.
//
// Solidity: e _GrantedAppealOverturned(listingAddress indexed address, challengeID indexed uint256, appealChallengeID indexed uint256, rewardPool uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterGrantedAppealOverturned(opts *bind.FilterOpts, listingAddress []common.Address, challengeID []*big.Int, appealChallengeID []*big.Int) (*CivilTCRContractGrantedAppealOverturnedIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}
	var appealChallengeIDRule []interface{}
	for _, appealChallengeIDItem := range appealChallengeID {
		appealChallengeIDRule = append(appealChallengeIDRule, appealChallengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_GrantedAppealOverturned", listingAddressRule, challengeIDRule, appealChallengeIDRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractGrantedAppealOverturnedIterator{contract: _CivilTCRContract.contract, event: "_GrantedAppealOverturned", logs: logs, sub: sub}, nil
}

// WatchGrantedAppealOverturned is a free log subscription operation binding the contract event 0xc49556ab8bcbdd0403e98b6dac260ac86008640cda2a5a229c895353b87f2feb.
//
// Solidity: e _GrantedAppealOverturned(listingAddress indexed address, challengeID indexed uint256, appealChallengeID indexed uint256, rewardPool uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchGrantedAppealOverturned(opts *bind.WatchOpts, sink chan<- *CivilTCRContractGrantedAppealOverturned, listingAddress []common.Address, challengeID []*big.Int, appealChallengeID []*big.Int) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}
	var appealChallengeIDRule []interface{}
	for _, appealChallengeIDItem := range appealChallengeID {
		appealChallengeIDRule = append(appealChallengeIDRule, appealChallengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_GrantedAppealOverturned", listingAddressRule, challengeIDRule, appealChallengeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractGrantedAppealOverturned)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_GrantedAppealOverturned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractListingRemovedIterator is returned from FilterListingRemoved and is used to iterate over the raw logs and unpacked data for ListingRemoved events raised by the CivilTCRContract contract.
type CivilTCRContractListingRemovedIterator struct {
	Event *CivilTCRContractListingRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractListingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractListingRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractListingRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractListingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractListingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractListingRemoved represents a ListingRemoved event raised by the CivilTCRContract contract.
type CivilTCRContractListingRemoved struct {
	ListingAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterListingRemoved is a free log retrieval operation binding the contract event 0x5aebb54d5afe29103adbc464fd4e0313af619f4d19f10a0209128b76cd9d0b07.
//
// Solidity: e _ListingRemoved(listingAddress indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterListingRemoved(opts *bind.FilterOpts, listingAddress []common.Address) (*CivilTCRContractListingRemovedIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_ListingRemoved", listingAddressRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractListingRemovedIterator{contract: _CivilTCRContract.contract, event: "_ListingRemoved", logs: logs, sub: sub}, nil
}

// WatchListingRemoved is a free log subscription operation binding the contract event 0x5aebb54d5afe29103adbc464fd4e0313af619f4d19f10a0209128b76cd9d0b07.
//
// Solidity: e _ListingRemoved(listingAddress indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchListingRemoved(opts *bind.WatchOpts, sink chan<- *CivilTCRContractListingRemoved, listingAddress []common.Address) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_ListingRemoved", listingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractListingRemoved)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_ListingRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractListingWithdrawnIterator is returned from FilterListingWithdrawn and is used to iterate over the raw logs and unpacked data for ListingWithdrawn events raised by the CivilTCRContract contract.
type CivilTCRContractListingWithdrawnIterator struct {
	Event *CivilTCRContractListingWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractListingWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractListingWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractListingWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractListingWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractListingWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractListingWithdrawn represents a ListingWithdrawn event raised by the CivilTCRContract contract.
type CivilTCRContractListingWithdrawn struct {
	ListingAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterListingWithdrawn is a free log retrieval operation binding the contract event 0x09a024f7311a15ac363521bddca1d9937c4244ab9a25e6c968e610b35ecc7503.
//
// Solidity: e _ListingWithdrawn(listingAddress indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterListingWithdrawn(opts *bind.FilterOpts, listingAddress []common.Address) (*CivilTCRContractListingWithdrawnIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_ListingWithdrawn", listingAddressRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractListingWithdrawnIterator{contract: _CivilTCRContract.contract, event: "_ListingWithdrawn", logs: logs, sub: sub}, nil
}

// WatchListingWithdrawn is a free log subscription operation binding the contract event 0x09a024f7311a15ac363521bddca1d9937c4244ab9a25e6c968e610b35ecc7503.
//
// Solidity: e _ListingWithdrawn(listingAddress indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchListingWithdrawn(opts *bind.WatchOpts, sink chan<- *CivilTCRContractListingWithdrawn, listingAddress []common.Address) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_ListingWithdrawn", listingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractListingWithdrawn)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_ListingWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the CivilTCRContract contract.
type CivilTCRContractRewardClaimedIterator struct {
	Event *CivilTCRContractRewardClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractRewardClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractRewardClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractRewardClaimed represents a RewardClaimed event raised by the CivilTCRContract contract.
type CivilTCRContractRewardClaimed struct {
	ChallengeID *big.Int
	Reward      *big.Int
	Voter       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x6f4c982acc31b0af2cf1dc1556f21c0325d893782d65e83c68a5534a33f59957.
//
// Solidity: e _RewardClaimed(challengeID indexed uint256, reward uint256, voter indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterRewardClaimed(opts *bind.FilterOpts, challengeID []*big.Int, voter []common.Address) (*CivilTCRContractRewardClaimedIterator, error) {

	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_RewardClaimed", challengeIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractRewardClaimedIterator{contract: _CivilTCRContract.contract, event: "_RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x6f4c982acc31b0af2cf1dc1556f21c0325d893782d65e83c68a5534a33f59957.
//
// Solidity: e _RewardClaimed(challengeID indexed uint256, reward uint256, voter indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *CivilTCRContractRewardClaimed, challengeID []*big.Int, voter []common.Address) (event.Subscription, error) {

	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_RewardClaimed", challengeIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractRewardClaimed)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_RewardClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractSuccessfulChallengeOverturnedIterator is returned from FilterSuccessfulChallengeOverturned and is used to iterate over the raw logs and unpacked data for SuccessfulChallengeOverturned events raised by the CivilTCRContract contract.
type CivilTCRContractSuccessfulChallengeOverturnedIterator struct {
	Event *CivilTCRContractSuccessfulChallengeOverturned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractSuccessfulChallengeOverturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractSuccessfulChallengeOverturned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractSuccessfulChallengeOverturned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractSuccessfulChallengeOverturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractSuccessfulChallengeOverturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractSuccessfulChallengeOverturned represents a SuccessfulChallengeOverturned event raised by the CivilTCRContract contract.
type CivilTCRContractSuccessfulChallengeOverturned struct {
	ListingAddress common.Address
	ChallengeID    *big.Int
	RewardPool     *big.Int
	TotalTokens    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSuccessfulChallengeOverturned is a free log retrieval operation binding the contract event 0x72506b3ce4d8f0cf8cf6ccb7cd5281af2b0d020121fb20abfa073eeb3f6d296e.
//
// Solidity: e _SuccessfulChallengeOverturned(listingAddress indexed address, challengeID indexed uint256, rewardPool uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterSuccessfulChallengeOverturned(opts *bind.FilterOpts, listingAddress []common.Address, challengeID []*big.Int) (*CivilTCRContractSuccessfulChallengeOverturnedIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_SuccessfulChallengeOverturned", listingAddressRule, challengeIDRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractSuccessfulChallengeOverturnedIterator{contract: _CivilTCRContract.contract, event: "_SuccessfulChallengeOverturned", logs: logs, sub: sub}, nil
}

// WatchSuccessfulChallengeOverturned is a free log subscription operation binding the contract event 0x72506b3ce4d8f0cf8cf6ccb7cd5281af2b0d020121fb20abfa073eeb3f6d296e.
//
// Solidity: e _SuccessfulChallengeOverturned(listingAddress indexed address, challengeID indexed uint256, rewardPool uint256, totalTokens uint256)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchSuccessfulChallengeOverturned(opts *bind.WatchOpts, sink chan<- *CivilTCRContractSuccessfulChallengeOverturned, listingAddress []common.Address, challengeID []*big.Int) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}
	var challengeIDRule []interface{}
	for _, challengeIDItem := range challengeID {
		challengeIDRule = append(challengeIDRule, challengeIDItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_SuccessfulChallengeOverturned", listingAddressRule, challengeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractSuccessfulChallengeOverturned)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_SuccessfulChallengeOverturned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractTouchAndRemovedIterator is returned from FilterTouchAndRemoved and is used to iterate over the raw logs and unpacked data for TouchAndRemoved events raised by the CivilTCRContract contract.
type CivilTCRContractTouchAndRemovedIterator struct {
	Event *CivilTCRContractTouchAndRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractTouchAndRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractTouchAndRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractTouchAndRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractTouchAndRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractTouchAndRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractTouchAndRemoved represents a TouchAndRemoved event raised by the CivilTCRContract contract.
type CivilTCRContractTouchAndRemoved struct {
	ListingAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTouchAndRemoved is a free log retrieval operation binding the contract event 0xc88462fa6972b64560d1dd34c4d66f2ff9841b2f974bdb0fab0827133d692f64.
//
// Solidity: e _TouchAndRemoved(listingAddress indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterTouchAndRemoved(opts *bind.FilterOpts, listingAddress []common.Address) (*CivilTCRContractTouchAndRemovedIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_TouchAndRemoved", listingAddressRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractTouchAndRemovedIterator{contract: _CivilTCRContract.contract, event: "_TouchAndRemoved", logs: logs, sub: sub}, nil
}

// WatchTouchAndRemoved is a free log subscription operation binding the contract event 0xc88462fa6972b64560d1dd34c4d66f2ff9841b2f974bdb0fab0827133d692f64.
//
// Solidity: e _TouchAndRemoved(listingAddress indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchTouchAndRemoved(opts *bind.WatchOpts, sink chan<- *CivilTCRContractTouchAndRemoved, listingAddress []common.Address) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_TouchAndRemoved", listingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractTouchAndRemoved)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_TouchAndRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CivilTCRContractWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the CivilTCRContract contract.
type CivilTCRContractWithdrawalIterator struct {
	Event *CivilTCRContractWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CivilTCRContractWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CivilTCRContractWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CivilTCRContractWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CivilTCRContractWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CivilTCRContractWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CivilTCRContractWithdrawal represents a Withdrawal event raised by the CivilTCRContract contract.
type CivilTCRContractWithdrawal struct {
	ListingAddress common.Address
	Withdrew       *big.Int
	NewTotal       *big.Int
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7b7771adeec078e4a338f627a52f307a7fd66d915cb133b5ace441bed26abc0b.
//
// Solidity: e _Withdrawal(listingAddress indexed address, withdrew uint256, newTotal uint256, owner indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) FilterWithdrawal(opts *bind.FilterOpts, listingAddress []common.Address, owner []common.Address) (*CivilTCRContractWithdrawalIterator, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CivilTCRContract.contract.FilterLogs(opts, "_Withdrawal", listingAddressRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CivilTCRContractWithdrawalIterator{contract: _CivilTCRContract.contract, event: "_Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7b7771adeec078e4a338f627a52f307a7fd66d915cb133b5ace441bed26abc0b.
//
// Solidity: e _Withdrawal(listingAddress indexed address, withdrew uint256, newTotal uint256, owner indexed address)
func (_CivilTCRContract *CivilTCRContractFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *CivilTCRContractWithdrawal, listingAddress []common.Address, owner []common.Address) (event.Subscription, error) {

	var listingAddressRule []interface{}
	for _, listingAddressItem := range listingAddress {
		listingAddressRule = append(listingAddressRule, listingAddressItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CivilTCRContract.contract.WatchLogs(opts, "_Withdrawal", listingAddressRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CivilTCRContractWithdrawal)
				if err := _CivilTCRContract.contract.UnpackLog(event, "_Withdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
