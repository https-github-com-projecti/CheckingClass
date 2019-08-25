import { Subscription } from 'rxjs';
import {newUser} from './../Entity/newUser.entity';
import {SingupService} from './../service/singup.service';
import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroupDirective, NgForm, Validators} from '@angular/forms';
import {ErrorStateMatcher} from '@angular/material/core';
import { DataSource } from '@angular/cdk/collections';
import { DomSanitizer } from '@angular/platform-browser';
import { HttpClient } from '@angular/common/http';

export interface User {
  tName: string;
  userName: string;
  tId: string;
  tEmail: string;
  tWorkPlace: string;
  tPassword: string;
}

export class MyErrorStateMatcher implements ErrorStateMatcher {
  isErrorState(control: FormControl | null, form: FormGroupDirective | NgForm | null): boolean {
    const isSubmitted = form && form.submitted;
    return !!(control && control.invalid && (control.dirty || control.touched || isSubmitted));
  }
}

@Component({
  selector: 'app-signup-component',
  templateUrl: './signup-component.component.html',
  styleUrls: ['./signup-component.component.css']
})
export class SignupComponentComponent implements OnInit {
  // tName:string = "";
  // userName:string = "";
  // tId:string = "";
  // tEmail:string = "";
  // tWorkPlace:string = "";
  // tPassword:string = "";
  tComfirmPassword = '';

  emailFormControl = new FormControl('', [
    Validators.required,
    Validators.email,
  ]);
  matcher = new MyErrorStateMatcher();

  constructor(private singupService: SingupService, private httpClient: HttpClient) {
  }

  newData: User = {
    tName: '',
    userName: '',
    tId: '',
    tEmail: '',
    tWorkPlace: '',
    tPassword: '',
  };


  // userFile: string = 'https://via.placeholder.com/150/FFFFFF/000000/?text=Check%20name%20';
  // readURL(event:any) {
  //   if (event.target.files && event.target.files[0]) {
  //     var reader = new FileReader();
  //     reader.onload = (event:any) => {
  //      this.userFile = event.target.result;
  //     }
  //     reader.readAsDataURL(event.target.files[0]);
  //   }
  // }
  ngOnInit() {
    
  }

  save() {

    console.log(this.newData);
    this.singupService.registerUsers(this.newData).subscribe(
      data => {
        if (data){
          alert(data);
        }
        else {
          alert('success');
        }
      }      
    );
  }
}
