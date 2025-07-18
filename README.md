# AtomicSwaps: Secure and Trustless Cryptocurrency Exchange

A decentralized solution for facilitating atomic swaps between different blockchain networks, enabling secure and trustless cryptocurrency exchange without intermediaries.

AtomicSwaps provides a framework for performing cross-chain atomic swaps in Go, offering a practical implementation of the Hashed Timelock Contract (HTLC) protocol. This repository aims to simplify the complexities involved in coordinating transactions across disparate blockchain environments, allowing users to exchange cryptocurrencies directly with each other, reducing reliance on centralized exchanges and mitigating associated risks such as custody concerns and single points of failure. This implementation focuses on modularity and extensibility, allowing developers to easily integrate it with existing cryptocurrency wallets, exchanges, or blockchain applications. The core logic handles the creation and management of HTLCs, the generation of cryptographic secrets, and the coordination of on-chain transactions to ensure the successful execution of atomic swaps.

This project addresses a crucial need in the decentralized finance (DeFi) space: interoperability. By enabling seamless atomic swaps, AtomicSwaps contributes to a more interconnected and efficient cryptocurrency ecosystem. This allows users to diversify their holdings, arbitrage price differences across different chains, and access new opportunities without the need for trusted third parties. The implementation also emphasizes security, employing robust cryptographic techniques and rigorous validation to prevent fraud and ensure the integrity of the swap process. Furthermore, the code is designed with clarity and maintainability in mind, facilitating easier auditing and contributions from the open-source community.

The long-term goal is to provide a production-ready solution for atomic swaps that can be easily deployed and integrated into various blockchain applications. This involves continuous improvement of the code, thorough testing, and active engagement with the community to gather feedback and address potential issues. This implementation also provides examples to connect and do swaps between two instances of regtest Bitcoin.

## Key Features

*   **Hashed Timelock Contract (HTLC) Implementation:** The core functionality revolves around creating and managing HTLCs on different blockchain networks. This involves generating cryptographic hashes and secrets, constructing and signing transactions that lock funds within the HTLC, and verifying the correctness of the secret provided by the counterparty. Specifically, the implementation supports the standard HTLC construction with SHA256 hashlocks and absolute timelocks.

*   **Cross-Chain Transaction Coordination:** The system facilitates the coordination of transactions across different blockchain networks. This includes managing the timing of transactions, monitoring blockchain states, and ensuring that the correct transactions are executed on each chain to complete the swap. This also handles the case where one party does not broadcast the transaction to claim the funds in the allotted time.

*   **Modular and Extensible Design:** The codebase is designed with modularity in mind, allowing developers to easily integrate it with different blockchain networks and customize the swap process to fit their specific needs. This includes abstracting the blockchain-specific details into separate modules and providing a clean API for interacting with the core swap logic. This allows for flexibility in the choice of blockchains and customization for individual requirements.

*   **Secure Secret Exchange:** The implementation handles the secure exchange of the cryptographic secret between the parties involved in the swap. This involves encrypting the secret using the counterparty's public key and ensuring that it is only revealed to the correct party at the appropriate time. The secret is protected using standard encryption algorithms to ensure confidentiality.

*   **Error Handling and Recovery:** The system includes robust error handling and recovery mechanisms to gracefully handle potential failures during the swap process. This includes detecting invalid transactions, handling network outages, and recovering funds in case of a failed swap. Extensive logging and error reporting are incorporated to facilitate debugging and troubleshooting.

*   **Regtest Support for Testing:** The provided examples are intended to be used with regtest mode, a local testing environment for Bitcoin. This enables developers to thoroughly test and debug their atomic swap implementations without risking real funds on the main network.

## Technology Stack

*   **Go (Golang):** The primary programming language used for developing the AtomicSwaps framework. Go was chosen for its performance, concurrency support, and suitability for blockchain applications.
*   **Bitcoin RPC Client:** Used to interact with Bitcoin Core nodes. Specifically, this implementation uses a standard RPC client to create transactions, sign transactions, and get blockchain data.
*   **SHA256 Hashing Algorithm:** Used for generating cryptographic hashes for locking funds within the HTLCs. The SHA256 algorithm provides the necessary security and collision resistance.
*   **Base58 Encoding:** Used for encoding and decoding Bitcoin addresses and other data. Base58 encoding is commonly used in Bitcoin and other cryptocurrencies to represent data in a human-readable format.

## Installation

1.  **Install Go:** Ensure that you have Go version 1.18 or later installed on your system. You can download the latest version of Go from the official Go website: [https://golang.org/dl/](https://golang.org/dl/)
2.  **Clone the Repository:** Clone the AtomicSwaps repository from GitHub:

    git clone https://github.com/ezozu/AtomicSwaps.git

3.  **Navigate to the Project Directory:** Change your current directory to the cloned repository:

    cd AtomicSwaps

4.  **Install Dependencies:** Download all go modules:

    go mod download

5.  **Building:** Build the project using the go build command.

    go build

## Configuration

The application relies on the following environment variables for configuration:

*   `NODE_ONE_RPC_URL`: The RPC URL of the first regtest Bitcoin node. Example: `http://user:password@localhost:18443`
*   `NODE_TWO_RPC_URL`: The RPC URL of the second regtest Bitcoin node. Example: `http://user:password@localhost:18443`

These environment variables are crucial for specifying the addresses of the Bitcoin nodes participating in the atomic swap. Ensure that these variables are properly set before running the application.

## Usage

The project provides examples of basic functions in the test folder. These are intended to be run against regtest Bitcoin nodes.
First make sure to have two Bitcoin nodes running with regtest enabled. Then use the following commands in the Bitcoin CLI to create necessary addresses and fund the nodes:

1.  `bitcoin-cli -regtest createwallet "testwallet"`
2.  `bitcoin-cli -regtest getnewaddress` (Store this address for later use)
3.  `bitcoin-cli -regtest generatetoaddress 101 <your_address>` (This mines 101 blocks to the address generated above, funding the wallet and allowing funds to mature)
4.  `bitcoin-cli -regtest sendtoaddress <another_address> 10` (Send some funds to a new address in the wallet to diversify)
5.  Repeat these steps for your second Bitcoin node.

To execute a swap between these two nodes, run the compiled binary after setting the `NODE_ONE_RPC_URL` and `NODE_TWO_RPC_URL` environment variables.

The code provides functions to create the HTLC on each chain, exchange secrets, and redeem the HTLC to complete the swap.

## Contributing

We welcome contributions from the community! Please follow these guidelines when contributing:

*   Fork the repository and create a new branch for your changes.
*   Follow the existing code style and conventions.
*   Write clear and concise commit messages.
*   Submit a pull request with a detailed description of your changes.
*   Ensure that your code is thoroughly tested.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/ezozu/AtomicSwaps/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the open-source community for their contributions to the development of blockchain technologies and the various libraries and tools that made this project possible.