FROM ghcr.io/rust-lang/rust:nightly-slim
COPY . /rust-http
WORKDIR /rust-http
RUN cargo build --release
EXPOSE 8080:8080
ENV WORKERS 1
CMD [ "./target/release/rust-http" ]
