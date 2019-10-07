import { createQrcode } from './../class/class.component';
import { Injectable } from '@angular/core';
import 'rxjs/add/operator/map';
import {Observable, throwError} from 'rxjs';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { classOrder } from './../home2/home2.component'
@Injectable({
  providedIn: 'root'
})
export class ClassService {
  public API = 'http://localhost:8080/';
  public APING = 'http://d23b31df.ap.ngrok.io/';
  constructor(private httpClient: HttpClient) { }

  passofClass(x){
    localStorage.setItem('passOfCouse', x);
  }

  ClientID(x){
    localStorage.setItem('clientID', x);
  }

  classSelect(x) {
    localStorage.setItem('classTeacheratSelectId' , x);
  }

  selectClass : classOrder = {
    subject_id : null,
    TSName: null,
    TSDescription: null,
    TSID: null,
    TSTeacher: null,
  };

  // public createQr (newQr:createQrcode) : Observable<createQrcode> {
  //   const headers = new HttpHeaders();
  //   headers.append('Access-Control-Allow-Origin', this.API);
  //   headers.append('Access-Control-Allow-Credentials', 'true');
  //   headers.append('Content-Type', 'application/json');
  //   console.log(JSON.stringify(newQr));
  //   return this.httpClient.post<createQrcode>(this.API + 'Attendance/' + 'Create', JSON.stringify(newQr), {headers});
  // };
  public createQr (newQr:createQrcode) : Observable<createQrcode> {
    const headers = new HttpHeaders();
    headers.append('Access-Control-Allow-Origin', this.API);
    headers.append('Access-Control-Allow-Credentials', 'true');
    headers.append('Content-Type', 'application/json');
    console.log(JSON.stringify(newQr));
    return this.httpClient.post<createQrcode>(this.APING + 'Attendance/' + 'Create', JSON.stringify(newQr), {headers});
  };

  // getmyClass(){
  //   var id = localStorage.getItem('classTeacheratSelectId');
  //   return this.httpClient.get(this.API + 'Subject/' + 'GetOneSubject/' + id);
  // }
  getmyClass(){
    var id = localStorage.getItem('classTeacheratSelectId');
    // console.log("subject_id = " + id);
    return this.httpClient.get(this.APING + 'Subject/' + 'GetOneSubject/' + id);
  }

  getmyQr(){
    var pass = localStorage.getItem('passOfCouse');
    return this.httpClient.get(this.APING + 'Attendance/' + 'getQRcode/' + pass);
  }
  // getmyQr(){
  //   var pass = localStorage.getItem('passOfCouse');
  //   return this.httpClient.get(this.API + 'Attendance/' + 'getQRcode/' + pass);
  // }

  getClientData(pass,date,timeAuthen){
    // var Cid = localStorage.getItem('clientID');
    var pass = pass;
    var date = date;
    var time = timeAuthen;
    return this.httpClient.get(this.APING + 'Attendance/' + 'info/' + pass + "/" + date + "/" + time);
  }
}
