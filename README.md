
# Unibase DA SDK

**Official SDK for interacting with Unibase DA — the first high-performance decentralized data availability (DA) and storage layer for AI and DePIN applications.**

Built for developers who need scalable, secure, and cost-efficient on-chain data storage and retrieval.

---

## 🚀 What is Unibase DA?

Unibase DA provides high-throughput, low-latency decentralized storage and retrieval services, optimized for:

- AI memory (conversation history, knowledge bases)
- DePIN networks (compute, bandwidth, storage nodes)
- Cross-chain data availability for decentralized applications (DeFi, SocialFi, GameFi)

It supports seamless file uploading, downloading, and verifiable storage on BNBChain, Optimism, and more.

---

## ✨ Key Features

- **High-Throughput, Low-Latency Storage**  
  Provides real-time access to decentralized storage with over 100+ GB/s throughput and sub-100ms latency.

- **On-Chain Verification via ZK Proofs**  
  All data storage and retrieval operations are verified with zero-knowledge proofs to ensure security and transparency.

- **Cross-Chain Compatibility**  
  Supports BNBChain Testnet, Optimism Sepolia, and opBNB Testnet out-of-the-box.

- **Developer-Friendly SDK**  
  Simple CLI tools for fast integration into AI, DePIN, and Web3 applications.

---

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

---

## 🌐 Supported Networks

- BNBChain Testnet
- Optimism Sepolia
- opBNB Testnet

---

## 🤝 Contributing

We welcome contributions!

- Fork the repository
- Create a feature branch
- Write tests and documentation
- Submit a Pull Request 🚀

---

## 📞 Contact

- Website: [https://www.unibase.com](https://www.unibase.com)
- Telegram: [@unibase_ai](https://t.me/unibase_ai)
- Email: <support@unibase.com>

---
