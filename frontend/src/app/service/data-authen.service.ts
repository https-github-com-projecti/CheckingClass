import { Injectable } from '@angular/core';
import 'rxjs/add/operator/map';
import {Observable, throwError} from 'rxjs';
import {HttpClient, HttpHeaders} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class DataAuthenService {
  public APING = 'http://d23b31df.ap.ngrok.io/';
  constructor(private httpClient: HttpClient) { }

  public getDataClassAndQrcode(x,y,z){
    var data : string = x + ";" + y + ";" + z;
    localStorage.setItem('dataStudentAuthen', data);
  }

  getDataStudent(){
    var data = localStorage.getItem('dataStudentAuthen');
    console.log("----------------" + data);
    var splitted = data.split(";", 3);
    var pass : string = splitted[0];
    var date : string = splitted[1];
    var time : string = splitted[2];
    console.log(pass + ";" + date + ";" + time);
    return this.httpClient.get(this.APING + "Attendance/" + "info/" + pass + "/" + date + "/" + time);
  }
}
