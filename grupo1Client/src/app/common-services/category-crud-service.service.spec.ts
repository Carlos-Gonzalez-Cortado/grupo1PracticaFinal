import { TestBed } from '@angular/core/testing';

import { CategoryCrudServiceService } from './category-crud-service.service';

describe('CategoryCrudServiceService', () => {
  let service: CategoryCrudServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CategoryCrudServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
