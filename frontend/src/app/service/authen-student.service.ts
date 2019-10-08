import { Injectable } from '@angular/core';
import 'rxjs/add/operator/map';
import {Observable, throwError} from 'rxjs';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { classOrder } from './../home2/home2.component'
import {PeriodicElement} from '../authenicat-student/authenicat-student.component'
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AuthenStudentService {
  public API = 'http://localhost:8080/';
  private APING = environment.baseUrl;
  constructor(private httpClient: HttpClient) { }

  // getAuthenData():Observable<PeriodicElement[]> {
  //   return this.httpClient.get<PeriodicElement[]>(this.API + 'authen/' + 'authenData');
  // }

  getAuthenData():Observable<PeriodicElement[]> {
    var pass = localStorage.getItem('passOfCouse');
    return this.httpClient.get<PeriodicElement[]>(this.APING + 'Attendance/' + 'getAttendance/' + pass);
  }
}
