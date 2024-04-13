```toml
# config.toml

# Gesiad Port
port = 

[chain_tree]
# ethereum hex address
address =
# ethereum hex private key without '0x' prefix
private_key = 
# bls hex secret key without '0x' prefix
bls_secret_key = 

[chain_tree.root]
# root chain id
chain_id = 
# root chain rpc url
rpc_url = 
# root chain ws url
ws_url = 
# poa network account contract ethereum hex address
network_account_address = 
# main app permission contract ethereum hex address
app_permission_address = 


[chain_tree.host]
# host chain id
chain_id = 
# host chain rpc url
rpc_url = 
# root chain ws url
ws_url = 
# main notary public contract ethereum hex address
notary_public_address = 

[keychain]
# {hostname}:{port}
host =
```