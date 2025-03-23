import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { MatButtonModule } from '@angular/material/button';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import { UserService } from '../../services/user.service';
import { User } from '../../models/user.model';

// Dialog sukses
@Component({
  selector: 'app-success-dialog',
  standalone: true,
  template: `
    <h2 mat-dialog-title>Success</h2>
    <mat-dialog-content>User has been successfully saved!</mat-dialog-content>
    <mat-dialog-actions>
      <button mat-button mat-dialog-close>OK</button>
    </mat-dialog-actions>
  `,
  imports: [
    CommonModule, 
    MatDialogModule, 
    MatButtonModule
  ]
})
export class SuccessDialogComponent {}

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
    MatDialogModule
  ]
})
export class UserFormComponent implements OnInit {
  userForm!: FormGroup;
  userId: number | null = null;
  user?: User;
  isEditMode: boolean = false;

  constructor(
    private fb: FormBuilder,
    private userService: UserService,
    private router: Router,
    private route: ActivatedRoute,
    private dialog: MatDialog
  ) {}

  ngOnInit(): void {
    this.initForm();

    // Ambil userId dari URL
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id) {
        this.userId = Number(id);
        this.loadUserData();
      }
    });
  }

  initForm(): void {
    this.userForm = this.fb.group({
      name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      dob: ['', Validators.required],
      role: ['Staff', Validators.required],
      status: ['Active', Validators.required]
    });
  }

  loadUserData(): void {
    if (this.userId !== null) {
      this.user = this.userService.getUserById(this.userId);
      if (this.user) {
        this.isEditMode = true;
        this.userForm.patchValue({
          name: this.user.name,
          email: this.user.email,
          dob: this.user.dob ? new Date(this.user.dob) : null,
          role: this.user.role,
          status: this.user.status
        });
      }
    }
  }

  onSubmit(): void {
    if (this.userForm.valid) {
      if (this.isEditMode && this.userId !== null) {
        this.userService.updateUser({ ...this.user, ...this.userForm.value });
      } else {
        this.userService.addUser(this.userForm.value);
      }

      this.dialog.open(SuccessDialogComponent, { width: '300px' });

      this.dialog.afterAllClosed.subscribe(() => {
        this.router.navigate(['/users']);
      });
    }
  }

  onCancel(): void {
    this.router.navigate(['/users']);
  }
}
