import { userLogin, classOrder } from './../home2/home2.component';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { Injectable } from '@angular/core';
import 'rxjs/add/operator/map';
import {Observable, throwError} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HomeService {
  public API = 'http://localhost:8080/';
  constructor(private httpClient: HttpClient) { }

  login(a,b){                
    localStorage.setItem('isLogin', a,);
    localStorage.setItem('stateLogin',b);
  }
  getUser() {
    return this.httpClient.get(this.API + 'home/' + 'ping');
  }

  public LoginUser(login : userLogin): Observable<userLogin> {
    const headers = new HttpHeaders();
    headers.append('Access-Control-Allow-Origin', this.API);
    headers.append('Access-Control-Allow-Credentials', 'true');
    headers.append('Content-Type', 'application/json');
    console.log(JSON.stringify(login));
    return this.httpClient.post<userLogin>(this.API + 'user/' + 'loginUser', JSON.stringify(login), {headers});
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
    return this.httpClient.get(this.API + 'class/' + 'myClass/' + us);
  }

  getAllClass() {
    return this.httpClient.get(this.API + 'class/' + 'allClass');
  }
}
