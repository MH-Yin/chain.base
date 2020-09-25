alias cleos="docker exec -it nodeos cleos --url http://127.0.0.1:8888 --wallet-url http://keosd:8901"

# 解锁wallet
cleos wallet unlock --password xxx

# 创建测试账号
cleos create account eosio eosio.hello xx xx

## 编译合约
docker exec -it nodeos bash -c "cd contracts/hello && eosio-cpp hello.cpp -o hello.wasm"

## 部署合约
cleos set contract eosio.hello contracts/hello -p eosio.hello@active

## 调用 合约
cleos push action eosio.hello hi '["world"]' -p eosio.hello@active