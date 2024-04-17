# @gesia/explorr-api

Gesia Carbon Explorer Backend

## Env
```bash
# 연결 몽고 디비 URI
MONGODB_URI=

# HTTP Port 
PORT=

# 중립체인 접속 WS RPC URL (protocol, port까지 포함)
CHAIN_NEUTRALITY_WS_URL=
# 배출체인 접속 WS RPC URL (protocol, port까지 포함)
CHAIN_EMISSION_WS_URL=
# 상쇄체인 접속 WS RPC URL (protocol, port까지 포함)
CHAIN_OFFSET_WS_URL=

# 중립체인 비콘노드 API URL (protocol, port까지 포함)
CHAIN_NEUTRALITY_BEACON_API_URL= 

# 1 > 3개 체인 최신 블록을 구독하여 디비에 입력합니다. (multiprocess 구동 시 1개 노드만 돌리도록 주의하세요.)
WEB3_SUBSCRIBING=

# 1 > 3개 체인 최신 블록부터 블록 1까지 연결된 디비에 Block, Tx 정보가 없다면 전체 블록 디비에 입력합니다. (multiprocess 구동 시 1개 노드만 돌리도록 주의하세요.)
WEB3_SYNCING=
```