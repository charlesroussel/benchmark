use actix_web::{error, App, Error, HttpRequest, HttpResponse, HttpServer, Responder, web, post};
use futures::StreamExt;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct MyObj {
    name: String,
}

const MAX_SIZE: usize = 262_144; // max payload size is 256k

#[post("/echo")]
async fn echo(mut payload: web::Payload) -> impl Responder {
    let mut body = web::BytesMut::new();
    while let Some(chunk) = payload.next().await {
        let chunk = chunk?;
        // limit max size of in-memory payload
        if (body.len() + chunk.len()) > MAX_SIZE {
            return Err(error::ErrorBadRequest("overflow"));
        }
        body.extend_from_slice(&chunk);
    }
    let obj = serde_json::from_slice::<MyObj>(&body)?;

    Ok(HttpResponse::Ok().json(obj))
}


#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(echo)
    })
    .bind(("0.0.0.0", 8080))?
    .run()
    .await
}