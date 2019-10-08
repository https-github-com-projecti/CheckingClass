import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { environment } from '../../environments/environment';


@Injectable({
  providedIn: 'root'
})
export class PepleService {

  private APING = environment.baseUrl;
  constructor(private httpClient : HttpClient,) {}
  
    getDataStudentOfCouse(){
      var pass = localStorage.getItem('passOfCouse');
      console.log("pass = " + pass);
      return this.httpClient.get(this.APING + "Subject/" + "GetStudent/" + pass);
    }
}
