# AtomicSwaps: Decentralized NFT Marketplace with On-Chain Order Book

A Go-based implementation of a decentralized NFT marketplace leveraging on-chain order books, atomic swaps via Hashed Time Lock Contracts (HTLCs), and gas-optimized ERC-721 metadata handling.

## Detailed Description

AtomicSwaps aims to revolutionize the NFT trading experience by providing a truly decentralized and trustless platform. Unlike centralized NFT marketplaces that act as intermediaries, AtomicSwaps empowers users to directly trade NFTs peer-to-peer using on-chain order books and Hashed Time Lock Contracts (HTLCs). This eliminates the need for a central authority to hold funds or NFTs, minimizing counterparty risk and enhancing security.

The core of AtomicSwaps lies in its on-chain order book, which allows users to create and execute limit orders directly on the blockchain. This facilitates transparent price discovery and efficient order matching. The use of HTLCs ensures that trades are atomic  either both parties receive what they agreed upon, or the entire transaction is reverted, guaranteeing that no one loses their assets. This is especially crucial in decentralized environments where trust is not assumed.

Furthermore, AtomicSwaps incorporates gas-optimized ERC-721 metadata handling techniques. By minimizing the amount of data stored on-chain and implementing efficient data structures, we significantly reduce transaction costs, making the platform more accessible and affordable for users. The platform is designed to be modular and extensible, allowing for future integrations and enhancements. Our goal is to provide a robust and scalable foundation for the next generation of decentralized NFT marketplaces.

## Key Features

*   **On-Chain Order Book:** Implemented using a Merkle tree-based data structure for efficient storage and retrieval of orders. Order matching logic is executed directly on the Ethereum Virtual Machine (EVM), ensuring transparency and immutability. Order creation involves crafting a signed message containing order details (NFT contract address, token ID, price, expiration) which is then broadcast to the smart contract.
*   **Atomic Swaps via HTLCs:** Uses a standardized HTLC implementation compatible with ERC-721 tokens. The swap process involves two transactions: first, the NFT seller locks the NFT in the HTLC with a secret hash and a timelock. Second, the NFT buyer reveals the secret, claiming the NFT and unlocking the funds for the seller. If the buyer fails to reveal the secret within the timelock, the seller can reclaim the NFT.
*   **Gas-Optimized ERC-721 Metadata Handling:** Employs techniques like URI storage and off-chain metadata caching to minimize on-chain data storage costs. This includes storing metadata on decentralized storage solutions like IPFS and referencing it via a single URI stored on-chain. Contract design focuses on minimal storage reads and writes.
*   **Decentralized Governance (Future):** Planned integration with a decentralized autonomous organization (DAO) for community governance and platform development. This will empower users to participate in decision-making processes regarding future upgrades and feature implementations.
*   **Go-based Smart Contract Interaction:** The backend is built using Go, allowing for efficient interaction with Ethereum smart contracts via the `go-ethereum` library. This simplifies development and deployment of the platform.
*   **Secure and Audited Smart Contracts:** All smart contracts are rigorously audited to ensure security and prevent vulnerabilities. Formal verification techniques are employed to mathematically prove the correctness of critical contract functions.

## Technology Stack

*   **Go:** The primary programming language used for backend development and smart contract interaction. Go's concurrency model and efficient memory management make it suitable for building high-performance decentralized applications.
*   **Ethereum:** The blockchain platform where the smart contracts are deployed and executed. Ethereum provides the infrastructure for decentralized applications and secure transactions.
*   **Solidity:** The programming language used to write the smart contracts that govern the NFT marketplace logic. Solidity is specifically designed for developing smart contracts on the Ethereum blockchain.
*   **IPFS:** A decentralized storage network used for storing NFT metadata. IPFS ensures that NFT metadata is immutable and accessible, even if the original provider goes offline.
*   **go-ethereum:** A Go library for interacting with the Ethereum blockchain. It provides functionalities for connecting to Ethereum nodes, deploying smart contracts, and executing transactions.
*   **Merkle Trees:** Used for efficient storage and verification of order data in the on-chain order book.

## Installation

1.  **Install Go:** If you don't have Go installed, download and install it from [https://golang.org/dl/](https://golang.org/dl/).
2.  **Clone the Repository:**
    git clone https://github.com/ezozu/AtomicSwaps.git
    cd AtomicSwaps
3.  **Install Dependencies:**
    go mod download
4.  **Install Ganache (Optional):** For local testing, install Ganache, a personal Ethereum blockchain. You can download it from [https://www.trufflesuite.com/ganache](https://www.trufflesuite.com/ganache).
5.  **Install Truffle (Optional):** For deploying and managing smart contracts, install Truffle:
    npm install -g truffle

## Configuration

1.  **Environment Variables:** Create a `.env` file in the root directory and define the following environment variables:

    *   `ETHEREUM_NODE_URL`: The URL of your Ethereum node (e.g., `http://localhost:8545` for Ganache or a provider like Infura).
    *   `PRIVATE_KEY`: Your Ethereum private key for deploying and interacting with contracts. Ensure this key is stored securely.
    *   `CONTRACT_ADDRESS`: (Optional) The address of the deployed AtomicSwaps smart contract. If you are deploying the contract for the first time, leave this empty.
    *   `IPFS_GATEWAY`: The URL of your IPFS gateway (e.g., `https://ipfs.io/ipfs/`).

    Example `.env` file:
    ETHEREUM_NODE_URL=http://localhost:8545
    PRIVATE_KEY=0xYourPrivateKeyHere
    CONTRACT_ADDRESS=
    IPFS_GATEWAY=https://ipfs.io/ipfs/

2.  **Smart Contract Configuration:**
    * If the `CONTRACT_ADDRESS` variable is empty, you need to deploy the smart contracts. Navigate to the `contracts` directory and deploy using Truffle:
    truffle migrate --network development

## Usage

1.  **Deploy Smart Contracts (if not already deployed):**
    Navigate to the `contracts` directory and run:
    truffle migrate --network development
    This will deploy the smart contracts to your local Ganache network. Note the contract address.
2.  **Run the Go application:**
    Navigate back to the root directory and run:
    go run main.go
    This will start the AtomicSwaps backend application.
3.  **API Documentation (Example - Extend as necessary):**

    The application exposes the following API endpoints:

    *   `POST /orders`: Create a new order. Requires a JSON payload containing order details (NFT contract address, token ID, price, expiration).
    *   `GET /orders`: Retrieve a list of all active orders.
    *   `GET /orders/{orderID}`: Retrieve a specific order by its ID.
    *   `POST /swap`: Execute an atomic swap. Requires a JSON payload containing the order ID and the secret.

    Example order creation request (POST /orders):
    {
        "nftContractAddress": "0xYourNFTContractAddress",
        "tokenId": 1,
        "price": "1000000000000000000", // 1 ETH in Wei
        "expiration": 1678886400 // Unix timestamp
    }

## Contributing

We welcome contributions to AtomicSwaps! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Write clear and concise code with comprehensive tests.
4.  Submit a pull request with a detailed description of your changes.
5.  Adhere to the existing code style and conventions.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/ezozu/AtomicSwaps/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the Ethereum community and the open-source developers who have contributed to the technologies used in this project.