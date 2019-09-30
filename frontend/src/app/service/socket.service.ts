import { Injectable, EventEmitter  } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SocketService {

  private socket: WebSocket;
  private listener: EventEmitter<any> = new EventEmitter();
  constructor() { 
    this.socket = new WebSocket("ws://localhost:8080/websocket/ws");
        this.socket.onopen = event => {
          console.log("this.socket.onopen : " + event);
            this.listener.emit({"type": "open", "data": event});
        }
        this.socket.onclose = event => {
          console.log("this.socket.onclose : " + event);
            this.listener.emit({"type": "close", "data": event});
        }
        this.socket.onmessage = event => {
          console.log("this.socket.onmessage : " + event);
            this.listener.emit({"type": "message", "data": JSON.parse(event.data)});
        }
  }
  public send(data: string) {
    console.log("Data from func public send : " + data);
    this.socket.send(data);
  }

  public close() {
    this.socket.close();
  }

  public getEventListener() {
    return this.listener;
  }
}
