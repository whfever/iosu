import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpExchange;
import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.util.concurrent.Executors;

public class HttpServerExample {
    // 等同于 Go 的 Engine struct
    static class Engine implements HttpHandler {
        @Override
        public void handle(HttpExchange exchange) throws IOException {
            String path = exchange.getRequestURI().getPath();
            String response = "";
            
            // 等同于 Go 的路由处理
            switch (path) {
                case "/":
                    response = String.format("URL.Path = \"%s\"\n", path);
                    break;
                case "/hello":
                    StringBuilder headers = new StringBuilder();
                    exchange.getRequestHeaders().forEach((key, value) -> 
                        headers.append(String.format("Header[\"%s\"] = \"%s\"\n", key, value))
                    );
                    response = headers.toString();
                    break;
                default:
                    response = String.format("404 NOT FOUND: %s\n", path);
            }
            
            byte[] responseBytes = response.getBytes();
            exchange.sendResponseHeaders(200, responseBytes.length);
            try (OutputStream os = exchange.getResponseBody()) {
                os.write(responseBytes);
            }
        }
    }

    public static void main(String[] args) throws IOException {
        HttpServer server = HttpServer.create(new InetSocketAddress(9999), 0);
        server.createContext("/", new Engine());
        server.setExecutor(Executors.newFixedThreadPool(10));
        server.start();
        
        System.out.println("Server is running on http://localhost:9999");
        
        // 优雅关闭
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.out.println("Shutting down server...");
            server.stop(5);
            System.out.println("Server stopped");
        }));
    }
}