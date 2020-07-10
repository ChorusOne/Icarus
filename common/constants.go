package common

//Paths for IPC and RPC connections to a Celo full node
//Currently hardcoded
//TODO : Set these via config.env
var (
	IPCFilePath = "/home/celo/.celo/geth.ipc"
	RPCURLPath  = ""
)

//Constant-ifying names of Celo network
const (
	TestnetBaklava   = "baklava"
	TestnetAlfajores = "alfajores"
	Mainnet          = "mainnet"
)

//NetActive refers to the Celo network that Icarus will extract data for
//Defaults to Mainnet.
var NetActive = Mainnet

//Constant-ifying names of contracts
const (
	GoldToken         = "GoldToken"
	LockedGold        = "LockedGold"
	StableToken       = "StableToken"
	Election          = "Election"
	Governance        = "Governance"
	Validators        = "Validators"
	DowntimeSlasher   = "DowntimeSlasher"
	Accounts          = "Accounts"
	GovernanceSlasher = "GovernanceSlasher"
)

//WrapperContractDeploymentAddress maps contracts to their deployment addresses
//for different networks of the Celo blockchain
var WrapperContractDeploymentAddress = map[string]map[string]string{
	TestnetAlfajores: {},

	TestnetBaklava: {
		DowntimeSlasher: "0x198D951Ec0E530C24A6b01714E1b1f1f0d6996E7",
		LockedGold:      "0xd01E94451aA66930Fb76287D502e6dc1689464FC",
		StableToken:     "0x4b84c2EF94A274DbF83E2F1Ec1608456c9B62D96",
		Election:        "0x7457F05389Bb197e3695E85994DE4b91F16B97Db",
		GoldToken:       "0x44f434E83A3179FCede28941b3a81953fb575217",
		Governance:      "0xfBBA5B0df7A52814301a4750Dab628b2F35f30f7",
		Validators:      "0x2870054a9F889655CC832763D01b438f3AA7b91D",
	},

	Mainnet: {
		DowntimeSlasher:   "0xeFE50f83bA23240A85c39afF429b31e556d2C80F",
		LockedGold:        "0x6cC083Aed9e3ebe302A6336dBC7c921C9f03349E",
		StableToken:       "0x765DE816845861e75A25fCA122bb6898B8B1282a",
		Election:          "0x8D6677192144292870907E3Fa8A5527fE55A7ff6",
		GoldToken:         "0x471EcE3750Da237f93B8E339c536989b8978a438",
		Governance:        "0xD533Ca259b330c7A88f74E000a3FaEa2d63B7972",
		Validators:        "0xaEb865bCa93DdC8F47b8e29F40C5399cE34d0C58",
		Accounts:          "0x7d21685C17607338b313a7174bAb6620baD0aaB7",
		GovernanceSlasher: "0x68f04Ab73B93f5175207296528454999475294d5",
	},
}

//DeploymentBlockNumber maps contracts to their deployment heights
//for different networks of the Celo blockchain
var DeploymentBlockNumber = map[string]map[string]int64{

	TestnetAlfajores: {
		GoldToken:  107,
		LockedGold: 158,
	},

	TestnetBaklava: {
		LockedGold:  706,
		StableToken: 682,
		Election:    718,
		GoldToken:   648,
		Governance:  782,
		Validators:  712,
	},
	Mainnet: {
		LockedGold:        3001,
		StableToken:       2957,
		Election:          3016,
		GoldToken:         2919,
		Governance:        3090,
		Validators:        3008,
		Accounts:          2993,
		GovernanceSlasher: 3061,
	},
}
