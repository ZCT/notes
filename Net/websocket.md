## WebSocket



pingpong是WebSocket的心跳机制



melody中现在pingpong超时会自动关闭连接



conn.SetReadDeadLine ，超过这个时间点之后，后续的read会返回 io timeout的错误

