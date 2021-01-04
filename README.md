# Icarus
## Extractor for the Celo network

Icarus is a data extractor for the Celo network. It currently powers the backend for
[Anthem](https://anthem.chorus.one), providing historical balances, governance and staking data for all accounts on the Celo network.

## Requirements

- Go >= 1.15.0
- Postgres >= 10

## Setup Instructions

To build, run `make build`. 

## Configuration

Config variables that can / should be set for running Icarus are as follows:

- `CELO_URI`: URI (either unix://.. or http(s):// accepted) to Celo blockchain node. A local node connected via IPC will sync much faster.
- `DB_HOST`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `DB_PORT`
- `DB_SSLMODE`: Defaults to `require`.
- `LISTEN_ADDRESS`: Defaults to `0.0.0.0:10101`.


## Runnning Icarus 

1. Create `config.env` in your working directory. Refer sample at `config.env.sample`, or set the environment variables referenced above.

2. Run the Extractor 
    ```
    ./build/icarus
    ```

3. Run the API server
    ```
    ./build/icarus-api
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

