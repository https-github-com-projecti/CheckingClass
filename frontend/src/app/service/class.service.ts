import { createQrcode, date } from './../class/class.component';
import { Injectable } from '@angular/core';
import 'rxjs/add/operator/map';
import {Observable, throwError} from 'rxjs';
import {HttpClient, HttpHeaders, HttpParams} from '@angular/common/http';
import { classOrder } from './../home2/home2.component'
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ClassService {
  public API = 'http://localhost:8080/';
  public APING = environment.baseUrl;
  private authType : string;
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
    TStimesubject: null,
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

  // public getmyClass(){
  //   var id = localStorage.getItem('classTeacheratSelectId');
  //   return this.httpClient.get(this.API + 'Subject/' + 'GetOneSubject/' + id);
  // }
  public getmyClass(){
    var id = localStorage.getItem('classTeacheratSelectId');
    // console.log("subject_id = " + id);
    return this.httpClient.get(this.APING + 'Subject/' + 'GetOneSubject/' + id);
  }

  public getmyQr(){
    var pass = localStorage.getItem('passOfCouse');
    return this.httpClient.get(this.APING + 'Attendance/' + 'getQRcode/' + pass);
  }
  // public getmyQr(){
  //   var pass = localStorage.getItem('passOfCouse');
  //   return this.httpClient.get(this.API + 'Attendance/' + 'getQRcode/' + pass);
  // }

  public getClientData(pass,date,timeAuthen){
    // var Cid = localStorage.getItem('clientID');
    var pass = pass;
    var date = date;
    var time = timeAuthen;
    return this.httpClient.get(this.APING + 'Attendance/' + 'info/' + pass + "/" + date + "/" + time);
  }

  public getTimeLimite(newDate : date){
    // console.log(newDate);
    const headers = new HttpHeaders();
    headers.append('Access-Control-Allow-Origin', this.API);
    headers.append('Access-Control-Allow-Credentials', 'true');
    headers.append('Content-Type', 'application/json');
    return this.httpClient.post(this.APING + "Attendance/" + "time", JSON.stringify(newDate), {headers}); 
  }

  public getTStimesubject(){
    var pass = localStorage.getItem('passOfCouse');
    return this.httpClient.get(this.APING + 'Subject/' + 'time/' + pass);
  }
}
