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
  constructor(private httpClient: HttpClient) { }

  classSelect(x) {
    console.log(x);
    localStorage.setItem('classTeacheratSelectId' , x);
  }

  selectClass : classOrder = {
    id : null,
    t_class_name: null,
    t_class_description: null,
    t_class_id: null,
    user: null,
  };

  public createQr (newQr:createQrcode) : Observable<createQrcode> {
    const headers = new HttpHeaders();
    headers.append('Access-Control-Allow-Origin', this.API);
    headers.append('Access-Control-Allow-Credentials', 'true');
    headers.append('Content-Type', 'application/json');
    console.log(JSON.stringify(newQr));
    return this.httpClient.post<createQrcode>(this.API + 'qr/' + 'createqr', JSON.stringify(newQr), {headers});
  };

  getmyClass(){
    var id = localStorage.getItem('classTeacheratSelectId');
    return this.httpClient.get(this.API + 'class/' + 'selectClass/' + id);
  }

  getmyQr(){
    var user = localStorage.getItem('isLogin');
    return this.httpClient.get(this.API + 'qr/' + 'myQr/' + user);
  }
}
