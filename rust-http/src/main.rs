use actix_web::{App, HttpRequest, HttpResponse, HttpServer, Responder, web, post};
use serde::{Deserialize, Serialize};

async fn greet(req: HttpRequest) -> impl Responder {
    let name = req.match_info().get("name").unwrap_or("World");
    format!("Hello {}!", &name)
}

#[derive(Serialize, Deserialize)]
struct MyObj {
    name: String,
}

#[post("/echo")]
async fn echo(body: web::Json<MyObj>) -> impl Responder {
    let json = serde_json::to_string(
        &body.0
    ).unwrap();
    HttpResponse::Ok().body(json)
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