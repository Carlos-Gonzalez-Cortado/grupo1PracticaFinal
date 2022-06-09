import { TestBed } from '@angular/core/testing';

import { VideoCrudServiceService } from './video-crud-service.service';

describe('VideoCrudServiceService', () => {
  let service: VideoCrudServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(VideoCrudServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
