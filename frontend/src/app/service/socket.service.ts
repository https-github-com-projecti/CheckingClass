import { Injectable, EventEmitter  } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class SocketService {

  private socket: WebSocket;
  private listener: EventEmitter<any> = new EventEmitter();
  private API = "http://localhost:8080/";
  constructor(private httpClient: HttpClient) { }

  public newSocket(){
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
            console.log("Evant data : " + event.data);
            console.log("Evant data type : " + typeof event.data);
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

  getClientID() {
    return this.httpClient.get(this.API + "Attendance/" + "getClientId")
  }
}
