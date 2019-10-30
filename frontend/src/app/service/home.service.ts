import { userLogin, classOrder, timeSubjectAuthen} from './../home2/home2.component';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { Injectable } from '@angular/core';
import 'rxjs/add/operator/map';
import {Observable, throwError} from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class HomeService {
  public API = 'http://localhost:8080/';
  public APING = environment.baseUrl;
  public data : any;
  constructor(private httpClient: HttpClient) { }

  setID(z){
    localStorage.setItem('id', z);
  }

  status(x){
    localStorage.setItem('statusLogin', x);
  }

  login(a,b){                
    localStorage.setItem('isLogin', a);
    localStorage.setItem('stateLogin',b);
  }
  getStatus() {
    return this.httpClient.get(this.API + 'home/' + 'ping');
  }

  // public LoginUser(login : userLogin): Observable<userLogin> {
  //   const headers = new HttpHeaders();
  //   headers.append('Access-Control-Allow-Origin', this.API);
  //   headers.append('Access-Control-Allow-Credentials', 'true');
  //   headers.append('Content-Type', 'application/json');
  //   console.log(JSON.stringify(login));
  //   return this.httpClient.post<userLogin>(this.API + 'user/' + 'login', JSON.stringify(login), {headers});
  // }
  public LoginUser(login : userLogin): Observable<userLogin> {
    const headers = new HttpHeaders();
    headers.append('Access-Control-Allow-Origin', this.API);
    headers.append('Access-Control-Allow-Credentials', 'true');
    headers.append('Content-Type', 'application/json');
    // console.log(JSON.stringify(login));
    return this.httpClient.post<userLogin>(this.APING + 'user/' + 'login', JSON.stringify(login), {headers});
  }

  // public CreateClass(newClass : classOrder): Observable<classOrder> {
  //   const headers = new HttpHeaders();
  //   headers.append('Access-Control-Allow-Origin', this.API);
  //   headers.append('Access-Control-Allow-Credentials', 'true');
  //   headers.append('Content-Type', 'application/json');
  //   console.log(JSON.stringify(newClass));
  //   return this.httpClient.post<classOrder>(this.API + 'Subject/' + 'Add', JSON.stringify(newClass), {headers});
  // }
  public CreateClass(newClass : classOrder): Observable<classOrder> {
    const headers = new HttpHeaders();
    headers.append('Access-Control-Allow-Origin', this.API);
    headers.append('Access-Control-Allow-Credentials', 'true');
    headers.append('Content-Type', 'application/json');
    // console.log(JSON.stringify(newClass));
    return this.httpClient.post<classOrder>(this.APING + 'Subject/' + 'Add', JSON.stringify(newClass), {headers});
  }

  public getClass(){
    var id = localStorage.getItem('id');
    return this.httpClient.get(this.APING + 'Subject/' + 'GetMySubject/' + id);
  }
  // public getClass(){
  //   var id = localStorage.getItem('id');
  //   return this.httpClient.get(this.API + 'Subject/' + 'GetMySubject/' + id);
  // }

  getUserdata(){
    var us = localStorage.getItem('isLogin');
    return this.httpClient.get(this.APING + 'user/' + 'GETONE/' + us );
  }
  // getUserdata(){
  //   var us = localStorage.getItem('isLogin');
  //   return this.httpClient.get(this.API + 'user/' + 'GETONE/' + us );
  // }

  getGetPic(){
    var id = localStorage.getItem('id');
    // console.log(id);
    return this.httpClient.get(this.APING + 'user/' + 'getMyPic/' + id);
  }
  // getGetPic(){
  //   var id = localStorage.getItem('id');
  //   console.log(id);
  //   return this.httpClient.get(this.API + 'user/' + 'getMyPic/' + id);
  // }

  public updateTime(newTime : timeSubjectAuthen){
    // console.log(newDate);
    var pass = localStorage.getItem('passOfCouse');
    const headers = new HttpHeaders();
    headers.append('Access-Control-Allow-Origin', this.API);
    headers.append('Access-Control-Allow-Credentials', 'true');
    headers.append('Content-Type', 'application/json');
    return this.httpClient.put(this.APING + "Subject/" + "TimeEdit/" + pass, JSON.stringify(newTime), {headers}); 
  }
}
