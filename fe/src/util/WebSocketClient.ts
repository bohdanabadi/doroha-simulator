export class WebSocketClient {
    private url: string;
    private maxRetries: number;
    private retryCount: number;
    private socket: WebSocket | null;
    private listeners: ((data: any) => void)[];

    constructor(url: string, maxRetries: number = 3) {
        this.url = url;
        this.maxRetries = maxRetries;
        this.retryCount = 0;
        this.socket = null;
        this.listeners = [];
    }

    connect(): void {
        this.socket = new WebSocket(this.url);

        this.socket.onopen = () => {
            console.log("WebSocket connected");
            this.retryCount = 0; // Reset retry count on successful connection
        };

        this.socket.onerror = (error: Event) => {
            console.error("WebSocket error:", error);
            this.socket?.close();
        };

        this.socket.onclose = () => {
            if (this.retryCount < this.maxRetries) {
                setTimeout(() => {
                    console.log("Attempting to reconnect WebSocket...");
                    this.retryCount++;
                    this.connect();
                }, 2000 * this.retryCount); // Exponential back-off
            }
        };

        this.socket.onmessage = (event: MessageEvent) => {
            this.listeners.forEach(listener => listener(event.data));
        };
    }

    addListener(listener: (data: any) => void): void {
        this.listeners.push(listener);
    }

    removeListener(listener: (data: any) => void): void {
        this.listeners = this.listeners.filter(l => l !== listener);
    }

    sendMessage(message: string): void {
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
            this.socket.send(message);
        } else {
            console.error("WebSocket is not connected.");
        }
    }

    close(): void {
        this.socket?.close();
    }
}