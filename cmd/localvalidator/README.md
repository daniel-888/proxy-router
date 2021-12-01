# Titan Validator STAGE 1

## The purpose of the validator is to ensure that a block hash has been correctly calculated

## The validator will NOT do the following
### 1 construct a unique block header
### 2 select which block it will analyze
### 3 route the block to the appropriate mining pool


## The validator has the following variables
### 1. BlockHeader - A custom Go type which contains the block header parameters as described in the stratum protocol
### 2. beginingTime - The time that the contract manager began allocating hashrate
### 3. blocksMined - the number of blocks which have actually been mined
### 4. contractHashRate - the contractually guarenteed hashrate of the contract
### 5. contractDuration - the length at which the contract will operate
