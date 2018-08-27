# go-contract #

## Prepare ##
Clone go-ethereum
```
mkdir $GOPATH/src/github.com/ethereum

git clone git@github.com:ethereum/go-ethereum.git
```

Install solc
```
npm install -g solc solc-cli --save-dev
```

## Compile abigen ##
```
cd $GOPATH/src/github.com/ethereum/go-ethereum/cmd/abigen/

go install
```

## Generate go version contract ##
Instead of operating abigen by [this tutorial](https://www.cnblogs.com/baizx/p/7469125.html) which never succeed for me,  
[another tutorial](https://hk.saowen.com/a/16e09b1d3e7f3099ac7b7ce9f891e96d8588c6d563c6dab8e344dbdff397cc73) is recommended.  

First, use solcjs(after npm install -g solc) to generate abi and bin.
```
solcjs sol/FixedSupplyToken.sol -o tmp/ctrabi --abi 

solcjs sol/FixedSupplyToken.sol -o tmp/ctrbin --bin
```

Then use abigen with abi and bin above to get contract wrapper in GO.  
Be aware of that --pkg not means which directory the contract wrapper file will be placed,  
it menas whcih package the contract wrapper will belong to. 
```
abigen \
--abi tmp/ctrabi/sol_FixedSupplyToken_sol_FixedSupplyToken.abi \
--bin tmp/ctrbin/sol_FixedSupplyToken_sol_FixedSupplyToken.bin \
--pkg wrapper --out wrapper/fstoken.go
```

## Ref ##
[https://www.cnblogs.com/baizx/p/7469125.html](https://www.cnblogs.com/baizx/p/7469125.html)  
[https://hackmd.io/V_-nYbXxR0WGpon8eylo5g?both](https://hackmd.io/V_-nYbXxR0WGpon8eylo5g?both)  
[https://hk.saowen.com/a/16e09b1d3e7f3099ac7b7ce9f891e96d8588c6d563c6dab8e344dbdff397cc73](https://hk.saowen.com/a/16e09b1d3e7f3099ac7b7ce9f891e96d8588c6d563c6dab8e344dbdff397cc73)  


## Note ##
abigen use the solc to compile *.sol.  
User may want to use [solc](https://medium.com/@ksin751119/ethereum-dapp%E5%88%9D%E5%BF%83%E8%80%85%E4%B9%8B%E8%B7%AF-7-web3-eth-compile-solidity-%E6%9B%BF%E4%BB%A3%E6%96%B9%E6%A1%88-a020a6763fbd) directly when [error occurs](https://blog.studygolang.com/topics/5215) with abigen.

```
solc FixedSupplyToken.sol > FixedSupplyToken.abi
```