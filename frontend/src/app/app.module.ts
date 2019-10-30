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

import {A11yModule} from '@angular/cdk/a11y';
import {DragDropModule} from '@angular/cdk/drag-drop';
import {PortalModule} from '@angular/cdk/portal';
import {ScrollingModule} from '@angular/cdk/scrolling';
import {CdkStepperModule} from '@angular/cdk/stepper';
import {CdkTableModule} from '@angular/cdk/table';
import {CdkTreeModule} from '@angular/cdk/tree';
import {MatBadgeModule} from '@angular/material/badge';
import {MatBottomSheetModule} from '@angular/material/bottom-sheet';
import {MatDividerModule} from '@angular/material/divider';
import {MatTreeModule} from '@angular/material/tree';
import {NoopAnimationsModule} from '@angular/platform-browser/animations';
import 'hammerjs';
import {AuthenStudentService } from './service/authen-student.service'
import { SocketService } from "./service/socket.service";
import { DataAuthenComponent } from './data-authen/data-authen.component';
import { SettingProfileComponent } from './setting-profile/setting-profile.component';
import { SimpleTimer } from 'ng2-simple-timer';
import { DataAuthenService } from './service/data-authen.service';
import { PepleService } from './service/peple.service';

const appRoutes: Routes = [
  {path: '', redirectTo: '/Home', pathMatch: 'full'},
  {path: 'Home', component: Home2Component},
  {path: 'Signup', component: SignupComponentComponent},
  {path: 'Setting', component: SettingComponent},
  {path: 'Class', component: ClassComponent},
  {path: 'People', component: PeopleComponent},
  {path: 'Authen', component: AuthenicatStudentComponent},
  {path: 'Score', component: StudentScoreComponent},
  {path: 'DataAuthen', component: DataAuthenComponent},
  {path: 'ProfileSetting', component: SettingProfileComponent},
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
    DataAuthenComponent,
    SettingProfileComponent,
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
    A11yModule,
    DragDropModule,
    PortalModule,
    ScrollingModule,
    CdkStepperModule,
    CdkTableModule,
    CdkTreeModule,
    MatBadgeModule,
    MatBottomSheetModule,
    MatDividerModule,
    MatTreeModule,
    NoopAnimationsModule,
  ],
  providers: [HomeService, SingupService, AuthenStudentService, SocketService, SimpleTimer, DataAuthenService, PepleService],
  bootstrap: [AppComponent]
})
export class AppModule {}
