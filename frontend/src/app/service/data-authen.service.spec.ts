import { TestBed } from '@angular/core/testing';

import { DataAuthenService } from './data-authen.service';

describe('DataAuthenService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: DataAuthenService = TestBed.get(DataAuthenService);
    expect(service).toBeTruthy();
  });
});
