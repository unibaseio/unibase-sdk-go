# Unibase DA

**Unibase DA** is a high-performance, decentralized data availability (DA) and storage network, purpose-built for AI agents and on-chain applications.  
It offers **100GB/s throughput**, **zk-proof-based on-chain verification**, and **EB-level scalability** through decentralized nodes.

---

## ✨ Key Features

- **High Performance**
  - 100 GB/s write throughput
  - 100 MB/s encoding speed
  - On-chain data verification powered by Zero-Knowledge Proofs (ZKPs)

- **Massive Scalability**
  - EB+ capacity with decentralized storage pools
  - 1M+ Depin devices contributing storage, compute, and bandwidth
  - Private enterprise storage pool support

- **On-Chain Programmability**
  - Smart contract-based access control
  - Native support for data assetization and customized ownership

---

## 🛠 Architecture Overview

- **Client Layer**: Submits data commitment and metadata to the chain.
- **Store Nodes**: Encode, store, and generate zk-proofs for stored data.
- **Smart Contracts**: Manage identity registration, access permissions, proof validation, and data availability settlement.
- **On-Chain Verification**: Continuous dual proofs ensure data persistence and reliability.

---

## 🔗 Data Upload / Download SDK

SDK repository: [https://github.com/MOSSV2/dimo-sdk-go](https://github.com/MOSSV2/dimo-sdk-go)

Supported Testnets:

- BNBChain Testnet
- Optimism Sepolia
- opBNB Testnet

Quick Start:

```bash
# Upload file or directory

## 📦 Installation

```bash
git clone https://github.com/unibaseio/unibase-sdk-go.git
cd unibase-sdk-go
go build
```

---

## 📚 Quick Usage

### Upload a file/directory

```bash
export CHAIN_TYPE=bnb-testnet
cd example/upload
go build
./upload --model=false --sk=<your_secret_key> --path=<your_local_path>
```

### Download a file/directory

```bash
export CHAIN_TYPE=bnb-testnet
cd example/download
go build
./download --model=false --sk=<your_secret_key> --name=<file_name> --path=<your_save_path>
```

### Public Hub (Optional)

- Download:

  ```bash
  wget http://54.151.130.2:8080/api/download?name=<your_file_name>&owner=<your_owner_address> -O <save_as_name>
  ```

- Upload:

  ```bash
  curl -X POST http://54.151.130.2:8080/api/upload -d '{"id":"test1","owner":"0xabcd","message":"sample message"}'
  ```

Public Hub API available for lightweight storage and retrieval.

---

## 📚 Documentation

- Full GitBook: Coming Soon
- Storage Network Overview
- On-Chain Proof Workflow

---

## 📞 Contact

- Website: [https://www.unibase.com](https://www.unibase.com)
- Support: [support@unibase.com](mailto:support@unibase.com)
- Telegram: [@unibase_ai](https://t.me/unibase_ai)

---

✅ **Unibase DA** redefines high-availability storage for decentralized AI and next-generation Web3 infrastructures.
