docker stack deploy -c ca.yaml ca_org2

docker stack deploy -c orderer4.yaml orderer4_org2

docker stack deploy -c orderer5.yaml orderer5_org2

docker stack deploy -c peer0.yaml peer0_org2

docker stack deploy -c peer1.yaml peer1_org2
