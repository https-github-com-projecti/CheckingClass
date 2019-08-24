import { User } from './../signup-component/signup-component.component';
import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import 'rxjs/add/operator/map';
import {Observable, throwError} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SingupService {
  public API = 'http://localhost:8080/';
  constructor(private httpClient: HttpClient) {}

  public registerUsers (user) : Observable<User> {
    const headers = new HttpHeaders();
    headers.append('Access-Control-Allow-Origin', this.API);
    headers.append('Access-Control-Allow-Credentials', 'true');
    headers.append('Content-Type', 'application/json');
    console.log(JSON.stringify(user));
    return this.httpClient.post<User>(this.API + 'newUser', JSON.stringify(user), {headers});
  }
}
