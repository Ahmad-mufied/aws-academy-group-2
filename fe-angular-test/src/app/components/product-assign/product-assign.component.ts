import { Component, Inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { FormsModule } from '@angular/forms';
import { User } from '../../models/user.model';

@Component({
  selector: 'app-product-assign',
  standalone: true,
  imports: [CommonModule, MatDialogModule, MatCheckboxModule, MatButtonModule, FormsModule],
  template: `
    <h2 mat-dialog-title>Assign Products</h2>
    <mat-dialog-content>
      <div *ngFor="let product of availableProducts" class="product-item">
        <mat-checkbox
          [(ngModel)]="product.selected"
          [checked]="isProductAssigned(product.name)">
          {{product.name}}
        </mat-checkbox>
      </div>
    </mat-dialog-content>
    <mat-dialog-actions align="end">
      <button mat-button mat-dialog-close>Cancel</button>
      <button mat-raised-button color="primary" (click)="onSave()">Save</button>
    </mat-dialog-actions>
  `,
  styles: [`
    .product-item {
      margin-bottom: 16px;
    }
  `]
})
export class ProductAssignComponent {
  availableProducts = [
    { name: 'Product Task 1', selected: false },
    { name: 'Child task 1', selected: false },
    { name: 'Child task 2', selected: false },
    { name: 'Child task 3', selected: false }
  ];

  constructor(
    public dialogRef: MatDialogRef<ProductAssignComponent>,
    @Inject(MAT_DIALOG_DATA) public data: User
  ) {
    this.availableProducts.forEach(product => {
      product.selected = this.isProductAssigned(product.name);
    });
  }

  isProductAssigned(productName: string): boolean {
    return this.data.products?.includes(productName) || false;
  }

  onSave(): void {
    const selectedProducts = this.availableProducts
      .filter(product => product.selected)
      .map(product => product.name);
    this.dialogRef.close(selectedProducts);
  }
}