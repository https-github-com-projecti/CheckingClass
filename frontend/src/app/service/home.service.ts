import { Article } from './../Entity/article.entity';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { Observable } from 'rxjs';

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
}
