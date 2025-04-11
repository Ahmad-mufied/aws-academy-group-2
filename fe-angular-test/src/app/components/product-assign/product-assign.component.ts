import { Component, Inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { FormsModule } from '@angular/forms';
import { User } from '../../models/user.model';
import { MatDialog } from '@angular/material/dialog';
import { ConfirmCancelDialogComponent } from '../confirm-cancel-dialog/confirm-cancel-dialog.component';

@Component({
  selector: 'app-product-assign',
  standalone: true,
  imports: [
    CommonModule,
    MatDialogModule,
    MatCheckboxModule,
    MatButtonModule,
    FormsModule
  ],
  templateUrl: './product-assign.component.html',
  styleUrls: ['./product-assign.component.css']
})
export class ProductAssignComponent implements OnInit {
  availableProducts = [
    {
      name: 'Parent Task 1',
      selected: false,
      subtasks: [
        { name: 'Child task 1', selected: false },
        { name: 'Child task 2', selected: false },
        { name: 'Child task 3', selected: false }
      ]
    }
  ];
  initialState: any; // Simpen state awal buat cek perubahan

  constructor(
    public dialogRef: MatDialogRef<ProductAssignComponent>,
    @Inject(MAT_DIALOG_DATA) public data: User,
    private dialog: MatDialog // Tambah MatDialog untuk confirm dialog
  ) {}

  ngOnInit(): void {
    // Inisialisasi state awal berdasarkan data user
    this.availableProducts.forEach(product => {
      product.selected = this.isProductAssigned(product.name);
      product.subtasks.forEach(subtask => {
        subtask.selected = this.isProductAssigned(subtask.name);
      });
    });
    // Simpen salinan state awal
    this.initialState = JSON.parse(JSON.stringify(this.availableProducts));
  }

  isProductAssigned(productName: string): boolean {
    return this.data.products?.includes(productName) || false;
  }

  toggleParentSelection(product: any) {
    product.subtasks.forEach((subtask: { selected: boolean }) => {
      subtask.selected = product.selected;
    });
  }

  isDirty(): boolean {
    // Bandingin state awal sama state sekarang
    return JSON.stringify(this.initialState) !== JSON.stringify(this.availableProducts);
  }

  onSave(): void {
    const selectedProducts = this.availableProducts
      .flatMap(product => product.subtasks.filter(subtask => subtask.selected))
      .map(subtask => subtask.name);
    this.dialogRef.close(selectedProducts);
  }

  onCancel(): void {
    if (this.isDirty()) {
      const dialogRef = this.dialog.open(ConfirmCancelDialogComponent, { width: '300px' });
      dialogRef.afterClosed().subscribe(result => {
        if (result) this.dialogRef.close(); // Tutup dialog kalau "Yes"
      });
    } else {
      this.dialogRef.close(); // Langsung tutup kalau nggak ada perubahan
    }
  }
}