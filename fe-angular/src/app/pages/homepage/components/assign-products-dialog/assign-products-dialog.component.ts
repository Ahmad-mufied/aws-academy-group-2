// src/app/pages/homepage/components/assign-products-dialog/assign-products-dialog.component.ts

import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef, MatDialogContent, MatDialogActions } from '@angular/material/dialog';
import { User } from '../user-list/user-list.component';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { FormsModule } from '@angular/forms';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';

interface ProductNode {
  name: string;
  children?: ProductNode[];
  completed: boolean;
}

@Component({
  selector: 'app-assign-products-dialog',
  standalone: true,
  templateUrl: './assign-products-dialog.component.html',
  styleUrls: ['./assign-products-dialog.component.css'],
  imports: [CommonModule, MatButtonModule, MatCheckboxModule, FormsModule, MatListModule, MatIconModule, MatDialogContent, MatDialogActions],
})
export class AssignProductsDialogComponent {
  products: ProductNode[] = [
    {
      name: 'Product A',
      completed: false,
      children: [
        { name: 'Subtask A1', completed: false },
        { name: 'Subtask A2', completed: false },
      ],
    },
    {
      name: 'Product B',
      completed: false,
      children: [
        { name: 'Subtask B1', completed: false },
        { name: 'Subtask B2', completed: false },
        { name: 'Subtask B3', completed: false },
      ],
    },
    { name: 'Product C', completed: false },
  ];

  constructor(
    public dialogRef: MatDialogRef<AssignProductsDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: User
  ) {}

  updateProduct(product: ProductNode, event: any) {
    product.completed = event.checked;
    if (product.children) {
      product.children.forEach(child => (child.completed = event.checked));
    }
  }

  updateSubtask(product: ProductNode, subtask: ProductNode, event: any) {
    subtask.completed = event.checked;
    product.completed = product.children ? product.children.every(child => child.completed) : product.completed;
  }

  onCancelClick(): void {
    this.dialogRef.close();
  }

  onSubmitClick(): void {
    const selectedProducts = this.products.filter(product => product.completed);
    console.log('Selected Products:', selectedProducts);
    this.dialogRef.close(selectedProducts);
  }
}
