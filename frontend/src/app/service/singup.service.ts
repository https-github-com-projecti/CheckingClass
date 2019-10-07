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
  public APING = 'http://d23b31df.ap.ngrok.io/';
  constructor(private httpClient: HttpClient) {}

  // public registerUsers (user:User) : Observable<User> {
  //   const headers = new HttpHeaders();
  //   headers.append('Access-Control-Allow-Origin', this.API);
  //   headers.append('Access-Control-Allow-Credentials', 'true');
  //   headers.append('Content-Type', 'application/json');
  //   console.log(JSON.stringify(user));
  //   return this.httpClient.post<User>(this.API + 'user/' + 'Add', JSON.stringify(user), {headers});
  // }
  public registerUsers (user:User) : Observable<User> {
    const headers = new HttpHeaders();
    headers.append('Access-Control-Allow-Origin', this.API);
    headers.append('Access-Control-Allow-Credentials', 'true');
    headers.append('Content-Type', 'application/json');
    console.log(JSON.stringify(user));
    return this.httpClient.post<User>(this.APING + 'user/' + 'Add', JSON.stringify(user), {headers});
  }
}
