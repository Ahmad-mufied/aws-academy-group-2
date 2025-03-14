import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AssignProductsDialogComponent } from './assign-products-dialog.component';

describe('AssignProductsDialogComponent', () => {
  let component: AssignProductsDialogComponent;
  let fixture: ComponentFixture<AssignProductsDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AssignProductsDialogComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AssignProductsDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
