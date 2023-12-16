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
    1: i32 Num;
    2: string Code;
    3: DbMsg Dbmsg;
}

struct InsertRequest {
    1: i32 Mid (api.form="mid");
    2: string Msg (api.form="msg");
}

struct InsertResponse {
    1: i32 Num;
    2: string Code;
}

struct DeleteRequest {
    1: i32 Mid (api.form="mid");
}

struct DeleteResponse {
    1: i32 Num;
    2: string Code;
}

struct UpdateRequest {
    1: i32 Mid (api.form="mid");
    2: string Msg (api.form="msg");
}

struct UpdateResponse {
    1: i32 Num;
    2: string Code;
}

service QueryService {
    QueryResponse queryMsg(1: QueryRequest request) (api.get="msg/query")
    InsertResponse insertMsg(1: QueryRequest request) (api.post="msg/insert")
    DeleteResponse deleteMsg(1: QueryRequest request) (api.post="msg/delete")
    UpdateResponse updateMsg(1: QueryRequest request) (api.post="msg/update")
}
