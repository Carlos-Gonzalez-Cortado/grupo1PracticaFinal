import { TestBed } from '@angular/core/testing';

import { CredentialControlService } from './credential-control.service';

describe('CredentialControlService', () => {
  let service: CredentialControlService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CredentialControlService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
