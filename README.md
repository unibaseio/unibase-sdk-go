
# Unibase DA

**Unibase DA** is a high-performance, decentralized data availability (DA) and storage layer optimized for AI and DePIN workloads. Built with zk-based on-chain verification, it supports massive data throughput and EB-scale capacity while ensuring trust and transparency via Ethereum-based settlement.

---

## 🔍 Why Unibase DA?

Current DA solutions (EigenDA, Celestia, Avail, etc.) were designed for rollup transaction data, not the large-scale, real-time data workloads required by AI and DePIN. Unibase DA solves this by offering:

- 🔐 On-chain verifiability with ZK Proofs
- ⚡️ 100GB/s write throughput & 100MB/s encoding
- 🌐 Horizontal scaling to millions of devices
- 🧠 Native support for AI/DePIN memory + model data
- 📦 Modular storage APIs for files, models, blobs, and streams

---

## 📊 DA Protocol Comparison

| Feature                      | Ethereum | EigenDA | Celestia | Avail | 0G | **Unibase DA** |
|-----------------------------|----------|---------|----------|-------|----|----------------|
| **Data Verification**       | —        | Off-chain | Off-chain | Off-chain | Off-chain | ✅ **On-chain (ZK Proofs)** |
| **Data Security Model**     | Honest Majority | Honest Majority | Honest Majority | Honest Majority | Honest Majority | ✅ **Honest One (On-chain)** |
| **Data Throughput**         | 0.083 MB/s | 10 MB/s | 1.5–10 MB/s | — | 50 GB/s | ✅ **100 GB/s** |
| **Consensus**               | Casper   | None    | Tendermint | BABE + GRANDPA | Tendermint | None |
| **DAS Support**             | ❌        | ❌      | ✅        | ✅     | ✅  | ✅ |
| **Proof Type**              | Validity Proof | Validity Proof | Fraud Proof | Validity Proof | Validity Proof | ✅ **Fraud Proof** |
| **DA Settlement Layer**     | Ethereum | Ethereum | Celestia | Avail | 0G | ✅ **Ethereum** |

---

## 📦 Features

### ✅ High Performance

- **Write Throughput:** 100 GB/s+
- **Encoding Speed:** 100 MB/s
- **Low-Latency Access:** Built for real-time AI/DePIN scenarios

### ✅ High Scalability

- **Capacity:** EB+
- **Nodes:** 1M+ edge/DePIN nodes
- **Storage Pools:** Public & private options

### ✅ Programmable

- **Ownership & Access Control:** On-chain via smart contracts
- **ZK Proofs:** Enable verifiable data integrity
- **Assetization:** Built-in support for data NFTs & streams

---

## 🧠 Architecture

Unibase DA follows a multi-layer architecture to ensure speed, reliability, and transparency:

- **Smart Contract Layer:** Receives commitment, validates encode proofs, settles data
- **Storage Nodes:** Store encoded data blocks
- **ZK Verifier:** Validates submitted proofs
- **Client SDK:** Uploads files, generates proofs, interacts with contracts

![Unibase DA Verification Diagram](https://path-to-image/da-verification.png)

---

## 🚀 SDK Quickstart

```bash

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

## 🔗 Links

- 🌐 Website: [https://www.unibase.com](https://www.unibase.com)
- 📖 GitBook: [Unibase Docs](https://openos-labs.gitbook.io/unibase-docs/)
- 🧠 GitHub: [github.com/unibaseio](https://github.com/unibaseio)
- 🐦 Twitter: [@Unibase_AI](https://twitter.com/Unibase_AI)
- 📢 Telegram: [@unibase_ai](https://t.me/unibase_ai)

---

## 📜 License

MIT License
