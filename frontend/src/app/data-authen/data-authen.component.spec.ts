import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DataAuthenComponent } from './data-authen.component';

describe('DataAuthenComponent', () => {
  let component: DataAuthenComponent;
  let fixture: ComponentFixture<DataAuthenComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DataAuthenComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DataAuthenComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
