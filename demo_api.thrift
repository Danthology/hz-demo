namespace go hz.demo

struct QueryRequest {
    1: i32 Num (api.query="num");
}

struct QueryResponse {
    1: i32 Num;
    2: string Code;
}

service QueryService {
    QueryResponse queryMsg(1: QueryRequest request) (api.get="msg/query")
}
