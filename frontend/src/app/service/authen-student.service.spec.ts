import { TestBed } from '@angular/core/testing';

import { AuthenStudentService } from './authen-student.service';

describe('AuthenStudentService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: AuthenStudentService = TestBed.get(AuthenStudentService);
    expect(service).toBeTruthy();
  });
});
