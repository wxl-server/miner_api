// idl/hello.thrift
namespace go miner_api

struct HelloReq {
    1: string ID (api.query="id"); // 添加 api 注解为方便进行参数绑定
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
}
