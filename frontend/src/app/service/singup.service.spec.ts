import { TestBed } from '@angular/core/testing';

import { SingupService } from './singup.service';

describe('SingupService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: SingupService = TestBed.get(SingupService);
    expect(service).toBeTruthy();
  });
});
