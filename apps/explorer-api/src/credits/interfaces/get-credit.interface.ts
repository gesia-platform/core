interface TransferEvent {
    transactionHash: string;
    methodName: string;
    blockNumber: string;
    creationTime: number;
    from: string;
    to: string;
    amount: bigint;
    id: string;
}

interface AggregatedData {
    amount: bigint;
    transactionHash: string;
    methodName: string;
    blockNumber: string;
    creationTime: number;
    from: string;
}
