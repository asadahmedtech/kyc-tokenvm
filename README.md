<p align="center">
  <img width="90%" alt="tokenvm" src="assets/logo.png">
</p>
<p align="center">
  Mint, Transfer, and Trade User-Generated Tokens, All On-Chain
</p>
<p align="center">
  <a href="https://github.com/ava-labs/hypersdk/actions/workflows/tokenvm-static-analysis.yml"><img src="https://github.com/ava-labs/hypersdk/actions/workflows/tokenvm-static-analysis.yml/badge.svg" /></a>
  <a href="https://github.com/ava-labs/hypersdk/actions/workflows/tokenvm-unit-tests.yml"><img src="https://github.com/ava-labs/hypersdk/actions/workflows/tokenvm-unit-tests.yml/badge.svg" /></a>
  <a href="https://github.com/ava-labs/hypersdk/actions/workflows/tokenvm-sync-tests.yml"><img src="https://github.com/ava-labs/hypersdk/actions/workflows/tokenvm-sync-tests.yml/badge.svg" /></a>
  <a href="https://github.com/ava-labs/hypersdk/actions/workflows/tokenvm-load-tests.yml"><img src="https://github.com/ava-labs/hypersdk/actions/workflows/tokenvm-load-tests.yml/badge.svg" /></a>
</p>

---

Avalanche created the [`tokenvm`](./examples/tokenvm) to showcase how to use the
`hypersdk` in an application most readers are already familiar with, token minting
and token trading. The `tokenvm` lets anyone create any asset, mint more of
their asset, modify the metadata of their asset (if they reveal some info), and
burn their asset. Additionally, there is an embedded on-chain exchange that
allows anyone to create orders and fill (partial) orders of anyone else. To
make this example easy to play with, the `tokenvm` also bundles a powerful CLI
tool and serves RPC requests for trades out of an in-memory order book it
maintains by syncing blocks.

If you are interested in the intersection of exchanges and blockchains, it is
definitely worth a read (the logic for filling orders is < 100 lines of code!).

This project extends TokenVM to support KYC Compliance over the assets and making life easier with aliases. 

The project contains KYC Data which Includes:

1. Country of Residence / Issuance
2. Authority of the individual

Authority can be of:

1. Government: Can issue KYC to adddress
2. Company: Can create and mint Tokens
3. Individual: Can transact Tokens with addresses of Same Country

Demo talks about:

1. Government address issuing KYC to Individual
2. Individual trying to Create and Mint Assets but fails
3. Government address issuing KYC to Company
4. Company tries and succeed to Create and Mint Assets

Commands:

./build/token-cli action create-asset
./build/token-cli action kyc
./build/token-cli action mint-asset

assetId:
chainId:

Project Implements:

1. Transfer checks on from and to addresses.
2. Mint and Burn checks for companies.
3. Cross-Border checks for transfers
4. Aliases based address mapping.

This can be expanded to:

1. Making AllowList precompiles dynamic, as admin can whitelist address on the fly without node restart.
2. Allowing app-chains to create custom transfer rules for allowing people to transact within the same company userbase.
3. Implementing "rule of the land" and having fair logs on cross-border payments.
