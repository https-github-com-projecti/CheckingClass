import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';


@Injectable({
  providedIn: 'root'
})
export class PepleService {

  private APING = "http://d23b31df.ap.ngrok.io/";
  constructor(private httpClient : HttpClient,) {}
  
    getDataStudentOfCouse(){
      var pass = localStorage.getItem('passOfCouse');
      console.log("pass = " + pass);
      return this.httpClient.get(this.APING + "Subject/" + "GetStudent/" + pass);
    }
}
