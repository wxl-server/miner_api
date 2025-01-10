// idl/hello.thrift
namespace go miner_api

struct HelloReq {
    1: string ID (api.query="id");
}

struct HelloResp {
    1: required i64 code;
    2: required string message;
    3: optional HelloData data;
}

struct HelloData {
    1: required string RespBody;
}

service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
}
