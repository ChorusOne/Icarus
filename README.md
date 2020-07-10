# Icarus
## Extractor for the Celo network

Icarus is a data extractor for the Celo network. It currently powers the backend for
[Anthem](https://anthem.chorus.one), providing historical balances, governance and staking data for all accounts on the Celo network.

## Requirements

- Go >= 1.13.9
- Postgres >= 10

## Setup Instructions

_Better instructions coming soon_


## Runnning Icarus 

1. Create `config.env` file in the `cmd/` folder (Refer sample at `cmd/config.env.sample`)

2. Run the Extractor 
    ```
    cd cmd
    go run icarus.go
    ```

3. Run the API server
    ```
    cd cmd
    go run api.go
    ```

## Contributing / Code Layout

Contributions are welcome! For a quick overview of the code structure, check
out this tree overview.

```
├── atlas         -- Functions to extract contract related data that is not trivially accessible via contract calls
├── binding       -- Go bindings for relevant Solidity contracts deployed on Celo
│  └── mainnet    -- Bindings for Celo Mainnet
│  └── baklava    -- Bindings for Baklava Testnet
├── blockchain    -- Functions to fetch information about blockchain structs like blocks, block headers, state dumps
├── blockshot     -- Entry point for the extractor, iterates through all blocks and performs extraction tasks    
├── cmd           -- Contains executables to run Icarus and the API server
├── common        -- Celo blockchain deployment constants (networks, contract address & deployment height, ... )
├── contract      -- Abstractions over relevant Celo smart contract calls
├── misc          -- Miscelleanous folders 
│  └── solidity   -- Copy of Solidity code for smart contracts deployed on the Celo network
├── rest          -- API server code to enable Anthem to consume data extracted by Icarus
├── snapshot      -- Code to generate detailed daily snapshots for all accounts on the network 
```

