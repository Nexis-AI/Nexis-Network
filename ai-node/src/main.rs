use tokio::net::TcpListener;
use tokio::io::{AsyncReadExt, AsyncWriteExt};

#[tokio::main]
async fn main() -> std::io::Result<()> {
    // Simple placeholder service for AI compute node
    let listener = TcpListener::bind("127.0.0.1:8081").await?;
    println!("AI node listening on 127.0.0.1:8081");

    loop {
        let (mut socket, _) = listener.accept().await?;
        tokio::spawn(async move {
            let mut buf = [0u8; 1024];
            match socket.read(&mut buf).await {
                Ok(n) if n == 0 => return,
                Ok(n) => {
                    // Echo back for now
                    let _ = socket.write_all(&buf[0..n]).await;
                }
                Err(_) => return,
            }
        });
    }
}
