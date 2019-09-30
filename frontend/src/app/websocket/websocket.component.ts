import { Component, OnInit } from '@angular/core';
import { SocketService } from "../service/socket.service";

@Component({
  selector: 'app-websocket',
  templateUrl: './websocket.component.html',
  styleUrls: ['./websocket.component.css']
})
export class WebsocketComponent implements OnInit {
  public messages: Array<any>;
  public chatBox: string;

  constructor(private socket: SocketService) { 
    this.messages = [];
    this.chatBox = "";
  }

  ngOnInit() {
    this.socket.getEventListener().subscribe(event => {
      console.log("event type : " + event.type);
      console.log("event data : " + event.data.content);
      console.log("event data sender : " + event.data.sender);
      if(event.type == "message") {
        console.log("event message : " + event);
          let data = event.data.content;
          if(event.data.sender) {
              data = event.data.sender + ": " + data;
              console.log("Data sender : " + data);
          }
          console.log("Data : " + data);
          this.messages.push(data);
          console.log("this.messages.push(data) : " + this.messages);
      }
      if(event.type == "close") {
        console.log("event close : " + event);
          this.messages.push("/The socket connection has been closed");
      }
      if(event.type == "open") {
        console.log("event open : " + event);
          this.messages.push("/The socket connection has been established");
      }
  });
  }

  public ngOnDestroy() {
    this.socket.close();
}

public send() {
    if(this.chatBox) {
      console.log("Data chatbox : " + this.chatBox);
        this.socket.send(this.chatBox);
        this.chatBox = "";
    }
}

public isSystemMessage(message: string) {
  console.log("message : " + message.substring(1));
  console.log("message : " + message);
  return message.startsWith("/") ? "<strong>" + message.substring(1) + "</strong>" : message;
}
}
