rm -R crypto-config/*

../bin/cryptogen generate --config=crypto-config.yaml

rm config/*

../bin/configtxgen -profile VLMOrgOrdererGenesis -outputBlock ./config/genesis.block

../bin/configtxgen -profile VLMOrgChannel -outputCreateChannelTx ./config/vlmchannel.tx -channelID vlmchannel
