# bridge-backend-service

This project is designed to provide bridge between different blockchain networks

1. Event Driven Design. Watch "NFT Mint" event from the source chain(compare with the block chain of orther side which called "destination") and make sure this NFT token which is going to be minted is not existed in the destination blockchain network. Then call the source bridge contract to allow/discard this mint progress, at the same time,record this NFT token's status so that it will be known as "already minted".
2. Watch "NFT Bridge" event from the source bridge contract which was deployed on the source blockchain network and call another bridge contract which was existed in the blockchain network which is going to bridge to.
3. Provide an API to query each bridge records.
4. Use Mysql as persistence database and etcd for distribute lock
