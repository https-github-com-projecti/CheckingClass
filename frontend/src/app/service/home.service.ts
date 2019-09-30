import { userLogin, classOrder} from './../home2/home2.component';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { Injectable } from '@angular/core';
import 'rxjs/add/operator/map';
import {Observable, throwError} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HomeService {
  public API = 'http://localhost:8080/';
  public APING = 'http://b894c639.ap.ngrok.io/';
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
  //   return this.httpClient.post<userLogin>(this.API + 'user/' + 'loginUser', JSON.stringify(login), {headers});
  // }
  public LoginUser(login : userLogin): Observable<userLogin> {
    const headers = new HttpHeaders();
    headers.append('Access-Control-Allow-Origin', this.API);
    headers.append('Access-Control-Allow-Credentials', 'true');
    headers.append('Content-Type', 'application/json');
    console.log(JSON.stringify(login));
    return this.httpClient.post<userLogin>(this.APING + 'user/' + 'login', JSON.stringify(login), {headers});
  }

  public CreateClass(newClass : classOrder): Observable<classOrder> {
    const headers = new HttpHeaders();
    headers.append('Access-Control-Allow-Origin', this.API);
    headers.append('Access-Control-Allow-Credentials', 'true');
    headers.append('Content-Type', 'application/json');
    console.log(JSON.stringify(newClass));
    return this.httpClient.post<classOrder>(this.API + 'class/' + 'newClass', JSON.stringify(newClass), {headers});
  }

  public getClass(){
    var us = localStorage.getItem('isLogin');
    return this.httpClient.get(this.APING + 'class/' + 'myClass/' + us);
  }

  getAllClass() {
    return this.httpClient.get(this.API + 'class/' + 'allClass');
  }

  getUserdata(){
    var us = localStorage.getItem('isLogin');
    return this.httpClient.get(this.APING + 'user/' + 'GETONE/' + us );
  }

  getGetPic(){
    var us = localStorage.getItem('isLogin');
    return this.httpClient.get(this.API + 'user/' + 'getMyPic/' + us);
  }
}
