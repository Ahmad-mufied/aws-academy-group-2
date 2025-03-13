import { Component, Inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import { User } from '../../models/user.model';

@Component({
  selector: 'app-user-detail',
  standalone: true,
  imports: [CommonModule, MatDialogModule, MatButtonModule],
  template: `
    <h2 mat-dialog-title>User Details</h2>
    <mat-dialog-content>
      <div class="detail-row">
        <strong>Name:</strong> {{data.name}}
      </div>
      <div class="detail-row">
        <strong>Email:</strong> {{data.email}}
      </div>
      <div class="detail-row">
        <strong>Date of Birth:</strong> {{data.dob | date}}
      </div>
      <div class="detail-row">
        <strong>Role:</strong> {{data.role}}
      </div>
      <div class="detail-row">
        <strong>Status:</strong> {{data.status}}
      </div>
      <div class="detail-row">
        <strong>Registered Date:</strong> {{data.registeredDate | date}}
      </div>
      <div class="detail-row">
        <strong>Assigned Products:</strong>
        <ul *ngIf="data.products?.length">
          <li *ngFor="let product of data.products">{{product}}</li>
        </ul>
        <span *ngIf="!data.products?.length">No products assigned</span>
      </div>
    </mat-dialog-content>
    <mat-dialog-actions align="end">
      <button mat-button mat-dialog-close>Close</button>
    </mat-dialog-actions>
  `,
  styles: [`
    .detail-row {
      margin-bottom: 16px;
    }
    strong {
      font-weight: 500;
      margin-right: 8px;
    }
    ul {
      margin: 8px 0 0 20px;
    }
  `]
})
export class UserDetailComponent {
  constructor(
    public dialogRef: MatDialogRef<UserDetailComponent>,
    @Inject(MAT_DIALOG_DATA) public data: User
  ) {}
}