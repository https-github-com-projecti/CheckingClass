import { TestBed } from '@angular/core/testing';

import { PepleService } from './peple.service';

describe('PepleService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: PepleService = TestBed.get(PepleService);
    expect(service).toBeTruthy();
  });
});
