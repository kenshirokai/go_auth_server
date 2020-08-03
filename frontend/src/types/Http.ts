export type HttpPost = (url: string, body: any, callback: HttpCallback) =>  void
export type HttpCallback = {
    success?: Function 
    failure?: Function
}