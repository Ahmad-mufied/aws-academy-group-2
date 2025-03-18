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
  templateUrl: './product-assign.component.html',
  styleUrls: ['./product-assign.component.css']
})
export class ProductAssignComponent {
  availableProducts = [
    {
      name: 'Parent Task 1',
      completed: false,
      selected: false, // ✅ Tambahkan selected untuk Parent Task
      subtasks: [
        { name: 'Child task 1', selected: false },
        { name: 'Child task 2', selected: false },
        { name: 'Child task 3', selected: false }
      ]
    }
  ];

  constructor(
    public dialogRef: MatDialogRef<ProductAssignComponent>,
    @Inject(MAT_DIALOG_DATA) public data: User
  ) {
    this.availableProducts.forEach(product => {
      product.selected = this.isProductAssigned(product.name);
      product.subtasks.forEach(subtask => {
        subtask.selected = this.isProductAssigned(subtask.name);
      });
    });
  }

  isProductAssigned(productName: string): boolean {
    return this.data.products?.includes(productName) || false;
  }

  toggleParentSelection(product: any) {
    product.subtasks.forEach((subtask: { selected: any; }) => {
      subtask.selected = product.selected;
    });
  }

  onSave(): void {
    const selectedProducts = this.availableProducts
      .flatMap(product => product.subtasks.filter(subtask => subtask.selected))
      .map(subtask => subtask.name);

    this.dialogRef.close(selectedProducts);
  }
}
