import {SingupService} from './service/singup.service';
import {HomeService} from './service/home.service';
import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {FileSelectDirective} from 'ng2-file-upload';
import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {RouterModule, Routes} from '@angular/router';
import {MatButtonModule} from '@angular/material/button';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatIconModule} from '@angular/material/icon';
import {
  MatAutocompleteModule, MatDatepickerModule, MatRadioModule, MatSlideToggleModule, MatSliderModule,
  MatMenuModule, MatGridListModule, MatSidenavModule, MatStepperModule, MatPaginatorModule, MatSortModule,
  MatTableModule, MatSnackBarModule, MatTooltipModule, MatDialogModule, MatProgressBarModule,
  MatProgressSpinnerModule, MatChipsModule, MatButtonToggleModule, MatExpansionModule, MatTabsModule,
  MatNativeDateModule, MatCardModule, MatInputModule, MatListModule, MatToolbarModule, MatFormFieldModule
} from '@angular/material';
import {HttpClientModule, HttpHeaders} from '@angular/common/http';
import {MatSelectModule} from '@angular/material/select';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {SignupComponentComponent} from './signup-component/signup-component.component';
import {Subscription} from 'rxjs';
import {Home2Component} from './home2/home2.component';
import {SettingComponent} from './setting/setting.component';
import {ClassComponent} from './class/class.component';
import {from} from 'rxjs';
import { PeopleComponent } from './people/people.component';
import { AuthenicatStudentComponent } from './authenicat-student/authenicat-student.component';
import { StudentScoreComponent } from './student-score/student-score.component';


const appRoutes: Routes = [
  {path: '', redirectTo: '/Home', pathMatch: 'full'},
  {path: 'Home', component: Home2Component},
  {path: 'Signup', component: SignupComponentComponent},
  {path: 'Setting', component: SettingComponent},
  {path: 'Class', component: ClassComponent},
  {path: 'People', component: PeopleComponent},
  {path: 'Authen', component: AuthenicatStudentComponent},
  {path: 'Score', component: StudentScoreComponent},
];

@NgModule({
  declarations: [
    AppComponent,
    SignupComponentComponent,
    Home2Component,
    SettingComponent,
    ClassComponent,
    FileSelectDirective,
    PeopleComponent,
    AuthenicatStudentComponent,
    StudentScoreComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    RouterModule.forRoot(appRoutes),
    MatButtonModule,
    MatCheckboxModule,
    BrowserAnimationsModule,
    MatIconModule,
    BrowserModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatCardModule,
    MatInputModule,
    MatListModule,
    MatSelectModule,
    MatToolbarModule,
    FormsModule,
    MatSelectModule,
    MatFormFieldModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatCardModule,
    MatInputModule,
    MatListModule,
    MatToolbarModule,
    FormsModule,
    MatCheckboxModule,
    MatAutocompleteModule,
    MatDatepickerModule,
    MatRadioModule,
    MatSelectModule,
    MatSliderModule,
    MatSlideToggleModule,
    MatMenuModule,
    MatSidenavModule,
    MatGridListModule,
    MatCardModule,
    MatStepperModule,
    MatTabsModule,
    MatExpansionModule,
    MatNativeDateModule,
    MatButtonToggleModule,
    MatChipsModule,
    MatIconModule,
    MatProgressSpinnerModule,
    MatProgressBarModule,
    MatDialogModule,
    MatTooltipModule,
    MatSnackBarModule,
    MatTableModule,
    MatSortModule,
    MatPaginatorModule,
    ReactiveFormsModule,
  ],
  providers: [HomeService, SingupService],
  bootstrap: [AppComponent]
})
export class AppModule {
}
