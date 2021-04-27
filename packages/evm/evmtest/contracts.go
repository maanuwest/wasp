package evmtest

import "github.com/ethereum/go-ethereum/common"

// pragma solidity >=0.7.0 <0.8.0;
//
// contract Storage {
//     uint32 n;
//
//     constructor(uint32 _n) {
//         n = _n;
//     }
//
//     function store(uint32 _n) public {
//         n = _n;
//     }
//
//     function retrieve() public view returns (uint32){
//         return n;
//     }
// }

const StorageContractABI = `
[
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "_n",
				"type": "uint32"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"inputs": [],
		"name": "retrieve",
		"outputs": [
			{
				"internalType": "uint32",
				"name": "",
				"type": "uint32"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "_n",
				"type": "uint32"
			}
		],
		"name": "store",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
`

var StorageContractBytecode = common.FromHex(`608060405234801561001057600080fd5b5060405161016f38038061016f8339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548163ffffffff021916908363ffffffff1602179055505060fc806100736000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80632e64cec1146037578063b9e95382146059575b600080fd5b603d608a565b604051808263ffffffff16815260200191505060405180910390f35b608860048036036020811015606d57600080fd5b81019080803563ffffffff16906020019092919050505060a3565b005b60008060009054906101000a900463ffffffff16905090565b806000806101000a81548163ffffffff021916908363ffffffff1602179055505056fea2646970667358221220f404641197f1bccb839a7e7e28ddc641f0559c5fa87cdd12dc34329023643e0b64736f6c63430007040033`)