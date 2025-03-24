import { Component, Inject } from '@angular/core';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { MatChipsModule } from '@angular/material/chips';
import { CommonModule } from '@angular/common';
import { User } from '../../models/user.model';

@Component({
  selector: 'app-user-detail',
  standalone: true,
  imports: [
    CommonModule,
    MatDialogModule,
    MatCardModule,
    MatButtonModule,
    MatListModule,
    MatIconModule,
    MatChipsModule
  ],
  templateUrl: './user-detail.component.html',
  styleUrls: ['./user-detail.component.css']
})
export class UserDetailComponent {
  constructor(
    public dialogRef: MatDialogRef<UserDetailComponent>,
    @Inject(MAT_DIALOG_DATA) public data: User,
    private dialog: MatDialog // Inject MatDialog service
  ) {}

  openDialog() {
    this.dialog.open(UserDetailComponent, {
      data: {
        name: 'John Doe',
        email: 'john.doe@example.com',
        dob: new Date('1990-01-01'),
        role: 'User',
        status: 'Active',
        registeredDate: new Date(),
        products: ['Product A', 'Product B']
      }
    });
  }
}
