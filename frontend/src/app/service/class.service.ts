import { createQrcode } from './../class/class.component';
import { Injectable } from '@angular/core';
import 'rxjs/add/operator/map';
import {Observable, throwError} from 'rxjs';
import {HttpClient, HttpHeaders} from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class ClassService {
  public API = 'http://localhost:8080/';
  constructor(private httpClient: HttpClient) { }

  public createQr (newQr:createQrcode) : Observable<createQrcode> {
    const headers = new HttpHeaders();
    headers.append('Access-Control-Allow-Origin', this.API);
    headers.append('Access-Control-Allow-Credentials', 'true');
    headers.append('Content-Type', 'application/json');
    console.log(JSON.stringify(newQr));
    return this.httpClient.post<createQrcode>(this.API + 'qr/' + 'createqr', JSON.stringify(newQr), {headers});
  }
}
