package bindings

// wget -nc "https://github.com/ethereum/solidity/releases/download/v0.8.20/solc-static-linux"
// chmod +x ./solc-static-linux

//go:generate ./solc-static-linux --overwrite --bin --abi -o ./ ../contracts/ERC5564Announcer.sol
//go:generate abigen --pkg bindings --type ERC5564Announcer --abi ERC5564Announcer.abi --bin ERC5564Announcer.bin --out ERC5564Announcer.go
