import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { RouterModule, Routes } from '@angular/router';
import {MatButtonModule} from '@angular/material/button';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatIconModule} from '@angular/material/icon';
import {MatAutocompleteModule,MatDatepickerModule,MatRadioModule,MatSlideToggleModule,MatSliderModule,
  MatMenuModule,MatGridListModule,MatSidenavModule,MatStepperModule,MatPaginatorModule,MatSortModule, 
  MatTableModule, MatSnackBarModule, MatTooltipModule,MatDialogModule, MatProgressBarModule, 
  MatProgressSpinnerModule, MatChipsModule, MatButtonToggleModule,MatExpansionModule, MatTabsModule, 
  MatNativeDateModule, MatCardModule, MatInputModule,MatListModule,MatToolbarModule, MatFormFieldModule
} from '@angular/material';
import { HttpClientModule } from '@angular/common/http';
import {MatSelectModule} from '@angular/material/select';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { SignupComponentComponent } from './signup-component/signup-component.component';
import { LoginComponentComponent } from './login-component/login-component.component';
import { Subscription } from 'rxjs';
import { Home2Component } from './home2/home2.component';
import { SettingComponent } from './setting/setting.component'

const appRoutes: Routes = [
  {path: '',   redirectTo: '/Home',    pathMatch: 'full'},
  { path: 'Home', component: Home2Component},
  { path: 'Signup', component: SignupComponentComponent},
  { path: 'Login', component: LoginComponentComponent},
  { path: 'Setting', component: SettingComponent},
]

@NgModule({
  declarations: [
    AppComponent,
    SignupComponentComponent,
    LoginComponentComponent,
    Home2Component,
    SettingComponent,
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
    RouterModule.forRoot(appRoutes),
    BrowserAnimationsModule,
    MatButtonModule,
    MatCardModule,
    MatInputModule,
    MatListModule,
    MatToolbarModule,
    FormsModule,
    RouterModule.forRoot(appRoutes),
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
    RouterModule.forRoot(appRoutes),
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
