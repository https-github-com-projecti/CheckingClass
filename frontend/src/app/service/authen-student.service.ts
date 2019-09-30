import { Injectable } from '@angular/core';
import 'rxjs/add/operator/map';
import {Observable, throwError} from 'rxjs';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { classOrder } from './../home2/home2.component'
import {PeriodicElement} from '../authenicat-student/authenicat-student.component'

@Injectable({
  providedIn: 'root'
})
export class AuthenStudentService {
  public API = 'http://localhost:8080/';
  constructor(private httpClient: HttpClient) { }

  getAuthenData():Observable<PeriodicElement[]> {
    return this.httpClient.get<PeriodicElement[]>(this.API + 'authen/' + 'authenData');
  }
}
