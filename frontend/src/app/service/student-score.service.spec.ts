import { TestBed } from '@angular/core/testing';

import { StudentScoreService } from './student-score.service';

describe('StudentScoreService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: StudentScoreService = TestBed.get(StudentScoreService);
    expect(service).toBeTruthy();
  });
});
