namespace go hz.demo

enum Code {
  Success = 0;
  ParamInvalid = 1;
  DBErr = 2;
}

struct DbMsg {
    1: i32 Mid;
    2: string Msg;
}

struct QueryRequest {
    1: i32 Num (api.query="num");
}

struct QueryResponse {
    1: Code Code;
    2: string Des;
    3: DbMsg Dbmsg;
}

struct InsertRequest {
    1: i32 Mid (api.form="mid");
    2: string Msg (api.form="msg");
}

struct InsertResponse {
    1: Code Code;
    2: string Des;
}

struct DeleteRequest {
    1: i32 Mid (api.form="mid");
}

struct DeleteResponse {
    1: Code Code;
    2: string Des;
}

struct UpdateRequest {
    1: i32 Mid (api.form="mid");
    2: string Msg (api.form="msg");
}

struct UpdateResponse {
    1: Code Code;
    2: string Des;
}

service QueryService {
    QueryResponse queryMsg(1: QueryRequest request) (api.get="msg/query")
    InsertResponse insertMsg(1: InsertRequest request) (api.post="msg/insert")
    DeleteResponse deleteMsg(1: DeleteRequest request) (api.post="msg/delete")
    UpdateResponse updateMsg(1: UpdateRequest request) (api.post="msg/update")
}

struct AddRequest {
    1: string K (api.form="k");
    2: string V (api.form="v");
}

struct AddResponse {
    1: Code Code;
    2: string Des;
}

struct GetRequest {
    1: string K (api.form="k");
}

struct GetResponse {
    1: Code Code;
    2: string Des;
    3: string V;
}

service RedisService {
    AddResponse addMsg(1: AddRequest request) (api.post="redis/add")
    GetResponse getMsg(1: GetRequest request) (api.post="redis/get")
}
