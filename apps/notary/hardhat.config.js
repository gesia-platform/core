module.exports = {
  defaultNetwork: "localhost",
  networks: {
    hardhat: {},
    localhost: {
      url: "http://127.0.0.1:8545",
      accounts: {
        mnemonic:
          "answer debate indoor acquire small lady impose mandate guitar airport win deputy",
        path: "m/44'/60'/0'/0",
        initialIndex: 0,
        count: 5,
        passphrase: "",
      },
    },
  },
  solidity: {
    version: "0.8.23",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  paths: {
    sources: "./contracts",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts",
  },
  mocha: {
    timeout: 40000,
  },
};
