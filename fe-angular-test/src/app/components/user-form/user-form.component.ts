import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { UserService } from '../../services/user.service';
import { User } from '../../models/user.model';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatButtonModule } from '@angular/material/button';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { MatDialogModule } from '@angular/material/dialog';
import { MatCardModule } from '@angular/material/card';

// Success Dialog Component
@Component({
  selector: 'app-success-dialog',
  standalone: true,
  template: `
    <h2 mat-dialog-title>Success</h2>
    <mat-dialog-content>User has been successfully {{ isEditMode ? 'updated' : 'added' }}!</mat-dialog-content>
    <mat-dialog-actions align="end">
      <button mat-button mat-dialog-close>OK</button>
    </mat-dialog-actions>
  `,
  imports: [CommonModule, MatDialogModule, MatButtonModule]
})
export class SuccessDialogComponent {
  isEditMode: boolean = false;
}

// Confirm Cancel Dialog Component (Opsional)
@Component({
  selector: 'app-confirm-cancel-dialog',
  standalone: true,
  template: `
    <h2 mat-dialog-title>Discard Changes?</h2>
    <mat-dialog-content>Are you sure you want to discard your changes?</mat-dialog-content>
    <mat-dialog-actions align="end">
      <button mat-button (click)="dialogRef.close(false)">No</button>
      <button mat-button color="warn" (click)="dialogRef.close(true)">Yes</button>
    </mat-dialog-actions>
  `,
  imports: [CommonModule, MatDialogModule, MatButtonModule]
})
export class ConfirmCancelDialogComponent {
  constructor(public dialogRef: MatDialogRef<ConfirmCancelDialogComponent>) {}
}

@Component({
  selector: 'app-user-form',
  standalone: true,
  templateUrl: './user-form.component.html',
  styleUrls: ['./user-form.component.css'],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatButtonModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatDialogModule,
    MatCardModule
  ]
})
export class UserFormComponent implements OnInit {
  userForm: FormGroup;
  isEditMode: boolean = false;
  userId: number | null = null;

  constructor(
    private fb: FormBuilder,
    private userService: UserService,
    private router: Router,
    private route: ActivatedRoute,
    private dialog: MatDialog
  ) {
    this.userForm = this.fb.group({
      name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      dob: ['', Validators.required],
      role: ['Staff', Validators.required],
      status: ['Active', Validators.required]
    });
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id) {
        this.isEditMode = true;
        this.userId = Number(id);
        this.loadUserData();
      }
    });
  }

  loadUserData(): void {
    if (this.userId !== null) {
      const user = this.userService.getUserById(this.userId);
      if (user) {
        this.userForm.patchValue({
          name: user.name,
          email: user.email,
          dob: user.dob ? new Date(user.dob) : null,
          role: user.role,
          status: user.status
        });
      }
    }
  }

  onSubmit(): void {
    if (this.userForm.valid) {
      const formData = this.userForm.value;
      if (this.isEditMode && this.userId !== null) {
        this.userService.updateUser({ id: this.userId, ...formData });
      } else {
        this.userService.addUser(formData);
      }

      const dialogRef = this.dialog.open(SuccessDialogComponent, {
        width: '300px',
        data: { isEditMode: this.isEditMode }
      });

      dialogRef.afterClosed().subscribe(() => {
        this.router.navigate(['/users']);
      });
    }
  }

  onCancel(): void {
    if (this.userForm.dirty) {
      const dialogRef = this.dialog.open(ConfirmCancelDialogComponent, { width: '300px' });
      dialogRef.afterClosed().subscribe(result => {
        if (result) this.router.navigate(['/users']);
      });
    } else {
      this.router.navigate(['/users']);
    }
  }
}