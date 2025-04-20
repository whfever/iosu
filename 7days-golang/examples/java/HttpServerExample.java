import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpExchange;
import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.util.concurrent.Executors;

public class HttpServerExample {
    // 等同于 Go 的 Engine struct
    static class Engine {
        public void handle(HttpExchange exchange) throws IOException {
            String path = exchange.getRequestURI().getPath();
            String response = "";
            
            // 等同于 Go 的路由处理
            switch (path) {
                case "/":
                    // 对应 Go 的 indexHandler
                    response = String.format("Java URL.Path = \"%s\"\n", path);
                    break;
                case "/hello":
                    // 对应 Go 的 helloHandler
                    StringBuilder headers = new StringBuilder();
                    exchange.getRequestHeaders().forEach((key, value) -> 
                        headers.append(String.format("Header[\"%s\"] = \"%s\"\n", key, value))
                    );
                    response = headers.toString();
                    break;
                default:
                    response = String.format("404 NOT FOUND: %s\n", path);
            }
            
            exchange.sendResponseHeaders(200, response.getBytes().length);
            try (OutputStream os = exchange.getResponseBody()) {
                os.write(response.getBytes());
            }
        }
    }

    public static void main(String[] args) throws IOException {
        // 创建HTTP服务器，监听9999端口
        HttpServer server = HttpServer.create(new InetSocketAddress(9999), 0);
        Engine engine = new Engine();
        
        // 注册处理器
        server.createContext("/", engine::handle);
        
        // 使用线程池
        server.setExecutor(Executors.newFixedThreadPool(10));
        
        // 启动服务器
        server.start();
        
        System.out.println("Server is running on http://localhost:9999");
        
        // 等待中断信号
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.out.println("Shutting down server...");
            server.stop(5); // 给5秒优雅关闭时间
            System.out.println("Server stopped");
        }));
    }
}