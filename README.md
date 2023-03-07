# zk rollup example

## Problem
Let's say we want to solve a math problem: `y = x^4 + x^3 + x^2 + x`, I can prove you I know the answer without telling you the answer.

This example use `gnark` to help understand how `zk SNARK` works, for the equation above, we know two answers are `(1,4)`, `(2,30)`, 
we use `(1,4)` in this example.  

## Procedure
1. Write circuit, compile circuit
2. Generate `PK(proving key)` and `VK(verifying key)`, export verifier contract(We're going to verify on blockchain),
this contract contains `VK`.
3. Generate witness
4. Generate proof by PK and witness 
5. Deploy verifier contract on chain
6. Verify on chain by proof

## Steps
1. Build example: 
	```
	make build
	```
2. Generate `PK`, `VK`, `verifier contract`, `proof`: 
	```
	build/zk-example
	```
3. Deploy `verifier contract` and verify on chain: 
	```
	cd solidity && npm install
	npx hardhat test
	```
 	Output:
	```
	Compiled 1 Solidity file successfully
	
	
	  Verifier contract
	input:  [ 4 ]
	Verify result:  true
		✔ Verify should success (1486ms)
	input:  [ 5 ]
	Verify result:  false
		✔ Verify should fail (274ms)
	
	
	  2 passing (2s)
	```
You can use other `x,y` to test this micro zk system, like `(2,30)`, `(3,120)` by running:
```
build/zk-example -x 2 -y 30
cd solidity && npx hardhat test
```
