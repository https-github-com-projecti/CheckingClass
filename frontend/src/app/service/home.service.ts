import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from "rxjs/Observable";

@Injectable({
  providedIn: 'root'
})
export class HomeService {
  public API = '//localhost:8080';
  private serviceUrl = 'http://localhost:8080/';
  private serviceUrl2 = 'http://localhost:8080/Results/'
  constructor(private httpClient: HttpClient) { }

  login(a,b){                
    localStorage.setItem('isLogin', a,);
    localStorage.setItem('stateLogin',b);
  }

  getTodoList(): Observable<any> {
    return this.httpClient.get<any>(this.serviceUrl);
  }
}
