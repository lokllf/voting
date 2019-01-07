# identity
docker cp sawtooth-validator-default:/root/.sawtooth/keys/my_key.priv .
docker cp my_key.priv sawtooth-shell-default:.

docker exec -it sawtooth-shell-default bash
sawset proposal create --key my_key.priv sawtooth.identity.allowed_keys=$(cat ~/.sawtooth/keys/root.pub) --url http://rest-api:8008
sawtooth settings list --url http://rest-api:8008

sawtooth identity policy create policy_1 "PERMIT_KEY *" --url http://rest-api:8008
sawtooth identity policy list --url http://rest-api:8008