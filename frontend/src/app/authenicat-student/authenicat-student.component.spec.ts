import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AuthenicatStudentComponent } from './authenicat-student.component';

describe('AuthenicatStudentComponent', () => {
  let component: AuthenicatStudentComponent;
  let fixture: ComponentFixture<AuthenicatStudentComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AuthenicatStudentComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AuthenicatStudentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
