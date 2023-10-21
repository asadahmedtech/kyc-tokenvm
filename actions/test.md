# How to run

KYC Data Includes:

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

This can be expanded to:

1. Making AllowList precompiles dynamic, as admin can whitelist address on the fly without node restart.
2. Allowing app-chains to create custom transfer rules for allowing people to transact within the same company userbase.
3. Implementing "rule of the land" and having fair logs on cross-border payments.
